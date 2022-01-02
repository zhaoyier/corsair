package recommend

import (
	"fmt"
	"strings"

	orm "git.ezbuy.me/ezbuy/corsair/digger/service/internal/model"

	"git.ezbuy.me/ezbuy/corsair/digger/service/internal/dawdle"
	log "github.com/Sirupsen/logrus"
)

func getReferDecrease(secucode string) int32 {
	result, err := orm.GPManualDecreaseMgr.FindOneBySecucodeDisabled(secucode, false)
	if err != nil {
		log.Errorf("get delay failed: %s|%q", secucode, err)
		return dawdle.GPShortDecrease
	}
	if result.DecreaseTag <= 0 {
		return dawdle.GPShortDecrease
	}
	return result.DecreaseTag
}

func getRMPrice(data []float64) string {
	list := make([]string, 0, len(data))
	for idx, val := range data {
		list = append(list, fmt.Sprintf("%.1f(%d)", val, idx+1))
	}
	return strings.Join(list, "->")
}
