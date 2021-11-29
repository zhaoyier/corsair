package dawdle

import (
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
	genShareholderOnce sync.Once
	shareholderData    []*WeightData
)

func GenLongLineTicker() {
	tk := time.NewTicker(time.Minute * 40)
	for range tk.C {
		weekday := time.Now().Weekday()
		nowHour := time.Now().Local().Hour()
		if weekday == time.Saturday || weekday == time.Sunday { //周
			continue
		}

		log.Infof("gen long line charging up: %d", nowHour)
		if nowHour >= 21 && nowHour < 22 {
			log.Infof("gen long line in progress: %d", nowHour)
			GenShareholder()
		}
	}
}

func GenLongLineOnce() {
	genShareholderOnce.Do(func() {
		GenShareholder()
	})
}

// 临时测试
func GenShareholderTmp(code string) error {
	start := time.Now().AddDate(0, -9, 0).Unix()
	getDawdleData(code, start)
	return nil
}

func GenShareholder() error {

	sess, col := orm.CNSecucodeMgr.GetCol()
	defer sess.Close()

	start := time.Now().AddDate(0, -9, 0).Unix()
	var secucode *orm.CNSecucode
	iter := col.Find(ezdb.M{}).Batch(100).Prefetch(0.25).Iter()
	for iter.Next(&secucode) {
		getDawdleData(secucode.Secucode, start)
	}

	// 更新任务
	job.UpdateJob("GenLongLine", "ok")

	return nil
}

func getDawdleData(secucode string, since int64) error {
	codes := strings.Split(secucode, ".")
	if len(codes) < 2 {
		log.Errorf("invalid secucode %s", secucode)
		return nil
	}
	secucode = codes[1] + "." + codes[0]
	query := ezdb.M{
		"Secucode": secucode,
	}
	gdResults, err := orm.GDRenshuMgr.Find(query, 20, 0, "-EndDate")
	if err != nil {
		log.Errorf("query gd renshu failed: %s|%q", secucode, err)
		return err
	}
	// log.Infof("==>>TODO 502: %+v", len(gdResults))
	dailyResult, err := orm.GPDailyMgr.FindOne(ezdb.M{"Secucode": codes[1]}, "-CreateDate")
	if err != nil {
		log.Errorf("query daily failed 11: %s|%q", secucode, err)
		return err
	}

	if len(gdResults) <= cumulantPrice {
		log.Warningf("maybe is new filter: %s|%d", secucode, len(gdResults))
		return nil
	}

	wv := NewWeightData(secucode)
	for _, r := range gdResults {
		wv.Price = append(wv.Price, r.Price)
		wv.Focus = append(wv.Focus, r.HoldFocus)
		wv.TotalNumRatio = append(wv.TotalNumRatio, r.TotalNumRatio)
		wv.AvgFreesharesRatio = append(wv.AvgFreesharesRatio, r.AvgFreesharesRatio)
		wv.Date = append(wv.Date, r.EndDate)
	}
	// log.Infof("==>>TODO 503: %+v", len(wv.Price))
	wv.RecentPrice = dailyResult.Closing

	if err := fillDawdleData(wv); err != nil {
		log.Errorf("fill file failed: %s|%q", secucode, err)
		return nil
	}

	return nil
}

func fillDawdleData(wv *WeightData) error {
	if cap(shareholderData) <= 0 {
		shareholderData = make([]*WeightData, 0, 8)
	}

	shareholderData = append(shareholderData, wv)

	if err := applyGPRecommend(wv); err != nil {
		log.Infof("apply recommend failed: %s|%+v", wv.Secucode, err)
	}

	return nil
}

// 记录数据库
func applyGPRecommend(wv *WeightData) error {
	enddate := time.Unix(wv.Date[0], 0).Format("2006-01-02")
	result, err := orm.GDHoldValueIndexMgr.FindOneBySecucodeEndDate(wv.Secucode, enddate)
	if err != nil && err != mgo.ErrNotFound {
		log.Errorf("apply recommend failed: %s|%s", wv.Secucode, err)
		return err
	}
	if result != nil {
		return nil
	}

	// log.Infof("==>>TODO 315:%+v|%+v|%+v", wv.Secucode, wv.Weight, wv.Cal().GetWeight() <= 50)
	result = orm.GDHoldValueIndexMgr.NewGDHoldValueIndex()
	result.EndDate = enddate
	result.Secucode = wv.Secucode
	result.ValueIndex = wv.Cal().GetWeight()
	result.CumulantPrice = intSlice2Str(wv.Price, "<-")
	result.CumulantFocus = utils.GetFocusStr(wv.Focus, "<-")
	result.CumulantDate = tmSlice2Str(wv.Date, "<-")
	result.CreateDate = time.Now().Unix()
	if result.ValueIndex <= 50 {
		return nil
	}

	if _, err := result.Save(); err != nil {
		log.Errorf("save recommend failed: %s|%q", wv.Secucode, err)
		return err
	}
	return nil
}
