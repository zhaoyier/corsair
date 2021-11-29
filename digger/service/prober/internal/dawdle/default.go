package dawdle

import (
	"fmt"
	"math"
	"strconv"
	"strings"
	"sync"
	"time"

	trpc "git.ezbuy.me/ezbuy/corsair/digger/rpc/digger"
	orm "git.ezbuy.me/ezbuy/corsair/digger/service/internal/model"
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
	Price              *trpc.WeightUnit //float64
	Focus              float64
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
		Price:              &trpc.WeightUnit{Value: 20}, //20,
		HoldRatioTotal:     10,
		FreeholdRatioTotal: 10,
	}
}

func (wv *WeightData) Cal() *WeightData {
	wv.calOnce.Do(func() {
		wv.CalPrice()
		wv.CalFocus()
		wv.CalTotalNumRatio()
		wv.CalAvgFreesharesRatio()
		wv.CalHoldRatioTotal()
		wv.CalFreeholdRatioTotal()
	})

	return wv
}

func (wv *WeightData) CalPrice() {
	unit := wv.wr.Price
	max, min, prev := wv.Price[0], wv.Price[0], wv.Price[0]
	for _, val := range wv.Price {
		// log.Infof("==>>TODO Price 321:%+v|%+v", val, val < 0)
		unit.SubNew = val <= 0
		if unit.SubNew {
			break
		}
	}

	if len(wv.Price) > cumulantPrice {
		wv.Price = wv.Price[0:cumulantPrice]
	}
	// log.Infof("==>>TODO Price 320:%+v", len(wv.Price))

	for idx, val := range wv.Price {
		unit.Counter++
		// log.Infof("==>>TODO Price 401:%s|%+v|%f|%+v", wv.Secucode, idx, val, val <= 0)
		if val <= 0 {
			continue
		}

		// 价格跳动太大就舍弃了
		// log.Infof("==>>TODO Price 322:%+v|%+v|%+v|%+v", prev, val, math.Abs(prev-val), val)
		// log.Infof("==>>TODO Price 323:%s|%+v|%f|%f", wv.Secucode, idx, (math.Abs(prev-val) / val), priceDiff*math.Pow(1.1, float64(idx)))
		if (math.Abs(prev-val) / val) > priceDiff*math.Pow(1.1, float64(idx)) {
			// wv.wr.Price = 0
			// log.Infof("==>>TODO Price 326:%+v|%+v", prev, val)
			unit.Value = 0
			return
		}

		if max < val {
			max = val
		}

		if min > val && val != 0 {
			min = val
		}
	}

	if unit.Counter <= 0 {
		unit.Value = 0
		return
	}

	// 与最近一次价格比较--越跌越好
	rate := Decimal((max - wv.Price[0]) / max)
	// log.Infof("==>>TODO Price 324:%+v|%+v|%+v", max, wv.Price[0], rate)
	if rate <= one_ten {
		unit.Value = 0
	} else if rate < 0.3 {
		unit.Value = unit.Value * 0.5
	} else if rate < 0.4 {
		unit.Value = unit.Value * 0.8
	} else if rate < 0.5 {
		unit.Value = unit.Value * 0.9
	}
	// log.Infof("==>>TODO Price 325:%+v|%+v|%+v", unit.Value, wv.GPDaily.Closing, min)
	//100%直接放弃--越涨越差
	rate2 := Decimal((wv.GPDaily.Closing - min) / min)
	if rate2 >= 1 {
		unit.Value = 0
	} else if rate2 > 0.5 {
		unit.Value = unit.Value * 0.5
	}

	// log.Infof("==>>TODO Price 329:%+v|%+v|%+v", unit.Value, 0, 0)
}

func (wv *WeightData) CalFocus() {
	var weight float64
	for idx, val := range wv.Focus {
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
	for idx, val := range wv.TotalNumRatio {
		// 计数器
		unit.Counter++
		// 累加值
		unit.Accum += val
		if idx == 0 && val <= -35 {
			rate = 1
			break
		}
		// 计算连续值类型
		if val < 0 && (pre == 0 || pre < 0) {
			unit.Consecutive++
		} else {
			unit.Consecutive--
		}
		// log.Infof("==>>TODO 302:%+v|%+v", unit, unit.Accum <= -15)
		if unit.Accum <= -15 {
			// log.Infof("==>>TODO 303:%+v", unit)
			break
		}

		pre = val
	}
	// log.Infof("==>>TODO 305:%+v", unit)
	// 最近是减少的
	if unit.Consecutive >= 2 && unit.Counter <= 4 {
		rate += 0.25
	}
	if unit.Accum <= -1*25 && unit.Consecutive <= 3 {
		rate += 0.75
	} else if unit.Accum <= -1*15 && unit.Consecutive <= 3 {
		rate += 0.6
	} else if unit.Accum <= -1*10 && unit.Consecutive <= 3 {
		rate += 0.45
	} else if unit.Accum <= -1*5 && unit.Consecutive <= 3 {
		rate += 0.3
	}
	// log.Infof("==>>TODO 309:%+v", rate)
	unit.Value = Decimal(unit.Value * math.Min(rate, 1))
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
	unit.Value = Decimal(unit.Value * math.Min(rate, 1))
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
	wv.weightOnce.Do(func() {
		wr, weight := wv.wr, float64(0)
		// log.Infof("==>>TODO 450: %+v", wv.Focus)
		// log.Infof("==>>TODO 451: %+v|%+v|%+v|%+v|%+v|%+v|%+v", wv.Secucode, wr.Price.Value, wr.Focus, wr.TotalNumRatio.Value, wr.AvgFreesharesRatio.Value, wr.HoldRatioTotal, wr.FreeholdRatioTotal)
		// 判断与当前的差价率
		if wv.GPDaily.Closing/wv.Price[0] > priceDiff {
			log.Infof("invalid price: %s|%f|%f", wv.Secucode, wv.GPDaily.Closing, wv.Price[0])
			wv.Weight = 0
			return
		}
		// 判断时间差值
		// log.Infof("==>>TODO 452: %+v|%+v|%+v", wv.Secucode, time.Unix(wv.Date[0], 0).Format("2006-01-02"), wv.Date[0])
		if time.Now().Unix()-wv.Date[0] > dateDiff && !wr.Price.SubNew {
			log.Infof("invalid date: %s|%s", wv.Secucode, time.Unix(wv.Date[0], 0).Format("2006-01-02"))
			wv.Weight = 0
			return
		} else if time.Now().Unix()-wv.Date[0] > dateDiff && !wr.Price.SubNew {
			wv.Weight -= 10
		}

		// log.Infof("==>>TODO 458: %+v|%+v|%+v|%+v|%+v|%+v", wr.Price, wr.Focus, wr.TotalNumRatio.Value, wr.AvgFreesharesRatio.Value, wr.HoldRatioTotal, wr.FreeholdRatioTotal)
		weight = wr.Price.Value + wr.Focus + wr.TotalNumRatio.Value + wr.AvgFreesharesRatio.Value + wr.HoldRatioTotal + wr.FreeholdRatioTotal
		wv.Weight = int32(weight)
		// wv.Weight, _ = strconv.ParseFloat(fmt.Sprintf("%.1f", weight), 64)
		log.Infof("cal weight: %+v|%+v", wv.Secucode, wv.Weight)
	})

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
