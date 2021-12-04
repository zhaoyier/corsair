package utils

import (
	"time"

	orm "git.ezbuy.me/ezbuy/corsair/digger/service/internal/model"
	trpc "git.ezbuy.me/ezbuy/corsair/digger/service/internal/rpc"
)

func CheckFuncValid(typ trpc.FunctionType) bool {
	nowHour := time.Now().Local().Hour()
	weekday := time.Now().Local().Weekday()
	createDate := time.Now().Format("2006-01-02")
	if weekday == time.Saturday || weekday == time.Sunday {
		return false
	}

	if nowHour < 17 {
		return false
	}

	result, err := orm.JobMgr.FindOneByCreateDate(createDate)
	if err != nil {
		return true
	}
	return int(typ)-len(result.Msg) == 1
}
