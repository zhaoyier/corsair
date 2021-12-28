package dawdle

import (
	"fmt"
	"time"

	orm "git.ezbuy.me/ezbuy/corsair/digger/service/internal/model"
	trpc "git.ezbuy.me/ezbuy/corsair/digger/service/internal/rpc"
	log "github.com/Sirupsen/logrus"
	mgo "gopkg.in/mgo.v2"
)

func SavePromptBuy(data *orm.GPRecommend, minPrice float64) error {
	switch trpc.RMState(data.State) {
	case trpc.RMState_RMStateInProgress:
	case trpc.RMState_RMStateStarted:
		break
	default:
		return fmt.Errorf("invalid state: %d", data.State)
	}

	result, err := orm.GPPromptBuyMgr.FindOneBySecucodeDisabled(data.Secucode, false)
	if err != nil && err != mgo.ErrNotFound {
		log.Errorf("get prompt failed: %s|%q", data.Secucode, err)
		return err
	}
	if result == nil {
		result = orm.GPPromptBuyMgr.NewGPPromptBuy()
		result.Secucode = data.Secucode
		result.Name = data.Name
	}

	if result.Start <= 0 {
		result.Start = time.Now().Unix()
	}

	if result.MinPrice <= 0 || result.MinPrice > minPrice {
		result.MinPrice = minPrice
	}

	result.UpdateDate = time.Now().Unix()
	if _, err := result.Save(); err != nil {
		log.Errorf("update prompt failed: %s|%q", data.Secucode, err)
		return err
	}
	return nil
}
