package dawdle

import (
	"fmt"
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
	recommendedLongOnce sync.Once
)

func RecommendedLongTicker() {
	tk := time.NewTicker(time.Second * 10)
	for range tk.C {
		if utils.CheckFuncValid(trpc.FunctionType_FunctionTypeRecommendedLong) {
			getRecommendedLongData()
		}
	}
}

func RecommendedLongOnce() {
	recommendedLongOnce.Do(func() {
		getRecommendedLongData()
	})
}

func RecommendedLongTmp(secucode string) {
	resutl, err := orm.GDLongLineMgr.FindOne(ezdb.M{
		"Secucode": secucode,
	}, "-CreateDate")
	if err != nil {
		log.Errorf("query long line failed: %s|%q", secucode, err)
		return
	}

	genLongLine(resutl)
}

func getRecommendedLongData() {
	sess, col := orm.GDLongLineMgr.GetCol()
	defer sess.Close()

	var gdll *orm.GDLongLine
	iter := col.Find(ezdb.M{"Disabled": false}).Batch(100).Prefetch(0.25).Iter()
	for iter.Next(&gdll) {
		genLongLine(gdll)
	}

	job.UpdateJob(trpc.FunctionType_FunctionTypeRecommendedLong)
}

func genLongLine(gdll *orm.GDLongLine) error {
	secucode := strings.Split(gdll.Secucode, ".")
	tm := time.Now().AddDate(0, -1, 0).Unix()
	query := ezdb.M{
		"Secucode":   secucode[1],
		"CreateDate": ezdb.M{"$gte": tm},
	}

	if gdll.ValueIndex < ValueIndexTag {
		return fmt.Errorf("invalid value index: %s|%d", gdll.Secucode, gdll.ValueIndex)
	}

	dailies, err := orm.GPDailyMgr.FindAll(query, "-CreateDate")
	if err != nil {
		log.Errorf("get daily failed: %s|%q", gdll.Secucode, err)
		return err
	}

	var max, current float64
	for idx, daily := range dailies {
		if idx == 0 {
			current = math.Min(daily.Closing, daily.MinPrice)
		}
		if daily.MaxPrice > max {
			max = daily.MaxPrice
		}
	}

	data := getGPRecommend(gdll.Secucode)
	data.GDDecrease = gdll.GDReduceRatio
	data.Decrease = utils.DecreasePercent(max, current)
	data.RMType = int32(trpc.RMType_RmTypeLong)
	if err := getLastDecrease(data); err != nil {
		log.Errorf("get last decrease failed: %s|%q", gdll.Secucode, err)
		return err
	}
	if err := applyGPRecommend(data); err != nil {
		log.Errorf("apply recommend failed: %s|%q", gdll.Secucode, err)
		return err
	}
	return nil
}

func applyGPRecommend(data *orm.GPRecommend) error {
	if data.Decrease > GPLongDecrease {
		data.State = int32(trpc.RMState_RMStateStarted)
	} else {
		data.State = int32(trpc.RMState_RMStatePrepared)
	}

	data.UpdateDate = time.Now().Unix()
	if _, err := data.Save(); err != nil {
		return err
	}

	return nil
}

func getGPRecommend(secucode string) *orm.GPRecommend {
	result, err := orm.GPRecommendMgr.FindOneBySecucodeDisabled(secucode, false)
	if err != nil || result == nil {
		result = orm.GPRecommendMgr.NewGPRecommend()
		result.Secucode = secucode
		result.CreateDate = time.Now().Unix()
	}
	return result
}

func getLastDecrease(data *orm.GPRecommend) error {
	secucode := utils.GetSecucode(data.Secucode)
	// log.Infof("==>>TODO 201: %+v|%+v", data.Secucode, secucode)
	tm := time.Now().AddDate(0, -2, 0).Unix()
	query := ezdb.M{
		"Secucode":   secucode,
		"CreateDate": ezdb.M{"$gte": tm},
	}

	results, err := orm.GPDailyMgr.FindAll(query, "-CreateDate")
	if err != nil {
		log.Errorf("query daily failed: %s|%q", secucode, err)
		return err
	}

	var createDate, counter int64
	for idx, result := range results {
		if idx == 0 {
			data.PresentPrice = math.Min(result.Closing, result.MinPrice)
		}

		if result.MaxPrice > data.MaxPrice {
			counter++
			data.MaxPrice = result.MaxPrice
			createDate = result.CreateDate
			data.MaxDay = utils.TS2Date(result.CreateDate)

		}
	}
	dateStr := time.Unix(createDate, 0).Format("2006-01-02")
	if counter == 1 {
		return fmt.Errorf("invalid data:%s", data.Secucode)
	}
	data.MaxPrice = utils.Decimal(data.MaxPrice)
	data.PresentPrice = utils.Decimal(data.PresentPrice)
	data.Decrease = utils.DecreasePercent(data.MaxPrice, data.PresentPrice)
	data.DecreaseDay = fmt.Sprintf("%d&%s", counter, dateStr)
	return nil
}
