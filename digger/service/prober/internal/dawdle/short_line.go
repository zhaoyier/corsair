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
	"git.ezbuy.me/ezbuy/corsair/digger/service/internal/utils"
	log "github.com/Sirupsen/logrus"
	ezdb "github.com/ezbuy/ezorm/db"
)

var (
	genShortLineOnce sync.Once
	// shareholderData    []*WeightData
)

func GenShortLineTicker() {
	tk := time.NewTicker(time.Minute * 40)
	for range tk.C {
		weekday := time.Now().Weekday()
		nowHour := time.Now().Local().Hour()
		if weekday == time.Saturday || weekday == time.Sunday { //周
			continue
		}

		log.Infof("gen share holder charging up: %d", nowHour)
		if nowHour >= 21 && nowHour < 22 {
			log.Infof("gen share holder in progress: %d", nowHour)
			GenShortLineData()
		}
	}
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
	// 更新任务
	job.UpdateJob("GenShortLine")
}

// 最近10日数据
func getShortLineData(secucode string) error {
	codes := strings.Split(secucode, ".")
	data := getGPRecommend(secucode)
	var half, onem, twom string
	data.HDecrease, half = getLastDecrease(codes[1], -15)
	data.MDecrease, onem = getLastDecrease(codes[1], -30)
	data.TDecrease, twom = getLastDecrease(codes[1], -60)
	if data.HDecrease < GPShortDecrease && data.MDecrease < GPShortDecrease && data.TDecrease < GPShortDecrease {
		return nil
	}
	data.RMPrice = calRecommendPrice(codes[1], -15, -30)
	data.DecreaseDay = fmt.Sprintf("%s|%s|%s", half, onem, twom)
	data.RMType = int32(digger.RMType_RmTypeShort)
	//

	if err := applyGPRecommend(data); err != nil {
		log.Errorf("apply recommend failed: %s|%q", secucode, err)
		return err
	}
	return nil
}

func getLastDecrease(secucode string, day int) (int32, string) {
	tm := time.Now().AddDate(0, 0, day).Unix()
	query := ezdb.M{
		"Secucode":   secucode,
		"CreateDate": ezdb.M{"$gte": tm},
	}

	results, err := orm.GPDailyMgr.FindAll(query, "-CreateDate")
	if err != nil {
		log.Errorf("query daily failed: %s|%q", secucode, err)
		return 0, ""
	}

	var counter int
	var createDate int64
	var max, current float64
	now := time.Now().Unix()
	for _, result := range results {
		if now-result.CreateDate <= int64(86400*1.5) {
			current = math.Min(result.Closing, result.MinPrice)
		}

		if result.MaxPrice > max {
			counter++
			max = result.MaxPrice
			createDate = result.CreateDate
		}
	}
	dateStr := time.Unix(createDate, 0).Format("2006-01-02")
	if counter == 1 {
		return 0, ""
	}

	return utils.DecreasePercent(max, current), fmt.Sprintf("%d&%s", counter, dateStr)
}

func calRecommendPrice(secucode string, latest, farthest int) string {
	lasttm := time.Now().AddDate(0, 0, latest).Unix()
	fasttm := time.Now().AddDate(0, 0, farthest).Unix()
	query := ezdb.M{
		"Secucode":   secucode,
		"CreateDate": ezdb.M{"$gte": fasttm},
	}

	results, err := orm.GPDailyMgr.FindAll(query, "-CreateDate")
	if err != nil {
		log.Errorf("query daily failed: %s|%q", secucode, err)
		return ""
	}

	var lastmax, fastmax, counter float64
	for _, result := range results {
		if result.CreateDate > lasttm && result.MaxPrice > lastmax {
			counter++
			lastmax = result.MaxPrice
		}

		if result.MaxPrice > fastmax {
			fastmax = result.MaxPrice
		}
	}

	if counter == 1 {
		return ""
	}

	max := math.Max(lastmax, fastmax)
	return fmt.Sprintf("%.1f-%.1f", math.Floor(max*0.55), math.Ceil(max*0.6))
}
