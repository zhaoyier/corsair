package dawdle

import (
	"math"
	"strings"
	"sync"
	"time"

	"git.ezbuy.me/ezbuy/corsair/digger/service/internal/job"
	orm "git.ezbuy.me/ezbuy/corsair/digger/service/internal/model"
	"git.ezbuy.me/ezbuy/corsair/digger/service/internal/utils"
	log "github.com/Sirupsen/logrus"
	ezdb "github.com/ezbuy/ezorm/db"
	mgo "gopkg.in/mgo.v2"
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

func GenShortLineData() {
	sess, col := orm.CNSecucodeMgr.GetCol()
	defer sess.Close()

	start := time.Now().AddDate(0, 0, -10).Unix()
	var secucode *orm.CNSecucode
	iter := col.Find(ezdb.M{}).Batch(100).Prefetch(0.25).Iter()
	for iter.Next(&secucode) {
		if err := getShortLineData(secucode.Secucode, start); err != nil {
			log.Infof("gen short line failed: %s|%q", secucode.Secucode, err)
		}
	}
	// 更新任务
	job.UpdateJob("GenShortLine", "ok")
}

// 最近10日数据
func getShortLineData(secucode string, start int64) error {
	codes := strings.Split(secucode, ".")
	query := ezdb.M{"Secucode": codes[1]}

	results, err := orm.GPDailyMgr.Find(query, 10, 0, "-CreateDate")
	if err != nil {
		log.Errorf("query daily failed: %s|%q", secucode, err)
		return err
	}

	var max, current, min float64
	for idx, result := range results {
		if idx == 0 {
			current = math.Min(result.Closing, result.MinPrice)
		}
		if result.MaxPrice > max {
			max = result.MaxPrice
		}

		if min > result.MinPrice {
			min = result.MinPrice
		}
	}

	decrease := int32((max-current)/max) * 100

	if decrease < GPDecrease { //幅度太小的不做考虑
		return nil
	}

	createDate := utils.GetZeroTS()
	result, err := orm.GPStubIncidentMgr.FindOneBySecucodeCreateDate(secucode, createDate)
	if err != nil && err != mgo.ErrNotFound {
		log.Errorf("find gp daily failed: %s|%q", secucode, err)
		return err
	}

	if result != nil {
		return nil
	}

	result = orm.GPStubIncidentMgr.NewGPStubIncident()
	result.Secucode = secucode
	result.PriceDiff = int32(max - min)
	result.Decrease = decrease
	result.CreateDate = createDate
	result.UpdateDate = time.Now().Unix()
	if _, err := result.Save(); err != nil {
		log.Errorf("save gp daily failed: %s|%q", secucode, err)
		return err
	}

	return nil
}
