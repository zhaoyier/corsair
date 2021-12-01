package utils

import (
	"fmt"
	"math"
	"strings"
	"time"

	log "github.com/Sirupsen/logrus"
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
			str += sep + ".."
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
	var counter int32
	var sum float64
	for _, cell := range cells {
		if cell > 0 {
			break
		}
		counter++
		sum += cell
	}
	return fmt.Sprintf("%d%s%.1f", counter, sep, sum)
}

func Decimal(value float64) float64 {
	return math.Ceil(value)
}

func DecreasePercent(max, min float64) int32 {
	if max <= 0 || min <= 0 {
		return 0
	}

	rate := (max - min) / max
	return int32(rate * 100)
}

func GetRate(max, min float64) float64 {
	if max <= 0 || min <= 0 {
		return 0
	}
	return math.Ceil((max - min) / max)
}

func FloatSlice2Str(values []float64, sep string) string {
	results := make([]string, 0, len(values))
	for _, val := range values {
		results = append(results, fmt.Sprintf("%v", val))
	}
	return strings.Join(results, sep)
}

func TS2Date(ts int64) string {
	return time.Unix(ts, 0).Format("2006-01-02 15:04:05")
}
