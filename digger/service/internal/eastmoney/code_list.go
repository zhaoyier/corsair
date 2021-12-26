package eastmoney

import (
	"fmt"
	"time"

	"git.ezbuy.me/ezbuy/corsair/digger/service/internal/common/webapi"
	"git.ezbuy.me/ezbuy/corsair/digger/service/internal/job"
	orm "git.ezbuy.me/ezbuy/corsair/digger/service/internal/model"
	trpc "git.ezbuy.me/ezbuy/corsair/digger/service/internal/rpc"
	log "github.com/Sirupsen/logrus"
	ezdb "github.com/ezbuy/ezorm/db"
	mgo "gopkg.in/mgo.v2"
)

func GetCodeListTicker() {
	GetCodeList()
	job.UpdateJob(trpc.FunctionType_FunctionTypeCodeList)
}

func GetCodeListOnce() {
	getCodeListOnce.Do(func() {
		GetCodeList()
	})
}

func GetCodeList() {
	sess, col := orm.CNSecucodeMgr.GetCol()
	defer sess.Close()

	resp, inc := new(StockList), int32(1)
	for {
		if inc > 70 { //5600个
			break
		}

		if err := webapi.GetEastmoneyCode(inc, 80, resp); err != nil || resp.Data == nil {
			log.Errorf("eastmoney get failed: %d|%+v\n", inc, err)
			inc++
			continue
		}

		if err := updateCodeList(resp, col); err != nil && mgo.IsDup(err) {
			log.Errorf("update code list failed: %+v\n", err)
			return
		}

		inc++
	}
}

func updateCodeList(req *StockList, col *mgo.Collection) error {
	if req.Data == nil || req.Data.Diff == nil {
		return nil
	}

	// 沪市A股票买卖的代码是以600、601或603开头的6位数编码。
	// 深市A股票买卖的代码是以00开头的6位数编码。
	// 深市创业板股票买卖的代码是以30开头的6位数编码，
	// 沪市科创板股票买卖的代码是以688开头的6位数编码。
	for _, val := range req.Data.Diff {
		exchange := "SH"
		prefix := val.Secucode[0:3]
		switch prefix {
		case "600":
		case "601":
		case "603":
		case "688":
			exchange = "SH"
		case "835":
		case "836":
			exchange = "BJ"
		default:
			exchange = "SZ"
		}

		if exchange == "BJ" {
			continue
		}

		if err := applyCNSecucode(val.Secucode, exchange, val.Name, col); err != nil {
			log.Errorf("apply failed: %+v|%+v\n", val, err)
		}
		if err := applyGPDaily(val, exchange); err != nil {
			log.Errorf("apply failed: %+v|%+v\n", val, err)
		}
	}

	return nil
}

func applyCNSecucode(secu, exchange, name string, col *mgo.Collection) error {
	r, err := orm.CNSecucodeMgr.FindOneBySecucode(secu)
	if r != nil {
		return fmt.Errorf("unupdate %+v", err)
	}

	secucode := exchange + "." + secu
	query := ezdb.M{"Secucode": secucode}

	change := mgo.Change{
		Update: ezdb.M{
			"$set": ezdb.M{
				"Name":       name,
				"CreateDate": time.Now().Unix(),
				"UpdateDate": time.Now().Unix(),
			},
		},
		Upsert:    true,
		ReturnNew: true,
	}

	if _, err := col.Find(query).Apply(change, nil); err != nil {
		return fmt.Errorf("%s|%q", secucode, err)
	}
	return nil
}

func applyGPDaily(data *CodeBase, exchange string) error {
	secucode := exchange + "." + data.Secucode
	createDate := getZeroTS()

	result, err := orm.GPDailyMgr.FindOneBySecucodeCreateDate(secucode, createDate)
	if err != nil && err != mgo.ErrNotFound {
		log.Errorf("find gp daily failed: %s|%q", secucode, err)
		return err
	}
	if result != nil {
		return nil
	}

	result = orm.GPDailyMgr.NewGPDaily()
	result.Secucode = data.Secucode
	result.Name = data.Name
	result.Closing = data.Closing
	result.Rise = data.Rise
	result.PRise = data.PRise
	result.Turnover = data.Turnover
	result.Business = data.Business
	result.Liangbi = data.Liangbi
	result.MaxPrice = data.MaxPrice
	result.MinPrice = data.MinPrice
	result.Opening = data.Opening
	result.Market = data.Market
	result.Traded = data.Traded
	result.BookRatio = data.BookRatio
	result.CreateDate = createDate
	result.UpdateDate = time.Now().Unix()

	if _, err := result.Save(); err != nil {
		log.Errorf("save gp daily failed: %s|%q", secucode, err)
		return err
	}

	return nil
}

func getZeroTS() int64 {
	timeStr := time.Now().Format("2006-01-02")
	t, _ := time.ParseInLocation("2006-01-02", timeStr, time.Local)
	return t.Unix()
}
