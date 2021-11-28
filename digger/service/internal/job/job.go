package job

import (
	"time"

	orm "git.ezbuy.me/ezbuy/corsair/digger/service/internal/model"
	log "github.com/Sirupsen/logrus"
	ezdb "github.com/ezbuy/ezorm/db"
	mgo "gopkg.in/mgo.v2"
)

func UpdateJob(name, msg string) {
	createDate := time.Now().Format("2006-01-02")

	sess, col := orm.JobMgr.GetCol()

	defer sess.Close()

	query := ezdb.M{"Name": name, "CreateDate": createDate}

	change := mgo.Change{
		Update: ezdb.M{
			"$set": ezdb.M{
				"Msg":        msg,
				"UpdateDate": time.Now().Unix(),
			},
		},
		Upsert:    true,
		ReturnNew: true,
	}

	if _, err := col.Find(query).Apply(change, nil); err != nil {
		log.Errorf("%s|%q", name, err)
	}
}
