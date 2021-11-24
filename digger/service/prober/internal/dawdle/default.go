package dawdle

import (
	"fmt"
	"math"
	"strconv"
	"strings"

	"qiniupkg.com/x/log.v7"
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
	Date               []string
	RecentPrice        float64
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
		Date:               make([]string, 0, 8),
		wr:                 defaultWeightRule(),
	}
}

func defaultWeightRule() *WeightRule {
	return &WeightRule{
		TotalNumRatio:      20,
		AvgFreesharesRatio: 10,
		Focus:              30,
		Price:              20,
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
	var max float64

	if len(wv.Price) <= 0 {
		return
	}

	for idx, val := range wv.Price {
		if idx > cumulantPrice {
			continue
		}

		if max < float64(val) {
			max = float64(val)
		}
	}

	rate := (max - wv.Price[0]) / max
	if rate < one_fifth && rate > one_ten {
		wv.wr.Price = wv.wr.Price * 0.5
	} else if rate < one_ten {
		wv.wr.Price = 0
	}
}

func (wv *WeightValue) CalFocus() float64 {
	var weight float64
	for idx, val := range wv.Focus {
		if idx <= cumulantRatio && (val == "非常集中" || val == "较集中") {
			weight += 0.25
		}
	}

	return float64(wv.wr.Focus) * weight
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
		wv.wr.TotalNumRatio = wv.wr.TotalNumRatio * rate
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
		wv.wr.AvgFreesharesRatio = wv.wr.AvgFreesharesRatio * rate
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
	if wv.RecentPrice/wv.Price[0] > priceDiff {
		return 0
	}

	log.Infof("==>>TODO 1002: %+v|%+v|%+v|%+v|%+v|%+v", wr.Price, wr.Focus, wr.TotalNumRatio, wr.AvgFreesharesRatio, wr.HoldRatioTotal, wr.FreeholdRatioTotal)
	weight = wr.Price + wr.Focus + wr.TotalNumRatio + wr.AvgFreesharesRatio + wr.HoldRatioTotal + wr.FreeholdRatioTotal

	value, _ := strconv.ParseFloat(fmt.Sprintf("%.1f", weight), 64)

	return value
}

func intSlice2Str(values []float64, sep string) string {
	results := make([]string, 0, len(values))
	for _, val := range values {
		results = append(results, fmt.Sprintf("%v", val))
	}
	return strings.Join(results, sep)
}
