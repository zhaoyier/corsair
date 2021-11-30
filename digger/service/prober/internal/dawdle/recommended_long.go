package dawdle

import (
	"fmt"
	"math"
	"strings"
	"sync"
	"time"

	"git.ezbuy.me/ezbuy/corsair/digger/rpc/digger"
	"git.ezbuy.me/ezbuy/corsair/digger/service/internal/job"
	orm "git.ezbuy.me/ezbuy/corsair/digger/service/internal/model"
	log "github.com/Sirupsen/logrus"
	ezdb "github.com/ezbuy/ezorm/db"
)

var (
	recommendedLongOnce sync.Once
)

func RecommendedLongTicker() {
	tk := time.NewTicker(time.Minute * 90)
	for range tk.C {
		weekday := time.Now().Weekday()
		nowHour := time.Now().Local().Hour()
		if weekday == time.Saturday || weekday == time.Sunday { //周
			continue
		}

		log.Infof("get code list charging up: %d", nowHour)
		if nowHour >= 23 && nowHour < 24 {
			log.Infof("get code list in progress: %d", nowHour)
			getRecommendedLongData()
			job.UpdateJob("RecommendedLong", "ok")
		}
	}
}

func RecommendedLongOnce() {
	recommendedLongOnce.Do(func() {
		getRecommendedLongData()
	})
}

func getRecommendedLongData() {
	sess, col := orm.GDLongLineMgr.GetCol()
	defer sess.Close()

	var gdll *orm.GDLongLine
	iter := col.Find(ezdb.M{"Disabled": false}).Batch(100).Prefetch(0.25).Iter()
	for iter.Next(&gdll) {
		genLongLine(gdll)
	}
}

func genLongLine(gdll *orm.GDLongLine) error {
	secucode := strings.Split(gdll.Secucode, ".")
	tm := time.Now().AddDate(0, 1, 0).Unix()
	query := ezdb.M{
		"Secucode":   secucode[1],
		"CreateDate": ezdb.M{"$gte": tm},
	}

	if gdll.ValueIndex < 80 {
		return fmt.Errorf("invalid value index: %s|%d", gdll.Secucode, gdll.ValueIndex)
	}

	dailies, err := orm.GPDailyMgr.FindAll(query, "-CreateDate")
	if err != nil {
		log.Errorf("get daily failed: %s|%q", gdll.Secucode, err)
		return err
	}

	var max, current float64
	data := getGPRecommend(gdll.Secucode)
	for idx, daily := range dailies {
		if idx == 0 {
			current = math.Min(daily.Closing, daily.MinPrice)
		}
		if daily.MaxPrice > max {
			max = daily.MaxPrice
		}
	}

	data.Secucode = gdll.Secucode
	data.MonthDrop = int32((max-current)/max) * 100
	data.RmType = digger.RMType_RmTypeLong.GetValue()
	if err := applyGPRecommend(data); err != nil {
		log.Errorf("apply recommend failed: %s|%q", gdll.Secucode, err)
		return err
	}
	return nil
}

func applyGPRecommend(data *orm.GPRecommend) error {
	if data.MonthDrop > GPDecrease {
		data.State = int32(digger.RMState_RMStateStarted)
	} else {
		data.State = int32(digger.RMState_RMStatePrepared)
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
		result.CreateDate = time.Now().Unix()
	}
	return result
}
