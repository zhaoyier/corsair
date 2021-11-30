package job

import (
	"time"

	orm "git.ezbuy.me/ezbuy/corsair/digger/service/internal/model"
	log "github.com/Sirupsen/logrus"
)

func UpdateJob(msg string) {
	createDate := time.Now().Format("2006-01-02")

	result, err := orm.JobMgr.FindOneByCreateDate(createDate)
	if err != nil {
		log.Errorf("query job failed: %s|%q", createDate, err)
		return
	}
	result.Msg = append(result.Msg, msg)
	result.UpdateDate = time.Now().Unix()
	if _, err := result.Save(); err != nil {
		log.Errorf("save job failed: %s|%q", createDate, err)
	}
}
