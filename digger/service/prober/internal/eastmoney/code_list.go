package eastmoney

import (
	"fmt"
	"time"

	"git.ezbuy.me/ezbuy/corsair/digger/service/internal/common/webapi"
	orm "git.ezbuy.me/ezbuy/corsair/digger/service/internal/model"
	ezdb "github.com/ezbuy/ezorm/db"
	mgo "gopkg.in/mgo.v2"
)

func GetCodeList() {
	sess, col := orm.CNSecucodeMgr.GetCol()
	defer sess.Close()

	// tk := time.NewTicker(time.Hour * 2)
	// for range tk.C {

	// }

	resp, inc := new(StockList), int32(1)
	for {
		if err := webapi.GetEastmoneyCode(inc, 80, resp); err != nil || resp.Data == nil {
			fmt.Printf("eastmoney get failed: %+v\n", err)
			return
		}
		inc++

		if err := updateCodeList(resp, col); err != nil {
			fmt.Printf("update code list failed: %+v\n", err)
		}
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
		prefix := val.Code[0:3]
		switch prefix {
		case "600":
		case "601":
		case "603":
		case "688":
			exchange = "SH"
		default:
			exchange = "SZ"
		}

		if err := applyCode(val.Code, exchange, val.Name, col); err != nil {
			fmt.Printf("apply failed: %+v|%+v\n", val, err)
		}
	}

	return nil

}

func applyCode(secu, exchange, name string, col *mgo.Collection) error {
	r, err := orm.CNSecucodeMgr.FindOneBySecucode(secu)
	if r != nil {
		return fmt.Errorf("unupdate %+v", err)
	}

	secucode := exchange + "." + secu
	query := ezdb.M{"secu": secucode}

	change := mgo.Change{
		Update: ezdb.M{
			"$set": ezdb.M{
				"SecurityCode": secu,
				"CreateDate":   time.Now().Unix(),
				"UpdateDate":   time.Now().Unix(),
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
