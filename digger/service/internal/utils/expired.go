package utils

import (
	"time"

	orm "git.ezbuy.me/ezbuy/corsair/digger/service/internal/model"
	trpc "git.ezbuy.me/ezbuy/corsair/digger/service/internal/rpc"
)

func CheckFuncValid(typ trpc.FunctionType) bool {
	nowHour := time.Now().Local().Hour()
	createDate := time.Now().Format("2006-01-02")
	if nowHour < 17 {
		return false
	}

	result, err := orm.JobMgr.FindOneByCreateDate(createDate)
	if err != nil {
		return true
	}
	return int(typ)-len(result.Msg) == 1
}

// func UpdateFunction(typ trpc.FunctionType) {

// 	job.UpdateJob(typ.String())
// }
