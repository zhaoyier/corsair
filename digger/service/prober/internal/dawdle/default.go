package dawdle

import (
	"math"
	"sync"
	"time"

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
	GPShortDecrease int32   = 35
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
		TotalNumRatio:      &trpc.WeightUnit{Value: 30}, //人均持股变化
		Focus:              &trpc.WeightUnit{Value: 20}, //集中度
		Price:              &trpc.WeightUnit{Value: 20}, //价格
		HoldRatioTotal:     15,
		FreeholdRatioTotal: 15,
	}
}

func (wv *WeightData) Cal() *WeightData {
	wv.calOnce.Do(func() {
		wv.CalPrice()
		wv.CalFocus()
		wv.CalTotalNumRatio()
		wv.CalHoldRatioTotal()
		wv.CalFreeholdRatioTotal()
	})

	return wv
}

func (wv *WeightData) CalPrice() {
	unit := wv.wr.Price
	// log.Infof("==>>TODO Price 3211:%+v", len(wv.Price))
	max, min, prev := wv.Price[0], wv.Price[0], wv.Price[0]
	for _, val := range wv.Price {
		// log.Infof("==>>TODO Price 321:%+v|%+v", val, val < 0)
		// 是否为新股, 新股跌幅要大
		unit.SubNew = val <= 0
		wv.wr.IsNew = val <= 0
		if wv.wr.IsNew {
			break
		}
	}

	for idx, val := range wv.Price {
		unit.Counter++
		if val <= 0 {
			continue
		}

		// 价格跳动太大就舍弃了
		// log.Infof("==>>TODO Price 322:%+v|%+v|%+v|%+v", prev, val, math.Abs(prev-val), val)
		// log.Infof("==>>TODO Price 323:%s|%+v|%f|%f", wv.Secucode, idx, (math.Abs(prev-val) / val), priceDiff*math.Pow(1.1, float64(idx)))
		if (math.Abs(prev-val) / val) > priceDiff*math.Pow(1.1, float64(idx)) {
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
	rate := utils.GetRate(max, wv.Price[0])
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
	// rate2 := utils.GetRate(wv.GPDaily.Closing, min)
	// if rate2 >= 1 {
	// 	unit.Value = 0
	// } else if rate2 > 0.5 {
	// 	unit.Value = unit.Value * 0.5
	// }

	// log.Infof("==>>TODO Price 329:%+v|%+v|%+v", unit.Value, 0, 0)
}

func (wv *WeightData) CalFocus() {
	var weight float64
	for _, val := range wv.Focus {
		if val == "非常集中" || val == "较集中" {
			weight += 0.25
		} else {
			break
		}
	}
	wv.wr.Focus.Value = wv.wr.Focus.Value * weight
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
		// 股东人数减少35就退出
		if unit.Accum <= -35 {
			// log.Infof("==>>TODO 303:%+v", unit)
			break
		}

		pre = val
	}
	// log.Infof("==>>TODO 305:%+v", unit)
	// 最近是减少的
	// if unit.Consecutive >= 2 && unit.Counter <= 4 {
	// 	rate += 0.25
	// }
	if unit.Accum <= -35 && unit.Consecutive <= 3 {
		rate += 1
	} else if unit.Accum <= -25 && unit.Consecutive <= 3 {
		rate += 0.75
	} else if unit.Accum <= -15 && unit.Consecutive <= 3 {
		rate += 0.5
	} else if unit.Accum <= -5 && unit.Consecutive <= 3 {
		rate += 0.25
	} else if unit.Accum < -35 {
		rate += 1
	} else if unit.Accum < -25 {
		rate += 0.5
	} else {
		rate = 0
	}

	if rate > 1 {
		rate = 1
	}
	// 股东人数变化率
	unit.Value = utils.Decimal(unit.Value * rate)
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

	if holdRatio > 65 {
		rate = 1
	} else if holdRatio > 45 {
		rate = 0.75
	} else {
		rate = 0.25
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

func (wv *WeightData) GetWeight() int32 {
	wv.weightOnce.Do(func() {
		wr, weight := wv.wr, float64(0)
		// log.Infof("==>>TODO 450: %+v", wv)
		// log.Infof("==>>TODO 451: %+v|%+v|%+v|%+v|%+v|%+v|%+v", wv.Secucode, wr.Price.Value, wr.Focus, wr.TotalNumRatio.Value, wr.AvgFreesharesRatio.Value, wr.HoldRatioTotal, wr.FreeholdRatioTotal)
		// 判断与当前的差价率
		if wv.Price[0] <= 0 || utils.GetRate(wv.GPDaily.Closing, wv.Price[0]) > 0.25 {
			log.Infof("invalid price: %s|%f|%f", wv.Secucode, wv.GPDaily.Closing, wv.Price[0])
			wv.Weight = 0
			return
		}
		// 判断时间差值
		// log.Infof("==>>TODO 452: %+v|%+v|%+v", wv.Secucode, time.Unix(wv.Date[0], 0).Format("2006-01-02"), wv.Date[0])
		if time.Now().Unix()-wv.Date[0] > dateDiff {
			log.Infof("invalid date: %s|%s", wv.Secucode, time.Unix(wv.Date[0], 0).Format("2006-01-02"))
			wv.Weight = 0
			return
		}

		log.Infof("==>>TODO 458: %+v|%+v|%+v|%+v|%+v|%+v", wr.Price, wr.Focus, wr.TotalNumRatio.Value, wr.AvgFreesharesRatio.Value, wr.HoldRatioTotal, wr.FreeholdRatioTotal)
		weight = wr.Price.Value + wr.Focus.Value + wr.TotalNumRatio.Value + wr.HoldRatioTotal + wr.FreeholdRatioTotal
		wv.Weight = int32(weight)
		// log.Infof("==>>TODO 459: %+v", args ...interface{})
		// wv.Weight, _ = strconv.ParseFloat(fmt.Sprintf("%.1f", weight), 64)
		log.Infof("cal weight: %+v|%+v", wv.Secucode, wv.Weight)
	})

	return wv.Weight
}
