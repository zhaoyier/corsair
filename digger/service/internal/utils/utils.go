package utils

import (
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
		if idx <= 0 {
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
