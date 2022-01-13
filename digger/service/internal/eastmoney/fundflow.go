package eastmoney

import (
	"time"

	"git.ezbuy.me/ezbuy/corsair/digger/service/internal/job"
	orm "git.ezbuy.me/ezbuy/corsair/digger/service/internal/model"
	"git.ezbuy.me/ezbuy/corsair/digger/service/internal/request"
	trpc "git.ezbuy.me/ezbuy/corsair/digger/service/internal/rpc"
	"git.ezbuy.me/ezbuy/corsair/digger/service/internal/utils"
	log "github.com/Sirupsen/logrus"
	ezdb "github.com/ezbuy/ezorm/db"
)

func GetFundFlowTicker() {
	GetFundFlow()
	job.UpdateJob(trpc.FunctionType_FunctionTypeFundFlow)
}

func GetFundFlowOnce() {
	getFundFlowOnce.Do(func() {
		GetFundFlow()
	})
}

func GetFundFlow() {
	sess, col := orm.GPRecommendMgr.GetCol()
	defer sess.Close()

	query := ezdb.M{}

	//查询最近6个月的数据
	var data orm.CNSecucode
	iter := col.Find(query).Batch(100).Prefetch(0.25).Iter()

	for iter.Next(&data) {
		secucode := utils.GetSecucodeNum(data.Secucode)
		if secucode <= 0 {
			continue
		}

		resp := new(EMFundFlow)
		if err := request.GetEastmoneyFundFlow(secucode, resp); err != nil {
			log.Errorf("eastmoney get five failed: %s|%+v\n", secucode, err)
			continue
		}

		if resp.Rc != 0 {
			// log.Infof("==>>TODO 812: %+v", secucode)
			continue
		}

		updateFundFlow(data.Secucode, resp)
	}
}

func updateFundFlow(secucode string, req *EMFundFlow) error {
	diff := req.Data.Diff[0]
	data, err := orm.GPFundFlowMgr.FindOneBySecucode(secucode)
	if err != nil {
		data = orm.GPFundFlowMgr.NewGPFundFlow()
		data.Secucode = secucode
	}

	data.Five = int32(diff.F164)
	data.Ten = int32(diff.F174)
	data.Twenty = int32(diff.F252)
	data.UpdateDate = time.Now().Unix()

	_, err = data.Save()
	return err
}
