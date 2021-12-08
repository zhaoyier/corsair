package utils

import (
	"strings"
	"time"

	orm "git.ezbuy.me/ezbuy/corsair/digger/service/internal/model"
	trpc "git.ezbuy.me/ezbuy/corsair/digger/service/internal/rpc"
	log "github.com/Sirupsen/logrus"
	ezdb "github.com/ezbuy/ezorm/db"
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

	for key, val := range result.Msg {
		if key == typ.String() {
			return false
		}

		if !strings.Contains(val, "&") {
			return false
		}
	}

	updateJob(typ, result)

	return true
}

func updateJob(typ trpc.FunctionType, result *orm.Job) {
	sess, col := orm.JobMgr.GetCol()
	defer sess.Close()

	if len(result.Msg) == 0 {
		result.Msg = make(map[string]string)
	}

	result.Msg[typ.String()] = TS2Date(time.Now().Unix())
	query := ezdb.M{"CreateDate": result.CreateDate}
	update := ezdb.M{
		"$set": ezdb.M{
			"Msg":        result.Msg,
			"UpdateDate": time.Now().Unix(),
		},
	}

	if err := col.Update(query, update); err != nil {
		log.Errorf("update job failed: %q", err)
	}
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
