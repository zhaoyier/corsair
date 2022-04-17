package dawdle

import (
	"math"
	"strings"
	"sync"
	"time"

	"git.ezbuy.me/ezbuy/corsair/digger/service/internal/job"
	orm "git.ezbuy.me/ezbuy/corsair/digger/service/internal/model"
	trpc "git.ezbuy.me/ezbuy/corsair/digger/service/internal/rpc"
	"git.ezbuy.me/ezbuy/corsair/digger/service/internal/utils"
	log "github.com/Sirupsen/logrus"
	ezdb "github.com/ezbuy/ezorm/db"
)

var (
	WaterfallOnce sync.Once
)

func GenWaterfallTicker() {
	genWaterfallData()
	// 更新任务
	job.UpdateJob(trpc.FunctionType_FunctionTypeWaterfall)
}

func GenWaterfallOnce() {
	zhouqiOnce.Do(func() {
		genWaterfallData()
	})
}

func GenWaterfallTmp(secucode string) {
	if err := genWaterfallItem(secucode); err != nil {
		log.Infof("gen short line failed: %s|%q", secucode, err)
	}
}

func genWaterfallData() {
	sess, col := orm.CNSecucodeMgr.GetCol()
	defer sess.Close()

	var secucode *orm.CNSecucode
	iter := col.Find(ezdb.M{"Disabled": false}).Batch(100).Prefetch(0.25).Iter()
	for iter.Next(&secucode) {
		if err := genWaterfallItem(secucode.Secucode); err != nil {
			log.Infof("gen short line failed: %s|%q", secucode.Secucode, err)
		}
	}
}

func genWaterfallItem(secucode string) error {
	code := utils.GetSecucode(secucode)
	tm := time.Now().AddDate(0, 0, -7).Unix()
	query := ezdb.M{
		"Secucode":   code,
		"CreateDate": ezdb.M{"$gte": tm},
	}

	results, err := orm.GPDailyMgr.FindAll(query, "-CreateDate")
	// fmt.Printf("==>>TODO 111: %+v|%+v\n", err, len(results))
	if err != nil {
		log.Errorf("query daily failed: %s|%q", secucode, err)
		return err
	}

	data := orm.GPWaterfallLineMgr.MustFindOneBySecucode(secucode)
	if data.CreateDate <= 0 {
		data.CreateDate = time.Now().Unix()
	}
	data.MaxPrice = 0
	data.MinPrice = 0

	// fmt.Printf("==>>TODO 220:%+v|%+v\n", 0, data)

	for idx, result := range results {
		// fmt.Printf("==>>TODO 221:%+v|%+v\n", idx, result)
		if idx == 0 {
			data.Name = result.Name
			data.PresentPrice = result.Closing
		}

		if data.Name == "" {
			data.Name = result.Name
		}

		if result.MaxPrice > data.MaxPrice {
			// fmt.Printf("==>>TODO 224:%+v\n", result)
			data.MaxPrice = result.MaxPrice
			data.MaxPDay = result.CreateDate
		}
		if data.MinPrice == 0 || (result.MinPrice != 0 && result.MinPrice < data.MinPrice) {
			// fmt.Printf("==>>TODO 225:%+v\n", result)
			data.MinPrice = result.MinPrice
		}
	}
	data.MaxPrice = utils.Decimal(data.MaxPrice)
	data.UpdateDate = time.Now().Unix()
	if strings.HasPrefix(strings.ToLower(data.Name), "st") {
		return nil
	}
	if strings.HasPrefix(strings.ToLower(data.Name), "*st") {
		return nil
	}

	if data.MaxPrice <= 0 || data.MinPrice <= 0 || data.PresentPrice <= 0 {
		return nil
	}

	// fmt.Printf("==>>TODO 115: %+v|%+v|%+v\n", data.MinPrice, data.MaxPrice, data.PresentPrice)
	// fmt.Printf("==>>TODO 1151: %+v|%+v|%+v\n", data.MaxPrice, data.PresentPrice, data.MaxPrice > data.PresentPrice)
	if data.MaxPrice > data.PresentPrice {
		rate := utils.Decimal((data.PresentPrice - data.MaxPrice) / data.MaxPrice)
		data.Decrease = int32(math.Ceil(rate * 100))
	} else {
		rate := utils.Decimal((data.PresentPrice - data.MinPrice) / data.MinPrice)
		// fmt.Printf("==>>TODO 116: %+v|%+v|%+v\n", rate, nil, nil)
		data.Decrease = int32(math.Ceil(rate * 100))
		// fmt.Printf("==>>TODO 117: %+v|%+v|%+v\n", data.Decrease, nil, nil)
	}

	if _, err := data.Save(); err != nil {
		return err
	}
	return nil
}

func getShortLineDecrease2(data *orm.GPShortLine, days int) (int32, error) {
	secucode := utils.GetSecucode(data.Secucode)
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

	// var minPrice float64
	for idx, result := range results {
		// log.Infof("==>>TODO 213: %+v", result)
		if idx == 0 {
			data.PresentPrice = result.Closing
			// minPrice = result.MinPrice
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
	return utils.DecreasePercent(data.MaxPrice, data.PresentPrice), nil
}
