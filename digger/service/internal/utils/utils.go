package utils

import (
	"fmt"
	"math"
	"strconv"
	"strings"
	"time"

	log "github.com/Sirupsen/logrus"
	"github.com/shopspring/decimal"
)

func GetZeroTS() int64 {
	timeStr := time.Now().Format("2006-01-02")
	t, _ := time.ParseInLocation("2006-01-02", timeStr, time.Local)
	return t.Unix()
}

func GetDateTS(date string) int64 {
	t, err := time.ParseInLocation("2006-01-02", date, time.Local)
	if err != nil {
		log.Errorf("get data ts failed: %q", err)
	}
	return t.Unix()
}

func IsSubNew(ts int64) bool {
	diff := time.Now().Unix() - ts
	return diff <= 365*86400
}

func GetFocusStr(cells []string, sep string) string {
	var str string
	if len(cells) <= 0 {
		return str
	}
	pre, str := cells[0], cells[0]
	for idx, cell := range cells {
		if idx <= 0 || idx > 5 {
			continue
		}
		if cell == pre {
			str += "."
		} else {
			str += sep + cell
		}
		pre = cell
	}
	return str
}

func GetDateStr(cells []int64, sep string) string {
	var results []string
	for idx, cell := range cells {
		if idx > 5 {
			continue
		}

		tm := time.Unix(cell, 0)
		tmp := fmt.Sprintf("%d.%d", tm.Month(), tm.Day())
		results = append(results, fmt.Sprintf("%s", tmp))
	}

	return strings.Join(results, sep)
}

// 最近
func GetGDReduceRatio(cells []float64, sep string) string {
	list := make([]string, 0, len(cells))
	for _, cell := range cells {
		list = append(list, fmt.Sprintf("%.1f", cell))
	}

	return strings.Join(list, sep)
}

func Decimal(value float64) float64 {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("recover...:", r)
		}
	}()

	result, _ := decimal.NewFromFloat(value).Round(2).Float64()
	return result
}

func DecreasePercent(max, min float64) int32 {
	if max <= 0 || min <= 0 {
		return 0
	}

	rate := Decimal((max - min) / max)
	return int32(math.Ceil(rate * 100))
}

func GetPercentum(data int32) float64 {
	val := float64(data) / float64(100)
	return Decimal(val)
}

func GetRate(max, min float64) float64 {
	if max <= 0 || min <= 0 {
		return 0
	}

	return Decimal((max - min) / max)
}

func FloatSlice2Str(values []float64, sep string) string {
	results := make([]string, 0, len(values))
	for _, val := range values {
		results = append(results, fmt.Sprintf("%v", val))
	}
	return strings.Join(results, sep)
}

func TS2Date(ts int64) string {
	if ts == 0 {
		return "unknown"
	}

	return time.Unix(ts, 0).Format("2006-01-02 15:04:05")
}

func TS2Day(ts int64) string {
	if ts == 0 {
		return "unknown"
	}

	return time.Unix(ts, 0).Format("2006-01-02")
}

func GetSecucode(secucode string) string {
	codes := strings.Split(secucode, ".")
	for _, code := range codes {
		if _, err := strconv.ParseInt(code, 10, 64); err == nil {
			return code
		}
	}
	return ""
}

func GetSecucodeNum(secucode string) int64 {
	codes := strings.Split(secucode, ".")
	for _, code := range codes {
		if val, err := strconv.ParseInt(code, 10, 64); err == nil {
			return val
		}
	}
	return 0
}

func GetPercent(a, b float64) float64 {
	if b <= 0 {
		return 0
	}
	ratio := a / b
	if ratio > 1 {
		return 0
	}
	result := TruncateFloat(ratio) * 100
	return result
}

func String2Float64(data string) float64 {
	result, err := strconv.ParseFloat(data, 64)
	if err != nil {
		return 0
	}
	return result
}

func String2I32(data string) int32 {
	result, err := strconv.ParseFloat(data, 64)
	if err != nil {
		return 0
	}
	return int32(result)
}

// uint 代表小数位数，格式位 0.000001 如果是几位就指定为几位
func TruncateFloat(f float64) float64 {
	str := fmt.Sprintf("%.2f", f)
	result, _ := strconv.ParseFloat(str, 64)
	return result
}
