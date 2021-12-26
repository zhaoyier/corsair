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
	disabledShortLine(secucode)
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

		if err := disabledShortLine(secucode.Secucode); err != nil {
			log.Errorf("disabled short line failed: %s|%q", secucode, err)
		}
	}
}

//
func getShortLineData(secucode string) error {
	datets := utils.GetZeroTS()
	code := utils.GetSecucode(secucode)
	// log.Infof("==>>TODO 231: %+v|%+v", secucode, datets)
	result, err := orm.GPShortLineMgr.FindOneBySecucodeCreateDate(secucode, datets)
	if err != nil && err != mgo.ErrNotFound {
		return err
	}
	if result == nil {
		// log.Infof("==>>TODO 232: %+v|%+v", secucode, datets)
		result = orm.GPShortLineMgr.NewGPShortLine()
		result.Secucode = secucode
		result.CreateDate = datets
	}

	result.DecreaseTag = getDecreaseValue(code)
	result.MDecrease, _ = getShortLineDecrease(result, -30)
	result.TDecrease, _ = getShortLineDecrease(result, -60)
	result.UpdateDate = time.Now().Unix()
	decreaseTag := result.DecreaseTag - 10
	// log.Infof("==>>TODO 235: %+v|%+v", result, decreaseTag)
	if result.MDecrease < decreaseTag && result.TDecrease < decreaseTag {
		log.Errorf("invalid decreaseTag: %s|%d|%d", secucode, result.MDecrease, result.TDecrease)
		return nil
	}

	if _, err := result.Save(); err != nil {
		log.Errorf("apply short line failed: %s|%q", secucode, err)
		return err
	}
	return nil
}

func getShortLineDecrease(data *orm.GPShortLine, days int) (int32, error) {
	secucode := utils.GetSecucode(data.Secucode)
	// log.Infof("==>>TODO 251: %+v|%+v", data, secucode)
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

	var minPrice float64
	for idx, result := range results {
		// log.Infof("==>>TODO 213: %+v", result)
		if idx == 0 {
			data.PresentPrice = result.Closing
			minPrice = result.MinPrice
		}

		if data.Name == "" {
			data.Name = result.Name
		}

		if result.MaxPrice > data.MaxPrice {
			data.MaxPrice = result.MaxPrice
			data.MaxPDay = result.CreateDate
		}
		// log.Infof("==>>TODO 256: %+v|%+v", data.MinPrice, result.MinPrice)
		if data.MinPrice == 0 || (result.MinPrice != 0 && result.MinPrice < data.MinPrice) {
			data.MinPrice = result.MinPrice
			// log.Infof("==>>TODO 257: %+v|%+v", data.MinPrice, result.MinPrice)
		}
	}
	// log.Infof("==>>TODO 258: %+v", data)
	data.MaxPrice = utils.Decimal(data.MaxPrice)
	return utils.DecreasePercent(data.MaxPrice, minPrice), nil
}

func getDecreaseValue(secucode string) int32 {
	query := ezdb.M{
		"Secucode": secucode,
	}
	results, err := orm.GPDailyMgr.Find(query, 2, 0, "-CreateDate")
	if err != nil {
		return GPShortDecrease
	}

	result, sdecrease := results[0], GPShortDecrease

	if result.Traded > int64(math.Pow10(10)*5) { //>= 500
		sdecrease = GPShortDecrease - 5
	} else if result.Traded > int64(math.Pow10(10)) { // >= 100
		sdecrease = GPShortDecrease
	} else if result.Traded > int64(math.Pow10(9)) { // >= 10
		sdecrease = GPShortDecrease + 2
	} else {
		sdecrease = GPShortDecrease + 3
	}

	return sdecrease
}

func disabledShortLine(secucode string) error {
	query := ezdb.M{
		"Secucode":   secucode,
		"CreateDate": ezdb.M{"$lt": utils.GetZeroTS()},
		"Disabled":   false,
	}

	results, err := orm.GPShortLineMgr.FindAll(query)
	// log.Infof("==>>TODO 242: %+v|%+v", err, len(results))
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
