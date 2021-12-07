package dawdle

import (
	"math"
	"sync"
	"time"

	"git.ezbuy.me/ezbuy/corsair/digger/service/internal/job"
	orm "git.ezbuy.me/ezbuy/corsair/digger/service/internal/model"
	trpc "git.ezbuy.me/ezbuy/corsair/digger/service/internal/rpc"
	"git.ezbuy.me/ezbuy/corsair/digger/service/internal/utils"
	log "github.com/Sirupsen/logrus"
	ezdb "github.com/ezbuy/ezorm/db"
	mgo "gopkg.in/mgo.v2"
)

var (
	genShortLineOnce sync.Once
)

func GenShortLineTicker() {
	GenShortLineData()

	// 更新任务
	job.UpdateJob(trpc.FunctionType_FunctionTypeShortLine)
}

func GenShortLineOnce() {
	genShortLineOnce.Do(func() {
		GenShortLineData()
	})
}

func GenShortLineTmp(secucode string) {
	getShortLineData(secucode)
}

func GenShortLineData() {
	sess, col := orm.CNSecucodeMgr.GetCol()
	defer sess.Close()

	var secucode *orm.CNSecucode
	iter := col.Find(ezdb.M{"Disabled": false}).Batch(100).Prefetch(0.25).Iter()
	for iter.Next(&secucode) {
		if err := getShortLineData(secucode.Secucode); err != nil {
			log.Infof("gen short line failed: %s|%q", secucode.Secucode, err)
		}
	}
}

//
func getShortLineData(secucode string) error {
	datets := utils.GetZeroTS()
	code := utils.GetSecucode(secucode)
	result, err := orm.GPShortLineMgr.FindOneBySecucodeCreateDate(secucode, datets)
	if err != nil && err != mgo.ErrNotFound {
		return err
	}
	if result == nil {
		result = orm.GPShortLineMgr.NewGPShortLine()
		result.Secucode = secucode
		result.CreateDate = datets
	}

	result.DecreaseTag = getDecreaseValue(code)
	result.MDecrease, _ = getShortLineDecrease(result, -30)
	result.TDecrease, _ = getShortLineDecrease(result, -60)
	result.UpdateDate = time.Now().Unix()
	decreaseTag := result.DecreaseTag - 10
	// log.Infof("==>>TODO 231: %+v|%+v", result, decreaseTag)
	if result.MDecrease < decreaseTag && result.TDecrease < decreaseTag {
		log.Errorf("invalid decreaseTag: %s|%d|%d", secucode, result.MDecrease, result.TDecrease)
		return nil
	}

	if _, err := result.Save(); err != nil {
		log.Errorf("apply short line failed: %s|%q", secucode, err)
		return err
	}

	if err := disabledShortLine(secucode); err != nil {
		log.Errorf("disabled short line failed: %s|%q", secucode, err)
		return err
	}
	return nil
}

func getShortLineDecrease(data *orm.GPShortLine, days int) (int32, error) {
	secucode := utils.GetSecucode(data.Secucode)
	// log.Infof("==>>TODO 201: %+v|%+v", data.Secucode, secucode)
	tm := time.Now().AddDate(0, 0, days).Unix()
	query := ezdb.M{
		"Secucode":   secucode,
		"CreateDate": ezdb.M{"$gte": tm},
	}

	results, err := orm.GPDailyMgr.FindAll(query, "-CreateDate")
	if err != nil {
		log.Errorf("query daily failed: %s|%q", secucode, err)
		return 0, err
	}

	for idx, result := range results {
		if idx == 0 {
			data.PresentPrice = math.Min(result.Closing, result.MinPrice)
		}

		if data.Name == "" {
			data.Name = result.Name
		}

		if result.MaxPrice > data.MaxPrice {
			data.MaxPrice = result.MaxPrice
			data.MaxPDay = utils.TS2Date(result.CreateDate)
		}

		if data.MinPrice == 0 || result.MinPrice < data.MinPrice {
			data.MinPrice = result.MinPrice
		}
	}

	data.MaxPrice = utils.Decimal(data.MaxPrice)
	return utils.DecreasePercent(data.MaxPrice, data.PresentPrice), nil
}

func getDecreaseValue(secucode string) int32 {
	query := ezdb.M{
		"Secucode": secucode,
	}
	result, err := orm.GPDailyMgr.FindOne(query, "-CreateDate")
	if err != nil {
		return GPShortDecrease
	}
	if result.Traded > int64(math.Pow10(11)*2) {
		return GPShortDecrease - 10
	} else if result.Traded > int64(math.Pow10(10)*5) {
		return GPShortDecrease - 5
	} else if result.Traded > int64(math.Pow10(10)) {
		return GPShortDecrease
	} else if result.Traded > int64(math.Pow10(9)) {
		return GPShortDecrease
	} else {
		return GPShortDecrease + 5
	}
}

func disabledShortLine(secucode string) error {
	query := ezdb.M{
		"Secucode":   secucode,
		"CreateDate": ezdb.M{"$lt": utils.GetZeroTS()},
		"Disabled":   true,
	}

	results, err := orm.GPShortLineMgr.FindAll(query)
	if err == mgo.ErrNotFound {
		return nil
	}
	if err != nil {
		return err
	}
	for _, result := range results {
		result.Disabled = true
		result.UpdateDate = time.Now().Unix()
		if _, err := result.Save(); err != nil {
			return err
		}
	}
	return nil
}
