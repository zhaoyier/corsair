package eastmoney

import (
	"strings"
	"time"

	"git.ezbuy.me/ezbuy/corsair/digger/service/internal/job"
	orm "git.ezbuy.me/ezbuy/corsair/digger/service/internal/model"
	"git.ezbuy.me/ezbuy/corsair/digger/service/internal/request"
	trpc "git.ezbuy.me/ezbuy/corsair/digger/service/internal/rpc"
	"git.ezbuy.me/ezbuy/corsair/digger/service/internal/utils"
	log "github.com/Sirupsen/logrus"
	ezdb "github.com/ezbuy/ezorm/db"
	// "/Users/zhaojianwei/Projects/ezbuy/goflow/src/git.ezbuy.me/ezbuy/corsair/digger/service/internal/utils"
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
	sess, col := orm.CNSecucodeMgr.GetCol()
	defer sess.Close()

	query := ezdb.M{}
	// query := ezdb.M{"Secucode": "SZ.300204"}

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

		// log.Infof("==>>TODO 812: %+v", resp)

		if resp.Rc != 0 {
			// log.Infof("==>>TODO 812: %+v", secucode)
			continue
		}
		parsingFFKlines(data.Secucode, resp)
		// updateFundFlow(data.Secucode, resp)
	}
}

func parsingFFKlines(secucode string, req *EMFundFlow) error {
	weekday := time.Now().Local().Weekday()
	if weekday == time.Saturday || weekday > time.Sunday {
		return nil
	}

	dateStr := time.Now().Local().Format("2006-01-02")
	if req.Data == nil || req.Data.Klines == nil || len(req.Data.Klines) <= 0 {
		return nil
	}

	// dateStr := "2022-01-14"
	log.Infof("==>>TODO 812: %+v", dateStr)
	rows := req.Data.Klines
	for _, row := range rows {
		if strings.Contains(row, dateStr) {
			log.Infof("==>>TODO 813: %+v", row)
			cells := strings.Split(row, ",")
			if len(cells) < 15 {
				continue
			}
			log.Infof("==>>TODO 814: %+v", cells)
			data := orm.GPFundFlowMgr.NewGPFundFlow()
			data.Secucode = secucode
			data.CreateDate = time.Now().Unix()

			for idx, cell := range cells {
				if idx == 0 {
					// date = cell
					data.FundDate = utils.GetDateTS(cell)
				}
				if idx > 0 && idx < 6 {
					data.Inflow += utils.String2I32(cell)
				}
				if idx >= 6 && idx <= 10 {
					data.InflowRatio += utils.String2I32(cell)
				}
				if idx == 11 {
					data.PresentPrice = utils.String2I32(cell)
				}
				if idx == 12 {
					data.IncreaseRatio = utils.String2I32(cell)
				}
			}
			log.Infof("==>>TODO 815: %+v", data)
			if _, err := data.Save(); err != nil {
				log.Errorf("save fund flow failed: %s|%q", secucode, err)
			}
		}
	}

	return nil
}

// func updateFundFlow(secucode string, req *EMFundFlow) error {
// 	diff := req.Data.Diff[0]
// 	data, err := orm.GPFundFlowMgr.FindOneBySecucode(secucode)
// 	if err != nil {
// 		data = orm.GPFundFlowMgr.NewGPFundFlow()
// 		data.Secucode = secucode
// 	}

// 	data.Five = int32(diff.F164)
// 	data.Ten = int32(diff.F174)
// 	data.Twenty = int32(diff.F252)
// 	data.UpdateDate = time.Now().Unix()

// 	_, err = data.Save()
// 	return err
// }
