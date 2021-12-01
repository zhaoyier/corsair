package job

import (
	"fmt"
	"time"

	orm "git.ezbuy.me/ezbuy/corsair/digger/service/internal/model"
	log "github.com/Sirupsen/logrus"
)

func UpdateJob(msg string) {
	createDate := time.Now().Format("2006-01-02")

	result := orm.JobMgr.MustFindOneByCreateDate(createDate)
	result.Msg = append(result.Msg, fmt.Sprintf("%s-%d", msg, time.Now().Unix()))
	result.UpdateDate = time.Now().Unix()
	if _, err := result.Save(); err != nil {
		log.Errorf("save job failed: %s|%q", createDate, err)
	}
}
