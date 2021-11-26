package sina

import (
	"fmt"
	"strconv"
	"strings"
	"sync"
	"time"

	"git.ezbuy.me/ezbuy/corsair/digger/service/internal/common/webapi"
	orm "git.ezbuy.me/ezbuy/corsair/digger/service/internal/model"
	log "github.com/Sirupsen/logrus"
	ezdb "github.com/ezbuy/ezorm/db"
	mgo "gopkg.in/mgo.v2"
)

var (
	getDailyDataOnce sync.Once
)

func GetDailyDataTicker() {
	tk := time.NewTicker(time.Minute * 90)

	for range tk.C {
		weekday := time.Now().Weekday()
		nowHour := time.Now().Local().Hour()
		if weekday == time.Saturday || weekday == time.Sunday { //周
			continue
		}

		log.Infof("get daily data charging up: %d", nowHour)
		if nowHour >= 18 && nowHour < 20 {
			log.Infof("get daily data in progress: %d", nowHour)
			GetDailyData()
			log.Infof("get daily data completed: %d", nowHour)
		}
	}
}

func GetDailyDataOnce() {
	getDailyDataOnce.Do(func() {
		GetDailyData()
	})
}

func GetDailyDataTmp(secucode string) {
	code := strings.Replace(secucode, ".", "", -1)
	now := time.Now()
	date := fmt.Sprintf("%d-%d-%d", now.Year(), now.Month(), now.Day())
	result, err := webapi.GetSinaDayDetail(code)
	if err != nil {
		log.Errorf("get sina daily failed: %s|%q", secucode, err)
		return
	}
	// log.Infof("==>>TODO daily 313:%+v|%+v", result, date)
	if err := applyDaily(secucode, date, result); err != nil {
		log.Errorf("get sina daily failed: %s|%q", secucode, err)
	}
}

func GetDailyData() {
	now := time.Now()
	date := fmt.Sprintf("%d-%d-%d", now.Year(), now.Month(), now.Day())
	sess, col := orm.CNSecucodeMgr.GetCol()
	defer sess.Close()

	var secucode *orm.CNSecucode
	iter := col.Find(ezdb.M{}).Batch(100).Prefetch(0.25).Iter()
	for iter.Next(&secucode) {
		code := strings.Replace(secucode.Secucode, ".", "", -1)
		result, err := webapi.GetSinaDayDetail(code)
		if err != nil {
			log.Errorf("get sina daily failed: %s|%q", secucode.Secucode, err)
			continue
		}
		if err := applyDaily(secucode.Secucode, date, result); err != nil {
			log.Errorf("get sina daily failed: %s|%q", secucode.Secucode, err)
			continue
		}
	}
}

func applyDaily(secucode, date string, results []string) error {
	result, err := orm.DailyMgr.FindOneBySecucodeEndDate(secucode, date)
	if err != nil && err != mgo.ErrNotFound {
		log.Errorf("find gd renshu failed: %s|%s", secucode, date)
		return err
	}
	if result != nil {
		// log.Infof("==>>TODO 322: %+v|%+v", secucode, date)
		return nil
	}

	if len(results) <= 0 {
		log.Errorf("get daily data failed: %s|%s", secucode, results)
		return nil
	}

	result = orm.DailyMgr.NewDaily()
	result.EndDate = date
	result.Secucode = secucode
	result.CreateDate = time.Now().Unix()

	if len(results) > 4 {
		if val, err := strconv.ParseFloat(results[3], 64); err == nil {
			result.Price = Decimal(val)
		}
	}

	if len(results) > 5 {
		if val, err := strconv.ParseFloat(results[4], 64); err == nil {
			result.Highest = Decimal(val)
		}
	}

	if len(results) > 6 {
		if val, err := strconv.ParseFloat(results[5], 64); err == nil {
			result.Minimum = Decimal(val)
		}
	}
	// log.Infof("==>>TODO daily 337:%+v", result)
	if _, err := result.Save(); err != nil {
		log.Errorf("save gd renshu failed: %s|%q", secucode, err)
		return err
	}

	return nil
}

func Decimal(value float64) float64 {
	value, _ = strconv.ParseFloat(fmt.Sprintf("%.2f", value), 64)
	return value
}
