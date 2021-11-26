package dawdle

import (
	"fmt"
	"math"
	"strconv"
	"strings"
	"time"

	log "github.com/Sirupsen/logrus"
)

const (
	shake_upper   float64 = 1.4 //不变
	shake_lower   float64 = 0.6
	one_ten       float64 = 0.1
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
)

type WeightRule struct {
	Price              float64
	Focus              float64
	TotalNumRatio      float64
	AvgFreesharesRatio float64
	HoldRatioTotal     float64
	FreeholdRatioTotal float64
}

type WeightValue struct {
	Secucode           string
	Price              []float64
	Focus              []string
	TotalNumRatio      []float64
	AvgFreesharesRatio []float64
	HoldRatioTotal     []float64
	FreeholdRatioTotal []float64
	Date               []int64
	RecentPrice        float64
	Weight             float64
	wr                 *WeightRule
}

func NewWeightValue(secucode string) *WeightValue {
	return &WeightValue{
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
		TotalNumRatio:      15,
		AvgFreesharesRatio: 10,
		Focus:              30,
		Price:              25,
		HoldRatioTotal:     10,
		FreeholdRatioTotal: 10,
	}
}

func (wv *WeightValue) Cal() *WeightValue {
	wv.CalPrice()
	wv.CalFocus()
	wv.CalTotalNumRatio()
	wv.CalAvgFreesharesRatio()
	wv.CalHoldRatioTotal()
	wv.CalFreeholdRatioTotal()
	return wv
}

func (wv *WeightValue) CalPrice() {
	max, prev := wv.Price[0], wv.Price[0]
	if len(wv.Price) > cumulantPrice {
		wv.Price = wv.Price[0:cumulantPrice]
	}

	for idx, val := range wv.Price {
		// log.Infof("==>>TODO Price 401:%s|%+v|%f|%+v", wv.Secucode, idx, val, val <= 0)
		if idx > cumulantPrice || val <= 0 {
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
	}

	rate := Decimal((max - wv.Price[0]) / max)
	// log.Infof("==>>TODO Price 405:%+v|%+v|%+v", max, wv.Price[0], rate)
	if rate < two_fifth && rate > one_ten {
		wv.wr.Price = wv.wr.Price * 0.5
	} else if rate <= one_ten {
		wv.wr.Price = 0
	}
}

func (wv *WeightValue) CalFocus() {
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

func (wv *WeightValue) CalTotalNumRatio() { //+-,越小越好
	var min float64
	for idx, val := range wv.TotalNumRatio {
		value := val
		if idx <= cumulantRatio && value < 0 {
			if value < min {
				min = value
			}
		}
	}

	if rate := math.Abs(min / TotalNumRatio); rate < 1 {
		wv.wr.TotalNumRatio = Decimal(wv.wr.TotalNumRatio * rate)
	}
}

func (wv *WeightValue) CalAvgFreesharesRatio() { //+-,越大越好
	var max float64
	for idx, val := range wv.AvgFreesharesRatio {
		value := val
		if idx <= cumulantRatio && value > 0 {
			if value > max {
				max = value
			}
		}
	}

	if rate := math.Abs(max / TotalNumRatio); rate < 1 {
		wv.wr.AvgFreesharesRatio = Decimal(wv.wr.AvgFreesharesRatio * rate)
	}
}

func (wv *WeightValue) CalHoldRatioTotal() {
	if len(wv.HoldRatioTotal) <= 0 {
		return
	}

	currentHold := wv.HoldRatioTotal[0]
	if currentHold > 60 {
		return
	}
	wv.wr.HoldRatioTotal = 0
}

func (wv *WeightValue) CalFreeholdRatioTotal() {
	if len(wv.HoldRatioTotal) <= 0 {
		return
	}

	currentFreeHold := wv.FreeholdRatioTotal[0]
	if currentFreeHold < 10 {
		return
	}
	wv.wr.FreeholdRatioTotal = 0
}

func (wv *WeightValue) CalPriceDiff() {

}

func (wv *WeightValue) GetWeight() float64 {
	wr, weight := wv.wr, float64(0)
	// log.Infof("==>>TODO 451: %+v|%+v|%+v|%+v|%+v|%+v", wr.Price, wr.Focus, wr.TotalNumRatio, wr.AvgFreesharesRatio, wr.HoldRatioTotal, wr.FreeholdRatioTotal)
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

	// log.Infof("==>>TODO 458: %+v|%+v|%+v|%+v|%+v|%+v", wr.Price, wr.Focus, wr.TotalNumRatio, wr.AvgFreesharesRatio, wr.HoldRatioTotal, wr.FreeholdRatioTotal)
	weight = wr.Price + wr.Focus + wr.TotalNumRatio + wr.AvgFreesharesRatio + wr.HoldRatioTotal + wr.FreeholdRatioTotal

	wv.Weight, _ = strconv.ParseFloat(fmt.Sprintf("%.1f", weight), 64)
	// log.Infof("==>>TODO 459: %+v|%+v", wv.Secucode, wv.Weight)
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
