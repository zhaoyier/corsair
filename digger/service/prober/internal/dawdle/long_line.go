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
	wv := NewWeightData(secucode)
	codes := strings.Split(secucode, ".")
	if len(codes) < 2 {
		log.Errorf("invalid secucode %s", secucode)
		return nil
	}

	dailyCode := codes[1]
	query := ezdb.M{
		"Secucode": secucode,
	}
	gdResults, err := orm.GDRenshuMgr.Find(query, 20, 0, "-EndDate")
	if err != nil {
		log.Errorf("query gd renshu failed: %s|%q", secucode, err)
		return err
	}
	// log.Infof("==>>TODO 502: %+v", len(gdResults))
	wv.GPDaily, err = orm.GPDailyMgr.FindOne(ezdb.M{"Secucode": dailyCode}, "-CreateDate")
	if err != nil {
		log.Errorf("query daily failed 11: %s|%q", secucode, err)
		return err
	}

	if len(gdResults) <= cumulantPrice {
		log.Warningf("maybe is new filter: %s|%d", secucode, len(gdResults))
		return nil
	}

	for _, r := range gdResults {
		wv.Price = append(wv.Price, r.Price)
		wv.Focus = append(wv.Focus, r.HoldFocus)
		wv.TotalNumRatio = append(wv.TotalNumRatio, r.TotalNumRatio)
		wv.AvgFreesharesRatio = append(wv.AvgFreesharesRatio, r.AvgFreesharesRatio)
		wv.Date = append(wv.Date, r.EndDate)
	}
	// log.Infof("==>>TODO 503: %+v", len(wv.Price))
	// wv.GPDaily = dailyResult

	if err := applyLongLine(wv); err != nil {
		log.Errorf("apply recommend failed: %s|%q", secucode, err)
		return nil
	}

	return nil
}

// 记录数据库
func applyLongLine(wv *WeightData) error {
	// log.Infof("==>>TODO 312:%+v|%+v|%+v", wv, nil, nil)
	enddate := time.Unix(wv.Date[0], 0).Format("2006-01-02")
	result, err := orm.GDLongLineMgr.FindOneBySecucodeEndDate(wv.Secucode, enddate)
	// log.Infof("==>>TODO 313:%+v|%+v|%+v", nil, result, err)
	if err != nil && err != mgo.ErrNotFound {
		log.Errorf("apply recommend failed: %s|%s", wv.Secucode, err)
		return err
	}
	if result != nil {
		return nil
	}

	// log.Infof("==>>TODO 315:%+v|%+v|%+v", wv.Secucode, wv.Weight, wv.Cal().GetWeight() <= 50)
	result = orm.GDLongLineMgr.NewGDLongLine()
	result.EndDate = enddate
	result.Secucode = wv.Secucode
	result.Name = wv.GPDaily.Name
	result.ValueIndex = wv.Cal().GetWeight()
	result.CumulantPrice = intSlice2Str(wv.Price, "<-")
	result.CumulantFocus = utils.GetFocusStr(wv.Focus, "<-")
	result.CumulantDate = utils.GetDateStr(wv.Date, "<-")
	result.GDReduceRatio = utils.GetGDReduceRatio(wv.TotalNumRatio, "&")
	result.CreateDate = time.Now().Unix()
	log.Infof("long line data:%+v|%+v", result.Name, result.ValueIndex)
	if result.ValueIndex <= 50 {
		return nil
	}

	if _, err := result.Save(); err != nil {
		log.Errorf("save recommend failed: %s|%q", wv.Secucode, err)
		return err
	}

	if err := disabledLongLine(result.Secucode, enddate); err != nil {
		log.Errorf("disabled long line failed: %s|%q", wv.Secucode, err)
		return err
	}
	return nil
}

func disabledLongLine(secucode, enddate string) error {
	query := ezdb.M{
		"Secucode": secucode,
	}
	results, err := orm.GDLongLineMgr.FindAll(query)
	if err != nil {
		return err
	}
	for _, result := range results {
		if result.EndDate == enddate {
			continue
		}
		result.Disabled = true
		if _, err := result.Save(); err != nil {
			log.Errorf("update long line failed: %s|%q", secucode, err)
		}
	}
	return nil
}
