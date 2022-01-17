package zwadmin

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
		return dawdle.GetConf().DecreaseTag
	}
	if result.DecreaseTag <= 0 {
		return dawdle.GetConf().DecreaseTag
	}
	return result.DecreaseTag
}

func getRMPrice(data []float64) string {
	list := make([]string, 0, len(data))
	for idx, val := range data {
		list = append(list, fmt.Sprintf("%.1f(%d)", val, idx+2))
	}
	return strings.Join(list, "->")
}

func getFocusBySecucode(secucode string) string {
	result, err := orm.GPFocusMgr.FindOneBySecucode(secucode)
	if err != nil {
		return "关注"
	}
	if result == nil {
		return "关注"
	}
	return "取消关注"
}

func getFocusByName(name string) string {
	result, err := orm.GPFocusMgr.FindOneByName(name)
	if err != nil {
		return "关注"
	}
	if result == nil {
		return "关注"
	}
	return "取消关注"
}

func getSecucodeWithExchange(secucode string) string {
	var exchange string
	prefix := secucode[0:3]
	switch prefix {
	case "600", "601", "603", "605", "688":
		exchange = "SH"
	case "835", "836":
		exchange = "BJ"
	default:
		exchange = "SZ"
	}
	return fmt.Sprintf("%s.%s", exchange, secucode)
}
