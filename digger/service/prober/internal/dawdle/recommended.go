package dawdle

import (
	"math"
	"sync"
	"time"

	"git.ezbuy.me/ezbuy/corsair/digger/service/internal/job"
	orm "git.ezbuy.me/ezbuy/corsair/digger/service/internal/model"
	trpc "git.ezbuy.me/ezbuy/corsair/digger/service/internal/rpc"
	"git.ezbuy.me/ezbuy/corsair/digger/service/internal/utils"
	"git.ezbuy.me/ezbuy/lib/log"
	ezdb "github.com/ezbuy/ezorm/db"
	mgo "gopkg.in/mgo.v2"
)

var (
	recommendOnce sync.Once
)

func GenRecommendedTicker() {
	tk := time.NewTicker(time.Second * 10)
	for range tk.C {
		if utils.CheckFuncValid(trpc.FunctionType_FunctionTypeRecommendedLong) {
			genRecommendData()
		}
	}
}

func GenRecommendOnce() {
	recommendOnce.Do(func() {
		genRecommendData()
	})
}

func genRecommendData() error {
	sess, col := orm.GPShortLineMgr.GetCol()
	defer sess.Close()

	var data *orm.GPShortLine
	iter := col.Find(ezdb.M{"Disabled": false}).Batch(100).Prefetch(0.25).Iter()
	for iter.Next(&data) {
		getShortRecommendedData(data)
	}

	// 更新任务
	job.UpdateJob(trpc.FunctionType_FunctionTypeLongLine)

	return nil
}

func getShortRecommendedData(data *orm.GPShortLine) error {
	query := ezdb.M{
		"Secucode": data.Secucode,
	}

	gdrenshu, err := orm.GDLongLineMgr.FindOne(query, "-CreateDate")
	if err != nil && err != mgo.ErrNotFound {
		return err
	}

	result := getGPRecommend(data.Secucode)
	decrease := math.Max(float64(data.MDecrease), float64(data.TDecrease))
	if decrease >= float64(data.DecreaseTag)+5 {
		result.State = int32(trpc.RMState_RMStateInProgress)
	} else if decrease >= float64(data.DecreaseTag) {
		result.State = int32(trpc.RMState_RMStateStarted)
	} else if decrease+5 >= float64(data.DecreaseTag) {
		result.State = int32(trpc.RMState_RMStatePrepared)
	} else {
		log.Errorf("condition not met %s|%d", data.Secucode, decrease)
		return nil
	}
	if gdrenshu != nil {
		result.GDDecrease = gdrenshu.GDReduceRatio
	}

	result.RMType = int32(trpc.RMType_RmTypeShort)
	result.Decrease = int32(decrease)
	result.DecreaseTag = data.DecreaseTag
	result.MaxPrice = data.MaxPrice
	result.PresentPrice = data.PresentPrice
	result.RMPrice = calRecommendPrice(result)
	result.UpdateDate = time.Now().Unix()

	if _, err := result.Save(); err != nil {
		log.Errorf("save recommend failed: %s|%q", data.Secucode, err)
		return err
	}

	return nil
}
