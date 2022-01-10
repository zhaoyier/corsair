package dawdle

import (
	"sync"
	"time"

	"git.ezbuy.me/ezbuy/corsair/digger/service/internal/job"
	orm "git.ezbuy.me/ezbuy/corsair/digger/service/internal/model"
	trpc "git.ezbuy.me/ezbuy/corsair/digger/service/internal/rpc"
	"git.ezbuy.me/ezbuy/corsair/digger/service/internal/utils"
	log "github.com/Sirupsen/logrus"
	ezdb "github.com/ezbuy/ezorm/db"
)

var (
	focusStateOnce sync.Once
)

func GenFocusStateTicker() {
	genZhouQiData()
	// 更新任务
	job.UpdateJob(trpc.FunctionType_FunctionTypeFocus)
}

func GenFocusStateOnce() {
	focusStateOnce.Do(func() {
		genFocusStateData()
	})
}

func GenFocusStateTemp(secucode string) {
	result, err := orm.GPFocusMgr.FindOneBySecucode(secucode)
	if err != nil {
		log.Errorf("get zhouqi failed: %s|%q", secucode, err)
		return
	}

	updateFocusState(result)
}

func genFocusStateData() {
	s, c := orm.GPFocusMgr.GetCol()
	defer s.Close()

	var data orm.GPFocus
	iter := c.Find(ezdb.M{}).Batch(1000).Prefetch(0.25).Iter()
	for iter.Next(&data) {
		updateFocusState(&data)
	}
}

func updateFocusState(data *orm.GPFocus) error {
	data.PresentPrice = getPresentPrice(data.Secucode)
	data.State = getFocusState(data)
	data.UpdateDate = time.Now().Unix()
	if _, err := data.Save(); err != nil {
		log.Errorf("update zhouqi price failed: %s|%q", data.Secucode, err)
		return err
	}
	return nil
}

func getFocusState(data *orm.GPFocus) int32 {
	state := trpc.GPFocusState_FocusStateUnknown
	if data.ExpectPrice <= 0 || data.PresentPrice <= 0 {
		return int32(state)
	}

	diff := data.PresentPrice - data.ExpectPrice
	percent := utils.GetPercent(diff, data.PresentPrice)
	if percent >= 0 && percent < 5 {
		state = trpc.GPFocusState_FocusStatePrepare
	} else if percent < 0 && percent >= -5 {
		state = trpc.GPFocusState_FocusStateStart
	} else {
		state = trpc.GPFocusState_FocusStateProgress
	}

	return int32(state)
}
