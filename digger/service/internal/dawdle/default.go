package dawdle

import (
	"math"
	"sync"

	orm "git.ezbuy.me/ezbuy/corsair/digger/service/internal/model"
	trpc "git.ezbuy.me/ezbuy/corsair/digger/service/internal/rpc"
	"git.ezbuy.me/ezbuy/corsair/digger/service/internal/utils"
	log "github.com/Sirupsen/logrus"
)

const (
	shake_upper     float64 = 1.4 //不变
	shake_lower     float64 = 0.6
	one_ten         float64 = 0.1
	three_ten       float64 = 0.3
	one_fifth       float64 = 0.2
	two_fifth       float64 = 0.4
	three_fifth     float64 = 0.6
	four_fifth      float64 = 0.8
	two_rate        float64 = 2.0
	three_rate      float64 = 3.0
	four_rate       float64 = 4.0
	five_rate       float64 = 5.0
	TotalNumRatio   float64 = 45
	cumulantRatio   int     = 3
	cumulantPrice   int     = 5
	priceDiff       float64 = 1.25
	dateDiff        int64   = 60 * 86400
	accum           float64 = 15
	GPLongDecrease  int32   = 30
	GPShortDecrease int32   = 38
	ValueIndexTag   int32   = 80
)

type WeightRule struct { //权重规则
	IsNew              bool
	Price              *trpc.WeightUnit //float64
	Focus              *trpc.WeightUnit
	TotalNumRatio      *trpc.WeightUnit
	AvgFreesharesRatio *trpc.WeightUnit
	HoldRatioTotal     float64
	FreeholdRatioTotal float64
}

type WeightData struct { //权重数据
	calOnce            sync.Once
	weightOnce         sync.Once
	Secucode           string
	Price              []float64
	Focus              []string
	TotalNumRatio      []float64
	AvgFreesharesRatio []float64
	HoldRatioTotal     []float64
	FreeholdRatioTotal []float64
	Date               []int64
	GPDaily            *orm.GPDaily
	Weight             int32
	wr                 *trpc.WeightRule
}

func NewWeightData(secucode string) *WeightData {
	return &WeightData{
		Secucode:           secucode,
		Price:              make([]float64, 0, 8),
		Focus:              make([]string, 0, 8),
		TotalNumRatio:      make([]float64, 0, 8),
		AvgFreesharesRatio: make([]float64, 0, 8),
		HoldRatioTotal:     make([]float64, 0, 8),
		FreeholdRatioTotal: make([]float64, 0, 8),
		Date:               make([]int64, 0, 8),
		wr:                 defaultWeightRule(),
	}
}

func defaultWeightRule() *trpc.WeightRule {
	return &trpc.WeightRule{
		TotalNumRatio:  &trpc.WeightUnit{Value: 55},  //人均持股变化
		Focus:          &trpc.WeightUnit{Value: 7.5}, //集中度
		Price:          &trpc.WeightUnit{Value: 30},  //价格
		HoldRatioTotal: 7.5,
	}
}

func (wv *WeightData) Cal() *WeightData {
	wv.calOnce.Do(func() {
		wv.CalPrice()
		wv.CalFocus()
		wv.CalTotalNumRatio()
		wv.CalHoldRatioTotal()
		// wv.CalFreeholdRatioTotal()
	})

	return wv
}

func (wv *WeightData) CalPrice() {
	var max float64
	unit := wv.wr.Price

	// log.Infof("==>>TODO 212: %+v|%+v", max, wv.Price[0])
	for _, val := range wv.Price {
		// 是否为新股, 新股跌幅要大
		unit.SubNew = val <= 0
		wv.wr.IsNew = val <= 0
		if wv.wr.IsNew {
			break
		}
	}

	for _, val := range wv.Price {
		unit.Counter++
		if val <= 0 {
			continue
		}

		if max < val {
			max = val
		}
	}
	// log.Infof("==>>TODO 212: %+v|%+v", max, wv.Price[0])
	// 与最近一次价格比较--越跌越好
	rate := utils.GetRate(max, wv.GPDaily.Closing)
	if rate >= 0.5 {
		unit.Value = unit.Value * 1
	} else if rate >= 0.45 {
		unit.Value = unit.Value * 0.95
	} else if rate >= 0.40 {
		unit.Value = unit.Value * 0.9
	} else if rate >= 0.35 {
		unit.Value = unit.Value * 0.85
	} else if rate >= 0.30 {
		unit.Value = unit.Value * 0.75
	} else if rate >= 0.25 {
		unit.Value = unit.Value * 0.5
	} else if rate >= 0.20 {
		unit.Value = unit.Value * 0.3
	} else {
		unit.Value = unit.Value * 0.2
	}
}

func (wv *WeightData) CalFocus() {
	var weight float64
	for _, val := range wv.Focus {
		if val != "非常集中" && val != "较集中" {
			continue
		}

		if val == "非常集中" || val == "较集中" {
			weight += 0.25
		} else {
			break
		}
	}
	if weight > 1 {
		weight = 1
	}

	wv.wr.Focus.Value = wv.wr.Focus.Value * weight
}

func checkTotalNumRatioConsecutive(accum, current float64) bool {
	absCur := math.Abs(current)
	if absCur <= 4 { //股东变化幅度忽略值
		return true
	}

	if accum >= 0 && current > 0 {
		return true
	}

	if accum <= 0 && current < 0 {
		return true
	}
	return false
}

func (wv *WeightData) CalTotalNumRatio() { //+-,越小越好
	var rate float64
	unit := wv.wr.TotalNumRatio
	for idx, val := range wv.TotalNumRatio {
		if !checkTotalNumRatioConsecutive(unit.Accum, val) {
			// log.Infof("==>>TODO 231:%+v|%+v", wv.Secucode, val)
			break
		}
		// log.Infof("==>>TODO 232:%+v|%+v", wv.Secucode, val)
		absVal := math.Abs(val)
		if absVal <= 10 {
			continue
		}

		if len(unit.Indexes) > 4 {
			continue
		}

		unit.Indexes = append(unit.Indexes, int32(idx))
		// log.Infof("==>>TODO 233:%+v|%+v|%+v", wv.Secucode, unit.Indexes, len(unit.Indexes))

		// log.Infof("==>>TODO 234:%+v|%+v|%+v", wv.Secucode, unit.Indexes, len(unit.Indexes))
		// 计数器
		unit.Counter++
		// 累加值
		unit.Accum += val
	}

	absAccum := math.Abs(unit.Accum)
	// log.Infof("==>>TODO 235: %+v|%+v|%+v", wv.Secucode, absAccum, unit.Accum)
	if absAccum >= 50 {
		rate = 1
	} else if absAccum >= 45 {
		rate = 0.85
	} else if absAccum >= 40 {
		rate = 0.7
	} else if absAccum >= 35 {
		rate = 0.6
	} else if absAccum >= 30 {
		rate = 0.5
	} else if absAccum >= 25 {
		rate = 0.4
	} else {
		rate = 0.2
	}
	// log.Infof("==>>TODO 238:%+v|%+v|%+v", wv.Secucode, unit.Value, rate)
	// 股东人数变化率
	unit.Value = utils.Decimal(unit.Value * rate)
	// log.Infof("==>>TODO 239:%+v|%+v|%+v", wv.Secucode, unit.Accum, unit.Value)
}

func (wv *WeightData) CalAvgFreesharesRatio() { //+-,越大越好
	var pre, rate float64
	unit := wv.wr.AvgFreesharesRatio
	for idx, val := range wv.AvgFreesharesRatio {
		// 计数器
		unit.Counter++
		// 累加值
		unit.Accum += val
		if idx == 0 && val >= 35 {
			rate = 1
			break
		}
		// 计算连续值类型
		if val > 0 && (pre == 0 || pre > 0) {
			unit.Consecutive++
		} else {
			unit.Consecutive--
		}

		if unit.Accum >= 15 {
			break
		}

		pre = val
	}

	// log.Infof("==>>TODO 421:%+v", unit)
	// 最近是减少的
	if unit.Consecutive >= 2 && unit.Counter <= 4 {
		rate += 0.25
	}
	if unit.Accum >= 25 && unit.Consecutive <= 3 {
		rate += 0.75
	} else if unit.Accum >= 15 && unit.Consecutive <= 3 {
		rate += 0.6
	} else if unit.Accum >= 10 && unit.Consecutive <= 3 {
		rate += 0.45
	} else if unit.Accum >= 5 && unit.Consecutive <= 3 {
		rate += 0.3
	}
	unit.Value = utils.Decimal(unit.Value * math.Min(rate, 1))
}

func (wv *WeightData) CalHoldRatioTotal() {
	if len(wv.HoldRatioTotal) <= 0 {
		wv.wr.HoldRatioTotal = 0
		return
	}

	var rate float64
	var holdRatio float64
	if len(wv.HoldRatioTotal) > 0 {
		holdRatio = wv.HoldRatioTotal[0]
	}
	if len(wv.HoldRatioTotal) > 1 {
		holdRatio = math.Max(wv.HoldRatioTotal[0], wv.HoldRatioTotal[1])
	}

	if holdRatio > 70 {
		rate = 1
	} else if holdRatio > 65 {
		rate = 0.95
	} else if holdRatio > 60 {
		rate = 0.9
	} else if holdRatio > 55 {
		rate = 0.85
	} else if holdRatio > 50 {
		rate = 0.8
	} else if holdRatio > 45 {
		rate = 0.75
	} else {
		rate = 0.5
	}

	wv.wr.HoldRatioTotal = utils.Decimal(wv.wr.HoldRatioTotal * rate)
}

func (wv *WeightData) CalFreeholdRatioTotal() {
	if len(wv.HoldRatioTotal) <= 0 {
		return
	}

	var lastRatio, rate float64
	if len(wv.FreeholdRatioTotal) > 0 {
		lastRatio = wv.FreeholdRatioTotal[0]
	}
	if len(wv.FreeholdRatioTotal) > 1 {
		lastRatio = math.Max(lastRatio, wv.FreeholdRatioTotal[1])
	}

	if lastRatio > 40 {
		rate = 1
	} else if lastRatio > 25 {
		rate = 0.75
	} else if lastRatio > 15 {
		rate = 0.5
	} else {
		rate = 0.25
	}
	wv.wr.FreeholdRatioTotal = utils.Decimal(wv.wr.FreeholdRatioTotal * rate)
}

func (wv *WeightData) CalPriceDiff() {

}

func (wv *WeightData) GetTotalNumRatio() int32 {
	return int32(wv.wr.TotalNumRatio.Accum)
}

func (wv *WeightData) GetCumulantRatio() string {
	// num := wv.wr.TotalNumRatio.Counter
	// if len(wv.Date) < int(num) {
	// 	return utils.GetGDReduceRatio(wv.TotalNumRatio, "<-")
	// }

	// return utils.GetGDReduceRatio(wv.TotalNumRatio[:int(num)], "<-")

	list := make([]float64, 0, 4)
	indexes := wv.wr.TotalNumRatio.GetIndexes()
	for _, idx := range indexes {
		list = append(list, wv.TotalNumRatio[idx])
	}

	return utils.GetGDReduceRatio(list, "<-")
}

func (wv *WeightData) GetCumulantDate() string {
	// num := wv.wr.TotalNumRatio.Counter
	// if len(wv.Date) < int(num) {
	// 	return utils.GetDateStr(wv.Date, "<-")
	// }

	// return utils.GetDateStr(wv.Date[:int(num)], "<-")

	list := make([]int64, 0, 4)
	indexes := wv.wr.TotalNumRatio.GetIndexes()
	for _, idx := range indexes {
		list = append(list, wv.Date[idx])
	}

	return utils.GetDateStr(list, "<-")
}

func (wv *WeightData) GetCumulantFocus() string {
	// num := wv.wr.TotalNumRatio.Counter
	// if len(wv.Focus) < int(num) {
	// 	return utils.GetFocusStr(wv.Focus, "<-")
	// }

	// return utils.GetFocusStr(wv.Focus[:int(num)], "<-")

	list := make([]string, 0, 4)
	indexes := wv.wr.TotalNumRatio.GetIndexes()
	for _, idx := range indexes {
		list = append(list, wv.Focus[idx])
	}

	return utils.GetFocusStr(list, "<-")
}

func (wv *WeightData) GetCumulantPrice() string {
	list := make([]float64, 0, 4)
	indexes := wv.wr.TotalNumRatio.GetIndexes()
	for _, idx := range indexes {
		list = append(list, wv.Price[idx])
	}

	return utils.FloatSlice2Str(list, "<-")
}

func (wv *WeightData) GetWeight() int32 {
	wv.weightOnce.Do(func() {
		wr, weight := wv.wr, float64(0)
		// log.Infof("==>>TODO 450: %+v", wv)
		// log.Infof("==>>TODO 451: %+v|%+v|%+v|%+v|%+v", wv.Secucode, wr.Price.Value, wr.Focus.Value, wr.TotalNumRatio.Value, wr.HoldRatioTotal)
		// 判断与当前的差价率
		// if wv.Price[0] <= 0 || utils.GetRate(wv.GPDaily.Closing, wv.Price[0]) > 0.25 {
		// 	log.Infof("invalid price: %s|%f|%f", wv.Secucode, wv.GPDaily.Closing, wv.Price[0])
		// 	wv.Weight = 0
		// 	return
		// }
		// // 判断时间差值
		// // log.Infof("==>>TODO 452: %+v|%+v|%+v", wv.Secucode, time.Unix(wv.Date[0], 0).Format("2006-01-02"), wv.Date[0])
		// if time.Now().Unix()-wv.Date[0] > dateDiff {
		// 	log.Infof("invalid date: %s|%s", wv.Secucode, time.Unix(wv.Date[0], 0).Format("2006-01-02"))
		// 	wv.Weight = 0
		// 	return
		// }

		log.Infof("==>>TODO 458: %+v|%+v|%+v|%+v", wr.Price.Value, wr.Focus.Value, wr.TotalNumRatio.Value, wr.HoldRatioTotal)
		weight = wr.Price.Value + wr.Focus.Value + wr.TotalNumRatio.Value + wr.HoldRatioTotal
		wv.Weight = int32(weight)
		// log.Infof("==>>TODO 459: %+v", args ...interface{})
		// wv.Weight, _ = strconv.ParseFloat(fmt.Sprintf("%.1f", weight), 64)
		log.Infof("cal weight: %+v|%+v", wv.Secucode, wv.Weight)
	})

	return wv.Weight
}
