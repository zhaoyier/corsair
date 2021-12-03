package utils

import (
	"fmt"
	"sync"
	"time"

	"git.ezbuy.me/ezbuy/corsair/digger/service/internal/job"
	trpc "git.ezbuy.me/ezbuy/corsair/digger/service/internal/rpc"
)

var (
	dataMap sync.Map
)

func CheckFuncValid(typ trpc.FunctionType) bool {
	now := time.Now()
	datets := GetZeroTS()
	weekday := now.Weekday()
	nowHour := now.Local().Hour()
	if weekday == time.Saturday || weekday == time.Sunday { //周
		return false
	}
	if nowHour < 18 {
		return false
	}

	max := int32(0)
	tnum := int32(typ)
	val, ok := dataMap.Load(datets)
	if !ok && tnum != 1 {
		return false
	}

	switch val.(type) {
	case nil:
		return true
	}

	for _, l := range val.([]int32) {
		if l > max {
			max = l
		}
		if l == tnum {
			return false
		}
	}

	return tnum-max == 1
}

func UpdateFunction(typ trpc.FunctionType) {
	datets := GetZeroTS()
	tnum := int32(typ)
	list := make([]int32, 0, 4)
	val, ok := dataMap.Load(datets)
	if ok {
		list = val.([]int32)
	}
	list = append(list, tnum)

	dataMap.Store(datets, list)

	job.UpdateJob(fmt.Sprintf("%s-%d", typ.String(), time.Now().Unix()))
}
