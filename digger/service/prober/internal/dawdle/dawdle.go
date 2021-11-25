package dawdle

import (
	"fmt"
	"os"
	"strings"
	"sync"
	"time"

	orm "git.ezbuy.me/ezbuy/corsair/digger/service/internal/model"
	log "github.com/Sirupsen/logrus"
	ezdb "github.com/ezbuy/ezorm/db"
	"github.com/tealeg/xlsx"
)

var (
	genShareholderOnce sync.Once
)

func GenShareholderTicker() {
	tk := time.NewTicker(time.Minute * 90)
	for range tk.C {
		nowHour := time.Now().Local().Hour()
		if time.Now().Weekday() != time.Saturday { //周
			continue
		}

		if nowHour >= 22 && nowHour < 24 {
			GenShareholder()
		}
	}
}

func GenShareholderOnce() {
	genShareholderOnce.Do(func() {
		GenShareholder()
	})
}

func GenShareholderTmp() error {
	file, err := genDawdleTitle()
	if err != nil {
		log.Errorf("generate file failed: %s|%q", "", err)
		return err
	}

	start := time.Now().AddDate(0, -9, 0).Unix()
	log.Infof("==>>TODO 201: %+v", start)
	getDawdleData("SZ.300943", start, file)
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

	now := time.Now()
	date := fmt.Sprintf("%d-%d-%d", now.Year(), now.Month(), now.Day())
	dailyResult, err := orm.DailyMgr.FindOneBySecucodeEndDate(strings.Join(codes, "."), date)
	if err != nil {
		log.Errorf("query daily failed: %s|%q", secucode, err)
		return err
	}

	if len(gdResults) <= 6 {
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
		wv.Date = append(wv.Date, time.Unix(r.EndDate, 0).Format("2006-01-02"))
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
	titleRow.AddCell().SetString("价格")
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

	sheet := file.Sheets[0]
	row := sheet.AddRow()
	row.AddCell().SetString(wv.Secucode)
	row.AddCell().SetString(wv.Weight)
	row.AddCell().SetString(intSlice2Str(wv.Price, "<-"))
	row.AddCell().SetString(strings.Join(wv.Focus, "<-"))
	row.AddCell().SetString(strings.Join(wv.Date, "<-"))

	filename := fmt.Sprintf("%s.xlsx", time.Now().Format("2006-01-02"))
	if err := saveToFile(file, filename); err != nil {
		log.Errorf("save file failed: %s|%q", filename, err)
		return err
	}

	return nil
}

func saveToFile(file *xlsx.File, filename string) error {
	_, err := os.Stat("export")
	if err != nil {
		os.Mkdir("export", os.ModePerm)
	}

	if err := file.Save("export/" + filename); err != nil {
		log.Errorf("save file failed: %s|%q", filename, err)
		return err
	}

	return nil
}
