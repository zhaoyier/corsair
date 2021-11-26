package dawdle

import (
	"fmt"
	"os"
	"sort"
	"strings"
	"sync"
	"time"

	orm "git.ezbuy.me/ezbuy/corsair/digger/service/internal/model"
	log "github.com/Sirupsen/logrus"
	ezdb "github.com/ezbuy/ezorm/db"
	"github.com/tealeg/xlsx"
	mgo "gopkg.in/mgo.v2"
)

var (
	genShareholderOnce sync.Once
	shareholderData    []*WeightValue
)

func GenShareholderTicker() {
	tk := time.NewTicker(time.Minute * 90)
	for range tk.C {
		weekday := time.Now().Weekday()
		nowHour := time.Now().Local().Hour()
		if weekday == time.Saturday || weekday == time.Sunday { //周
			continue
		}

		log.Infof("gen share holder charging up: %d", nowHour)
		if nowHour >= 22 && nowHour < 24 {
			log.Infof("gen share holder in progress: %d", nowHour)
			GenShareholder()
			log.Infof("gen share holder completed: %d", nowHour)
		}
	}
}

func GenShareholderOnce() {
	genShareholderOnce.Do(func() {
		GenShareholder()
	})
}

func GenShareholderTmp(code string) error {
	file, err := genDawdleTitle()
	if err != nil {
		log.Errorf("generate file failed: %s|%q", "", err)
		return err
	}

	start := time.Now().AddDate(0, -9, 0).Unix()
	log.Infof("==>>TODO 201: %+v", start)
	// getDawdleData("SZ.003039", start, file)
	getDawdleData(code, start, file)
	saveToFile(file)
	return nil
}

func GenShareholder() error {
	file, err := genDawdleTitle()
	if err != nil {
		log.Errorf("generate file failed: %s|%q", "", err)
		return err
	}

	sess, col := orm.CNSecucodeMgr.GetCol()
	defer sess.Close()

	start := time.Now().AddDate(0, -9, 0).Unix()
	var secucode *orm.CNSecucode
	iter := col.Find(ezdb.M{}).Batch(100).Prefetch(0.25).Iter()
	for iter.Next(&secucode) {
		getDawdleData(secucode.Secucode, start, file)
	}

	saveToFile(file)

	return nil
}

func getDawdleData(secucode string, since int64, file *xlsx.File) error {
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

	// now := time.Now()
	// date := fmt.Sprintf("%d-%d-%d", now.Year(), now.Month(), now.Day())
	// dailyResult, err := orm.DailyMgr.FindOneBySecucodeEndDate(strings.Join(codes, "."), date)
	dailyResult, err := orm.DailyMgr.FindOne(ezdb.M{"Secucode": strings.Join(codes, ".")}, "-CreateDate")
	if err != nil {
		log.Errorf("query daily failed 11: %s|%q", secucode, err)
		return err
	}

	if len(gdResults) <= cumulantPrice {
		log.Warningf("maybe is new filter: %s|%d", secucode, len(gdResults))
		return nil
	}

	wv := NewWeightValue(secucode)
	for idx, r := range gdResults {
		if idx > 6 {
			continue
		}
		wv.Price = append(wv.Price, r.Price)
		wv.Focus = append(wv.Focus, r.HoldFocus)
		wv.TotalNumRatio = append(wv.TotalNumRatio, r.TotalNumRatio)
		wv.AvgFreesharesRatio = append(wv.AvgFreesharesRatio, r.AvgFreesharesRatio)
		wv.Date = append(wv.Date, r.EndDate)
	}

	wv.RecentPrice = dailyResult.Price

	if err := fillDawdleData(file, wv); err != nil {
		log.Errorf("fill file failed: %s|%q", secucode, err)
		return nil
	}

	return nil
}

func genDawdleTitle() (*xlsx.File, error) {
	file := xlsx.NewFile()
	sheet, err := file.AddSheet("gdrs")
	if err != nil {
		return nil, err
	}
	titleRow := sheet.AddRow()
	titleRow.AddCell().SetString("码")
	titleRow.AddCell().SetString("指数")
	titleRow.AddCell().SetString("最新价")
	titleRow.AddCell().SetString("参考价")
	titleRow.AddCell().SetString("集中度")
	titleRow.AddCell().SetString("日期")
	return file, nil
}

func fillDawdleData(file *xlsx.File, wv *WeightValue) error {
	if file == nil {
		return fmt.Errorf("sheet is nil")
	}

	if weight := wv.Cal().GetWeight(); weight < 70 {
		return fmt.Errorf("%s underweighting of stocks", wv.Secucode)
	}

	if cap(shareholderData) <= 0 {
		shareholderData = make([]*WeightValue, 0, 8)
	}

	shareholderData = append(shareholderData, wv)

	if err := applyGpRecommend(wv); err != nil {
		log.Infof("apply recommend failed: %s|%+v", wv.Secucode, err)
	}

	return nil
}

// 记录数据库
func applyGpRecommend(wv *WeightValue) error {
	// log.Infof("==>>TODO 311:%+v", wv.Secucode)
	enddate := time.Unix(wv.Date[0], 0).Format("2006-01-02")
	result, err := orm.GPRecommendMgr.FindOneBySecucodeEndDate(wv.Secucode, enddate)
	// result, err := orm.GPRecommendMgr.FindOneBySecucodeEndDate(wv.Secucode, "2021-11-12")
	// log.Infof("==>>TODO 312:%+v|%+v", result, err)
	// log.Infof("==>>TODO 313:%+v|%+v", err != nil, result != nil)
	if err != nil && err != mgo.ErrNotFound {
		log.Errorf("apply recommend failed: %s|%s", wv.Secucode, err)
		return err
	}
	if result != nil {
		return nil
	}

	// log.Infof("==>>TODO 315:%+v|%+v", result, err)
	result = orm.GPRecommendMgr.NewGPRecommend()
	result.Level = wv.Weight
	result.EndDate = enddate
	result.Secucode = wv.Secucode
	result.CumulantPrice = intSlice2Str(wv.Price, "<-")
	result.CumulantFocus = strings.Join(wv.Focus, "<-")
	result.CumulantDate = tmSlice2Str(wv.Date, "<-")
	result.CreateDate = time.Now().Unix()

	// log.Infof("==>>TODO 318:%+v", result)
	if _, err := result.Save(); err != nil {
		log.Errorf("save recommend failed: %s|%q", wv.Secucode, err)
		return err
	}
	return nil
}

func saveToFile(file *xlsx.File) error {
	filename := fmt.Sprintf("export/%s.xlsx", time.Now().Format("2006-01-02"))
	_, err := os.Stat("export")
	if err != nil {
		os.Mkdir("export", os.ModePerm)
	}

	sheet := file.Sheets[0]
	sort.Slice(shareholderData, func(i, j int) bool {
		return shareholderData[i].Weight > shareholderData[j].Weight
	})

	for _, val := range shareholderData {
		row := sheet.AddRow()
		row.AddCell().SetString(val.Secucode)
		row.AddCell().SetString(fmt.Sprintf("%.1f", val.Weight))
		row.AddCell().SetString(fmt.Sprintf("%.1f", val.RecentPrice))
		row.AddCell().SetString(intSlice2Str(val.Price, "<-"))
		row.AddCell().SetString(strings.Join(val.Focus, "<-"))
		row.AddCell().SetString(tmSlice2Str(val.Date, "<-"))
	}

	if err := file.Save(filename); err != nil {
		log.Errorf("save file failed: %s|%q", filename, err)
		return err
	}

	return nil
}
