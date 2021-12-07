package utils

import (
	"strings"
	"time"

	"git.ezbuy.me/ezbuy/corsair/digger/service/internal/job"
	orm "git.ezbuy.me/ezbuy/corsair/digger/service/internal/model"
	trpc "git.ezbuy.me/ezbuy/corsair/digger/service/internal/rpc"
	log "github.com/Sirupsen/logrus"
	"gopkg.in/mgo.v2"
)

func CheckFuncValid(typ trpc.FunctionType) bool {
	nowHour := time.Now().Local().Hour()
	weekday := time.Now().Local().Weekday()
	createDate := time.Now().Format("2006-01-02")
	if weekday == time.Saturday || weekday == time.Sunday {
		return false
	}

	if nowHour < 17 { //TODO 17
		return false
	}

	result, err := orm.JobMgr.FindOneByCreateDate(createDate)
	if err != nil && err != mgo.ErrNotFound {
		return false
	}

	if result == nil {
		result = orm.JobMgr.MustFindOneByCreateDate(createDate)
	}

	if len(result.Msg) >= int(trpc.FunctionType_FunctionTypeRecommend) {
		return false
	}

	if _, ok := result.Msg[typ.String()]; ok {
		return false
	} else {
		job.UpdateJob(trpc.FunctionType_FunctionTypeCodeList)
	}

	return true
}

func CheckFuncValid2(typ trpc.FunctionType) bool {

	nowHour := time.Now().Local().Hour()
	weekday := time.Now().Local().Weekday()
	createDate := time.Now().Format("2006-01-02")
	if weekday == time.Saturday || weekday == time.Sunday {
		return false
	}

	if nowHour > 17 { //TODO
		return false
	}

	result, err := orm.JobMgr.FindOne(createDate)
	log.Infof("==>>TODO 201: %+v", typ.String())
	if err != nil {
		return typ == trpc.FunctionType_FunctionTypeCodeList
	}

	log.Infof("==>>TODO 203: %+v", typ.String())
	if _, ok := result.Msg[typ.String()]; ok { //重复的
		return false
	}

	var max int32
	for key := range result.Msg {
		val := getFunctionNum(key)
		if max < val {
			max = val
		}
	}

	if max == 0 && typ != trpc.FunctionType_FunctionTypeCodeList { //不是第一个
		return false
	}

	diff := int32(typ) - max
	if diff != 1 {
		return false
	}

	mtp := trpc.FunctionType(max)
	if diff == 1 && !strings.Contains(result.Msg[mtp.String()], "-") {
		return false
	}

	return true
}

func getFunctionNum(typ string) int32 {
	switch typ {
	case "FunctionTypeCodeList":
		return 1
	case "FunctionTypeShareholder":
		return 2
	case "FunctionTypeLongLine":
		return 3
	case "FunctionTypeShortLine":
		return 4
	case "FunctionTypeRecommend":
		return 5
	default:
		return 1
	}
}
