package job

import (
	"fmt"
	"time"

	orm "git.ezbuy.me/ezbuy/corsair/digger/service/internal/model"
	trpc "git.ezbuy.me/ezbuy/corsair/digger/service/internal/rpc"
	"git.ezbuy.me/ezbuy/corsair/digger/service/internal/utils"
	ezdb "github.com/ezbuy/ezorm/db"
	mgo "gopkg.in/mgo.v2"
)

func UpdateJob(typ trpc.FunctionType) {
	createDate := time.Now().Format("2006-01-02")
	sess, col := orm.JobMgr.GetCol()
	defer sess.Close()

	var result orm.Job
	query := ezdb.M{"CreateDate": createDate}
	if err := col.Find(query).One(&result); err != nil {
		result = *orm.JobMgr.NewJob()
		result.CreateDate = createDate
		result.Msg = make(map[string]string)
	}

	val, ok := result.Msg[typ.String()]
	if !ok {
		result.Msg[typ.String()] = utils.TS2Date(time.Now().Unix())
	} else {
		result.Msg[typ.String()] = fmt.Sprintf("%s&%s", val, utils.TS2Date(time.Now().Unix()))
	}

	change := mgo.Change{
		Update: ezdb.M{
			"$set": ezdb.M{
				"Msg":        result.Msg,
				"CreateDate": createDate,
				"UpdateDate": time.Now().Unix(),
			},
		},
		Upsert:    true,
		ReturnNew: true,
	}

	col.Find(query).Apply(change, nil)
}
