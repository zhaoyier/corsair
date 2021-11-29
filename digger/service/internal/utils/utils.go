package utils

import (
	"fmt"
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
