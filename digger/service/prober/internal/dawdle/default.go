package dawdle

import (
	"fmt"
	"math"
	"strconv"
	"strings"
	"time"

	trpc "git.ezbuy.me/ezbuy/corsair/digger/rpc/digger"
	log "github.com/Sirupsen/logrus"
)

const (
	shake_upper   float64 = 1.4 //不变
	shake_lower   float64 = 0.6
	one_ten       float64 = 0.1
	three_ten     float64 = 0.3
	one_fifth     float64 = 0.2
	two_fifth     float64 = 0.4
	three_fifth   float64 = 0.6
	four_fifth    float64 = 0.8
	two_rate      float64 = 2.0
	three_rate    float64 = 3.0
	four_rate     float64 = 4.0
	five_rate     float64 = 5.0
	TotalNumRatio float64 = 45
	cumulantRatio int     = 3
	cumulantPrice int     = 5
	priceDiff     float64 = 1.25
	dateDiff      int64   = 10 * 86400
	accum         float64 = 15
)

// type WeightUnit struct { //权重单元
// 	Value       float64 //权重值
// 	Accum       float64 //连续累加值
// 	Consecutive int32   //连续计数
// 	Counter     int     //计数器
// }

type WeightRule struct { //权重规则
	Price              float64
	Focus              float64
	TotalNumRatio      *trpc.WeightUnit
	AvgFreesharesRatio *trpc.WeightUnit
	HoldRatioTotal     float64
	FreeholdRatioTotal float64
}

type WeightData struct { //权重数据
	Secucode           string
	Price              []float64
	Focus              []string
	TotalNumRatio      []float64
	AvgFreesharesRatio []float64
	HoldRatioTotal     []float64
	FreeholdRatioTotal []float64
	Date               []int64
	RecentPrice        float64
	Weight             int32
	wr                 *WeightRule
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

func defaultWeightRule() *WeightRule {
	return &WeightRule{
		TotalNumRatio:      &trpc.WeightUnit{Value: 20},
		AvgFreesharesRatio: &trpc.WeightUnit{Value: 20},
		Focus:              20,
		Price:              20,
		HoldRatioTotal:     10,
		FreeholdRatioTotal: 10,
	}
}

func (wv *WeightData) Cal() *WeightData {
	wv.CalPrice()
	wv.CalFocus()
	wv.CalTotalNumRatio()
	wv.CalAvgFreesharesRatio()
	wv.CalHoldRatioTotal()
	wv.CalFreeholdRatioTotal()
	return wv
}

func (wv *WeightData) CalPrice() {
	max, min, prev := wv.Price[0], float64(0), wv.Price[0]
	if len(wv.Price) > cumulantPrice {
		wv.Price = wv.Price[0:cumulantPrice]
	}

	for idx, val := range wv.Price {
		// log.Infof("==>>TODO Price 401:%s|%+v|%f|%+v", wv.Secucode, idx, val, val <= 0)
		if idx > cumulantPrice && val <= 0 {
			wv.wr.Price = 0
			return
		}

		// 价格跳动太大就舍弃了
		// log.Infof("==>>TODO Price 402:%s|%+v|%f|%f", wv.Secucode, idx, math.Pow(1.1, float64(idx)), priceDiff*math.Pow(1.1, float64(idx)))
		if (math.Abs(prev-val) / val) > priceDiff*math.Pow(1.1, float64(idx)) {
			wv.wr.Price = 0
			return
		}

		if max < val {
			max = val
		}

		if min > val {
			min = val
		}
	}

	// 与最近一次价格比较--越跌越好
	rate := Decimal((max - wv.Price[0]) / max)
	if rate <= one_ten {
		wv.wr.Price = 0
	} else if rate < 0.3 {
		wv.wr.Price = wv.wr.Price * 0.5
	} else if rate < 0.4 {
		wv.wr.Price = wv.wr.Price * 0.8
	} else if rate < 0.5 {
		wv.wr.Price = wv.wr.Price * 0.9
	}

	//100% 暂时不考虑--越涨越差
	rate2 := Decimal((wv.RecentPrice - min) / min)
	if rate2 >= 1 {
		wv.wr.Price = 0
	} else if rate2 > 0.5 {
		wv.wr.Price = wv.wr.Price * 0.5
	}
}

func (wv *WeightData) CalFocus() {
	var weight float64
	for idx, val := range wv.Focus {
		// log.Infof("==>>401: %+v|%+v|%+v", idx, val, val == "非常集中")
		if idx <= cumulantRatio && (val == "非常集中" || val == "较集中") {
			weight += 0.25
		}
	}
	wv.wr.Focus = wv.wr.Focus * weight
	// log.Infof("==>>405: %+v|%+v", wv.wr.Focus, weight)
}

func (wv *WeightData) CalTotalNumRatio() { //+-,越小越好
	var pre, rate float64
	unit := wv.wr.TotalNumRatio
	for _, val := range wv.TotalNumRatio {
		// 计数器
		unit.Counter++
		// 累加值
		unit.Accum += val
		// 计算连续值类型
		if val < 0 && (pre == 0 || pre < 0) {
			unit.Consecutive++
		} else {
			unit.Consecutive--
		}

		if unit.Accum <= -15 {
			break
		}

		pre = val
	}

	// 最近是减少的
	if unit.Consecutive >= 2 && unit.Counter <= 4 {
		rate += 0.25
	}
	//
	if unit.Accum <= -1*25 && unit.Consecutive <= 3 {
		rate += 0.75
	} else if unit.Accum <= -1*15 && unit.Consecutive <= 3 {
		rate += 0.6
	} else if unit.Accum <= -1*10 && unit.Consecutive <= 3 {
		rate += 0.45
	} else if unit.Accum <= -1*5 && unit.Consecutive <= 3 {
		rate += 0.3
	}

	unit.Value = Decimal(unit.Value * rate)
}

func (wv *WeightData) CalAvgFreesharesRatio() { //+-,越大越好
	var pre, rate float64
	unit := wv.wr.AvgFreesharesRatio
	for _, val := range wv.AvgFreesharesRatio {
		// 计数器
		unit.Counter++
		// 累加值
		unit.Accum += val
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
	unit.Value = Decimal(unit.Value * rate)
}

func (wv *WeightData) CalHoldRatioTotal() {
	if len(wv.HoldRatioTotal) <= 0 {
		return
	}

	currentHold := wv.HoldRatioTotal[0]
	if currentHold > 60 {
		wv.wr.HoldRatioTotal = 10
	} else if currentHold > 45 {
		wv.wr.HoldRatioTotal = 5
	} else {
		wv.wr.HoldRatioTotal = 0
	}
}

func (wv *WeightData) CalFreeholdRatioTotal() {
	if len(wv.HoldRatioTotal) <= 0 {
		return
	}

	currentFreeHold := wv.FreeholdRatioTotal[0]
	if currentFreeHold < 10 {
		return
	}
	wv.wr.FreeholdRatioTotal = 0
}

func (wv *WeightData) CalPriceDiff() {

}

func (wv *WeightData) GetWeight() int32 {
	wr, weight := wv.wr, float64(0)
	log.Infof("==>>TODO 451: %+v|%+v|%+v|%+v|%+v|%+v", wr.Price, wr.Focus, wr.TotalNumRatio.Value, wr.AvgFreesharesRatio.Value, wr.HoldRatioTotal, wr.FreeholdRatioTotal)
	// 判断与当前的差价率
	if wv.RecentPrice/wv.Price[0] > priceDiff {
		log.Infof("invalid price: %s|%f|%f", wv.Secucode, wv.RecentPrice, wv.Price[0])
		return 0
	}
	// 判断时间差值
	// log.Infof("==>>TODO 452: %+v|%+v|%+v", wv.Secucode, time.Unix(wv.Date[0], 0).Format("2006-01-02"), wv.Date[0])
	if time.Now().Unix()-wv.Date[0] > dateDiff {
		log.Infof("invalid date: %s|%s", wv.Secucode, time.Unix(wv.Date[0], 0).Format("2006-01-02"))
		return 0
	}

	// log.Infof("==>>TODO 458: %+v|%+v|%+v|%+v|%+v|%+v", wr.Price, wr.Focus, wr.TotalNumRatio.Value, wr.AvgFreesharesRatio.Value, wr.HoldRatioTotal, wr.FreeholdRatioTotal)
	weight = wr.Price + wr.Focus + wr.TotalNumRatio.Value + wr.AvgFreesharesRatio.Value + wr.HoldRatioTotal + wr.FreeholdRatioTotal

	wv.Weight = int32(weight)
	// wv.Weight, _ = strconv.ParseFloat(fmt.Sprintf("%.1f", weight), 64)
	log.Infof("cal weight: %+v|%+v", wv.Secucode, wv.Weight)
	return wv.Weight
}

func intSlice2Str(values []float64, sep string) string {
	results := make([]string, 0, len(values))
	for _, val := range values {
		results = append(results, fmt.Sprintf("%v", val))
	}
	return strings.Join(results, sep)
}

func tmSlice2Str(dates []int64, sep string) string {
	var results []string
	for _, val := range dates {
		tm := time.Unix(val, 0)
		results = append(results, fmt.Sprintf("%d-%d", tm.Month(), tm.Day()))
	}
	return strings.Join(results, sep)
}

func Decimal(value float64) float64 {
	value, _ = strconv.ParseFloat(fmt.Sprintf("%.2f", value), 64)
	return value
}
