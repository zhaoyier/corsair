package dawdle

import (
	"fmt"
	"strings"
	"sync"
	"sync/atomic"
	"time"

	"git.ezbuy.me/ezbuy/corsair/digger/service/internal/job"
	orm "git.ezbuy.me/ezbuy/corsair/digger/service/internal/model"
	trpc "git.ezbuy.me/ezbuy/corsair/digger/service/internal/rpc"

	"git.ezbuy.me/ezbuy/corsair/digger/service/internal/utils"
	log "github.com/Sirupsen/logrus"
	ezdb "github.com/ezbuy/ezorm/db"
	"github.com/tealeg/xlsx"
)

var (
	offset           int   = 0
	_fourDay         int64 = 4 * 86400
	_million         int64 = 1000000
	StatFundFlowOnce sync.Once
	_path            = "./" + strings.Replace(utils.TS2Day(time.Now().Unix()), "-", "_", -1) + ".xlsx"
)

func GenStatFundFlowTicker() {
	getStatFundFlowData()
	// 更新任务
	job.UpdateJob(trpc.FunctionType_FunctionTypeStatFundFlow)
}

func GenStatFundFlowOnce() {
	StatFundFlowOnce.Do(func() {
		getStatFundFlowData()
	})
}

func GenStatFundFlowTmp(secucode string) {
	if err := getStatFundFlowItem(secucode, nil); err != nil {
		log.Infof("gen short line failed: %s|%q", secucode, err)
	}
}

func getStatFundFlowData() {
	sess, col := orm.CNSecucodeMgr.GetCol()
	defer sess.Close()

	file := generateFile()
	var count int32
	var secucode *orm.CNSecucode
	workCh := make(chan struct{}, 50)
	iter := col.Find(ezdb.M{"Disabled": false}).Batch(100).Prefetch(0.25).Iter()
	for iter.Next(&secucode) {
		secucode := secucode
		workCh <- struct{}{}
		go func(secucode *orm.CNSecucode) {
			atomic.AddInt32(&count, 1)
			if err := getStatFundFlowItem(secucode.Secucode, file); err != nil {
				log.Infof("gen short line failed: %s|%q", secucode.Secucode, err)
			}
			<-workCh
		}(secucode)

		if val := atomic.LoadInt32(&count); val%100 == 0 {
			log.Infof("==>>TODO 213: %+v", count)
		}
	}
	log.Infof("==>>TODO 215: %+v", count)
	file.Save(_path)
}

func getStatFundFlowItem(secucode string, file *xlsx.File) error {
	now := time.Now().Unix()
	// log.Infof("==>>TODO 222: %+v", secucode)
	if !strings.Contains(secucode, ".") {
		return nil
	}
	numcode := strings.Split(secucode, ".")[1]
	dailies, err := orm.GPDailyMgr.Find(ezdb.M{"Secucode": numcode}, 180, offset, "-_id") //TODOZ
	if err != nil {
		fmt.Printf("stat daily failed: %s|%q\n", secucode, err)
		return err
	}
	if len(dailies) <= 0 || dailies[0].Traded >= 20000*_million {
		return nil
	}
	daily := dailies[0]
	stat, err := orm.GPStatFundFlowMgr.FindOneBySecucode(secucode)
	if err != nil {
		stat = orm.GPStatFundFlowMgr.MustFindOneBySecucode(secucode)
	}

	// 统计数值
	flowes, err := orm.GPFundFlowMgr.Find(ezdb.M{"Secucode": secucode}, 21, offset, "-_id") //TODOZ
	if err != nil {
		fmt.Printf("stat fund flow failed: %s|%q\n", secucode, err)
		return err
	}
	if len(flowes) <= 2 || now-flowes[0].FundDate > _fourDay || now-daily.CreateDate > _fourDay { //TODOZ
		return nil
	}

	stat.Name = daily.Name
	stat.Secucode = secucode
	stat.Twenty = getFundFlowAmount(flowes, 20)
	stat.TwentyRatio = getFundFlowRatio(flowes, 20)
	stat.Ten = getFundFlowAmount(flowes, 10)
	stat.TenRatio = getFundFlowRatio(flowes, 10)
	stat.Five = getFundFlowAmount(flowes, 5)
	stat.FiveRatio = getFundFlowRatio(flowes, 5)
	stat.Three = getFundFlowAmount(flowes, 3)
	stat.ThreeRatio = getFundFlowRatio(flowes, 3)
	stat.Inflow = getInflowAmount(flowes, 3)
	stat.Traded = daily.Traded
	stat.MonthRise = getPriceIncrease(dailies, 30)
	stat.TermRise = getPriceIncrease(dailies, 180)
	stat.UpdateDate = time.Now().Unix()
	stat.Rising = genStatFundFlowRising(stat, dailies)
	go generateLineRow(file, stat, dailies)
	// log.Infof("==>>TODO 335: %+v", stat)
	if !checkStatParam(stat) {
		return nil
	}

	generateStatRow(file, stat)

	return nil
}

func generateStatRow(file *xlsx.File, stat *orm.GPStatFundFlow) {
	if file == nil {
		out := fmt.Sprintf("%s,%d,%s,", stat.Secucode, stat.Rising, stat.Inflow)
		out += fmt.Sprintf("%d,%d,%d,%d,", stat.Twenty, stat.Ten, stat.Five, stat.Three)
		out += utils.TS2Day(stat.UpdateDate)
		fmt.Println(out)
	} else {
		row := file.Sheets[0].AddRow()
		row.AddCell().Value = stat.Name
		row.AddCell().Value = stat.Secucode
		row.AddCell().Value = fmt.Sprintf("%d", stat.Rising)
		row.AddCell().Value = stat.Inflow
		row.AddCell().Value = fmt.Sprintf("%d", stat.Twenty)
		row.AddCell().Value = fmt.Sprintf("%d", stat.Ten)
		row.AddCell().Value = fmt.Sprintf("%d", stat.Five)
		row.AddCell().Value = fmt.Sprintf("%d", stat.Three)
		row.AddCell().Value = fmt.Sprintf("%d", stat.MonthRise)
		row.AddCell().Value = fmt.Sprintf("%d", stat.TermRise)
		row.AddCell().Value = utils.TS2Day(stat.UpdateDate)
	}

	if _, err := stat.Save(); err != nil {
		log.Errorf("stat fund flow failed: %s|%q", stat.Secucode, err)
	}
}

func generateLineRow(file *xlsx.File, stat *orm.GPStatFundFlow, dailies []*orm.GPDaily) {
	if file == nil {
		return
	}

	daily := dailies[0]
	ratio, max, min := getPriceDecrease(dailies, 7)
	if ratio <= -20 {
		r1 := file.Sheets[1].AddRow()
		r1.AddCell().Value = daily.Secucode
		r1.AddCell().Value = fmt.Sprintf("%d", ratio)
		r1.AddCell().Value = fmt.Sprintf("%d", stat.Three)
		r1.AddCell().Value = fmt.Sprintf("%d", stat.Five)
		r1.AddCell().Value = fmt.Sprintf("%d", min)
		r1.AddCell().Value = fmt.Sprintf("%d", max)
		r1.AddCell().Value = utils.TS2Day(stat.UpdateDate)
	}

	ratio, max, min = getPriceDecrease(dailies, 30)
	if ratio <= -40 {
		r2 := file.Sheets[2].AddRow()
		r2.AddCell().Value = daily.Secucode
		r2.AddCell().Value = fmt.Sprintf("%d", ratio)
		r2.AddCell().Value = fmt.Sprintf("%d", stat.Three)
		r2.AddCell().Value = fmt.Sprintf("%d", stat.Five)
		r2.AddCell().Value = fmt.Sprintf("%d", min)
		r2.AddCell().Value = fmt.Sprintf("%d", max)
		r2.AddCell().Value = utils.TS2Day(stat.UpdateDate)
	}

	ratio, max, min = getPriceDecrease(dailies, 180)
	if ratio <= -50 {
		r3 := file.Sheets[3].AddRow()
		r3.AddCell().Value = daily.Secucode
		r3.AddCell().Value = fmt.Sprintf("%d", ratio)
		r3.AddCell().Value = fmt.Sprintf("%d", stat.Three)
		r3.AddCell().Value = fmt.Sprintf("%d", stat.Five)
		r3.AddCell().Value = fmt.Sprintf("%d", min)
		r3.AddCell().Value = fmt.Sprintf("%d", max)
		r3.AddCell().Value = utils.TS2Day(stat.UpdateDate)
	}
}

func checkStatParam(stat *orm.GPStatFundFlow) bool {
	if stat.Inflow == "" {
		return false
	}
	if stat.Three < 10 || stat.Twenty > -50 {
		return false
	}

	if stat.Rising <= 0 {
		return false
	}
	return true
}

func getPriceIncrease(dailies []*orm.GPDaily, num int) int32 {
	if len(dailies) <= 0 {
		return 0
	}
	if len(dailies) > num {
		dailies = dailies[:num]
	}

	var maxPrice float64
	for _, daily := range dailies {
		if daily.MaxPrice > maxPrice {
			maxPrice = daily.MaxPrice
		}
	}

	result := int32(((dailies[0].Closing / maxPrice) - 1) * 100)
	return result
}

func getPriceDecrease(dailies []*orm.GPDaily, num int) (int32, int32, int32) {
	if len(dailies) <= 0 {
		return 0, 0, 0
	}
	if len(dailies) > num {
		dailies = dailies[:num]
	}

	var maxPrice, minPrice float64
	for idx, daily := range dailies {
		if idx == 0 {
			minPrice = daily.Closing
		}
		if daily.MaxPrice > maxPrice {
			maxPrice = daily.MaxPrice
		}
		if daily.MinPrice < minPrice {
			minPrice = daily.MinPrice
		}
	}

	result := int32(((dailies[0].Closing / maxPrice) - 1) * 100)
	return result, int32(maxPrice), int32(minPrice)
}

// 判断最近N日的涨幅
func getRecentlyRatio(dailies []*orm.GPDaily, num int) float64 {
	if len(dailies) <= 0 {
		return 100
	}
	// dailies = dailies[1:]
	first := dailies[0]
	var last *orm.GPDaily
	if len(dailies) > num {
		last = dailies[num]
	} else {
		last = dailies[len(dailies)-1]
	}

	diff := first.Closing - last.Closing
	result := utils.TruncateFloat((diff / last.Closing) * 100)
	return result
}

// 计算最近N日的流入, 单位: kw
func getInflowAmount(flowes []*orm.GPFundFlow, num int) string {
	// flowes = flowes[1:] //TODOZ
	results := make([]string, 0, num)
	for idx, flow := range flowes {
		if idx >= num {
			continue
		}
		if idx == 0 && flow.Inflow < 0 { //TODOZ 资金必须为正
			break
		}

		results = append(results, fmt.Sprintf("%.1f", float64(flow.Inflow)/float64(10*_million)))
	}
	if len(results) < num {
		return ""
	}

	return strings.Join(results, ";")
}

// 计算流入, 三五十天
func getFundFlowAmount(flowes []*orm.GPFundFlow, num int) int64 {
	var result int64
	// flowes = flowes[1:] //TODOZ
	for idx, item := range flowes {
		if idx >= num {
			continue
		}

		result += int64(item.Inflow)
	}
	return result / _million
}

// 计算最近流入比例
func getFundFlowRatio(flowes []*orm.GPFundFlow, num int) int64 {
	var result int64
	// flowes = flowes[1:] //TODOZ
	for idx, item := range flowes {
		if idx >= num {
			continue
		}

		result += int64(item.Inflow)
	}
	return result
}

// 计算流入趋势
func genStatFundFlowRising(stat *orm.GPStatFundFlow, dailies []*orm.GPDaily) int32 {
	// log.Infof("==>>TODO 350: %+v", stat)
	var rising int32
	if stat.Three < 0 {
		// log.Infof("==>>TODO 351: %+v", stat.Three)
		return 0
	}
	// daily := dailies[1] //TODOZ dailies[0]
	var totalRise, riseNum, limitUp int32
	// dailies = dailies[1:] //TODOZ
	for idx, daily := range dailies { //最近
		if idx >= 3 {
			continue
		}
		if daily.Rise >= 9.8 {
			limitUp++
		}
		riseNum++
		totalRise += int32(daily.Rise)

		// 高开低走
		dailyRise := utils.TruncateFloat((daily.MaxPrice - daily.Closing) * 100 / daily.Closing)
		if dailyRise >= 4 {
			// log.Infof("==>>TODO 352: %+v", dailyRise)
			return 0
		}
	}

	if ratio := getRecentlyRatio(dailies, 5); ratio <= 3 {
		// log.Infof("==>>TODO 355: %+v", ratio)
		return 0
	} else if ratio >= 25 {
		rising -= 30
	}

	// 趋势不对
	if riseNum <= 0 || (totalRise/riseNum) < 0 {
		// log.Infof("==>>TODO 356: %+v|%+v", riseNum, totalRise)
		return 0
	}
	if limitUp >= 1 {
		rising -= limitUp * 15
	}

	if stat.Twenty < 0 && stat.Ten < 0 && stat.Five > 0 && stat.Three > 0 {
		rising += 40
		if stat.Twenty < -30 && stat.Ten < -10 && (stat.Five > 10 || stat.Three > 20) {
			rising += 20
		}
	}

	max := float64((stat.Five+stat.Three-stat.Twenty)*_million) / float64(stat.Traded)
	flow := float64((stat.Five+stat.Three+stat.Twenty)*_million) / float64(stat.Traded)
	if flow > max {
		max = flow
	}
	if max <= 0.1 {
		max *= float64(20)
	} else {
		max = 1
	}
	if max > 1 {
		max = 1
	}

	rising += int32(max * 40)
	return rising
}

func generateFile() *xlsx.File {
	file := xlsx.NewFile()
	s1, err := file.AddSheet("sheet1")
	if err != nil {
		panic(err)
	}
	s2, _ := file.AddSheet("周线")
	s3, _ := file.AddSheet("月线")
	s4, _ := file.AddSheet("半年线")
	// style := xlsx.NewStyle()
	// style.Font.Name = "Georgia"
	// style.Font.Size = 12
	// style.Font.Bold = true
	// style.Alignment.WrapText = true
	trS1 := s1.AddRow()
	trS1.AddCell().Value = "名称"
	trS1.AddCell().Value = "代码"
	trS1.AddCell().Value = "趋势"
	trS1.AddCell().Value = "资金"
	trS1.AddCell().Value = "二十日"
	trS1.AddCell().Value = "十日"
	trS1.AddCell().Value = "五日"
	trS1.AddCell().Value = "三日"
	trS1.AddCell().Value = "月线"
	trS1.AddCell().Value = "半年线"
	trS1.AddCell().Value = "日期"

	trS2 := s2.AddRow()
	trS2.AddCell().Value = "代码"
	trS2.AddCell().Value = "周线"
	trS2.AddCell().Value = "三日"
	trS2.AddCell().Value = "五日"
	trS2.AddCell().Value = "最低"
	trS2.AddCell().Value = "最高"
	trS2.AddCell().Value = "日期"

	trS3 := s3.AddRow()
	trS3.AddCell().Value = "代码"
	trS3.AddCell().Value = "月线"
	trS3.AddCell().Value = "三日"
	trS3.AddCell().Value = "五日"
	trS3.AddCell().Value = "最低"
	trS3.AddCell().Value = "最高"
	trS3.AddCell().Value = "日期"

	trS4 := s4.AddRow()
	trS4.AddCell().Value = "代码"
	trS4.AddCell().Value = "月线"
	trS4.AddCell().Value = "三日"
	trS4.AddCell().Value = "五日"
	trS4.AddCell().Value = "最低"
	trS4.AddCell().Value = "最高"
	trS4.AddCell().Value = "日期"

	go func() {
		file.Save(_path)
	}()

	return file
}
