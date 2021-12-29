package recommend

import (
	orm "git.ezbuy.me/ezbuy/corsair/digger/service/internal/model"

	"git.ezbuy.me/ezbuy/corsair/digger/service/internal/dawdle"
	log "github.com/Sirupsen/logrus"
)

func getReferDecrease(secucode string) int32 {
	result, err := orm.GPDelayMgr.FindOneBySecucodeDisabled(secucode, false)
	if err != nil {
		log.Errorf("get delay failed: %s|%q", secucode, err)
		return dawdle.GPShortDecrease
	}
	if result.DecreaseTag <= 0 {
		return dawdle.GPShortDecrease
	}
	return result.DecreaseTag
}
