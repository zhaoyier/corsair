package dawdle

import (
	"math"
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
	zhouqiOnce sync.Once
)

func GenZhouQiTicker() {
	genZhouQiData()
	// 更新任务
	job.UpdateJob(trpc.FunctionType_FunctionTypeZhouQi)
}

func GenZhouQiOnce() {
	zhouqiOnce.Do(func() {
		genZhouQiData()
	})
}

func GenZhouQiTemp(secucode string) {
	result, err := orm.GPZhouQiMgr.FindOneBySecucode(secucode)
	if err != nil {
		log.Errorf("get zhouqi failed: %s|%q", secucode, err)
		return
	}

	updateZhouQiState(result)
}

func genZhouQiData() {
	s, c := orm.GPZhouQiMgr.GetCol()
	defer s.Close()

	var data orm.GPZhouQi
	iter := c.Find(ezdb.M{}).Batch(1000).Prefetch(0.25).Iter()
	for iter.Next(&data) {
		updateZhouQiState(&data)
	}
}

func updateZhouQiState(data *orm.GPZhouQi) error {
	data.UpdateDate = time.Now().Unix()
	data.PresentPrice = getPresentPrice(data.Secucode)
	data.State = genZhouQiState(data)

	if _, err := data.Save(); err != nil {
		log.Errorf("update zhouqi price failed: %s|%q", data.Secucode, err)
		return err
	}
	return nil
}

func getPresentPrice(secucode string) float64 {
	secucode = utils.GetSecucode(secucode)
	result, err := orm.GPDailyMgr.FindOne(ezdb.M{
		"Secucode": secucode,
	}, "-CreateDate")
	if err != nil {
		log.Errorf("get daily price failed: %s|%q", secucode, err)
		return 0
	}
	return math.Min(result.Closing, result.MinPrice)
}

func genZhouQiState(data *orm.GPZhouQi) int32 {
	state := trpc.GPZhouQiState_GPZhouQiStateUnknown
	now := time.Now().Unix()
	if data.ExpectStart > 0 && now > data.ExpectStart {
		state = trpc.GPZhouQiState_GPZhouQiStateDate
	}

	if data.PresentPrice < data.ExpectMax {
		state = trpc.GPZhouQiState_GPZhouQiStatePrice
	}

	return int32(state)
}
