package dawdle

import (
	"fmt"
	"math"
	"sync"
	"time"

	"git.ezbuy.me/ezbuy/corsair/digger/service/internal/job"
	orm "git.ezbuy.me/ezbuy/corsair/digger/service/internal/model"
	trpc "git.ezbuy.me/ezbuy/corsair/digger/service/internal/rpc"
	"git.ezbuy.me/ezbuy/corsair/digger/service/internal/utils"
	log "github.com/Sirupsen/logrus"
	ezdb "github.com/ezbuy/ezorm/db"
	mgo "gopkg.in/mgo.v2"
)

var (
	recommendOnce sync.Once
)

func GenRecommendTicker() {
	genRecommendData()
	// 更新任务
	job.UpdateJob(trpc.FunctionType_FunctionTypeRecommend)
}

func GenRecommendOnce() {
	recommendOnce.Do(func() {
		genRecommendData()
	})
}

func GenRecommendTmp(secucode string) {
	result, err := orm.GPShortLineMgr.FindOne(ezdb.M{"Secucode": secucode}, "-CreateDate")
	if err != nil {
		log.Errorf("query short line failed: %s|%q", secucode, err)
		return
	}
	getShortRecommendedData(result)
}

func genRecommendData() error {
	sess, col := orm.GPShortLineMgr.GetCol()
	defer sess.Close()

	var data *orm.GPShortLine
	iter := col.Find(ezdb.M{}).Batch(100).Prefetch(0.25).Iter()
	for iter.Next(&data) {
		getShortRecommendedData(data)
	}

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
	} else if decrease >= float64(data.DecreaseTag)-5 {
		result.State = int32(trpc.RMState_RMStatePrepared)
	} else {
		result.State = int32(trpc.RMState_RMStateUnknown)
		return nil
	}
	if gdrenshu != nil {
		result.GDDecrease = gdrenshu.GDReduceRatio
	}

	result.Name = data.Name
	result.RMType = int32(trpc.RMType_RmTypeShort)
	result.PDecrease = int32(decrease)
	result.DecreaseTag = data.DecreaseTag
	result.MaxPrice = data.MaxPrice
	result.MaxPDay = data.MaxPDay
	result.PresentPrice = data.PresentPrice
	result.RMPrice = calRecommendPrice(result)
	result.UpdateDate = time.Now().Unix()
	result.RMIndex = getRecommendIndex(result)
	result.Disabled = data.Disabled

	log.Infof("==>>335:%+v", result, data)

	if _, err := result.Save(); err != nil {
		log.Errorf("save recommend failed: %s|%q", data.Secucode, err)
		return err
	}

	return nil
}

func getRecommendIndex(data *orm.GPRecommend) int32 {
	var rate, gd int32
	if data.State == int32(trpc.RMState_RMStatePrepared) || data.State == int32(trpc.RMState_RMStateInProgress) {
		rate = 70
	}

	if data.GDDecrease != 0 {
		num := data.GDDecrease / -10
		gd = num * 6
	}

	if gd > 30 {
		gd = 30
	}

	if gd > 0 {
		rate += gd
	}

	return rate
}

func calRecommendPrice(data *orm.GPRecommend) string {
	price := data.MaxPrice

	tag := utils.Decimal(1 - utils.GetPercentum(data.DecreaseTag))
	// log.Infof("==>>TODO 311: %+v|%+v", price, tag)
	max, per, min := utils.Decimal(tag+0.03), utils.Decimal(tag), utils.Decimal(tag-0.05)
	// log.Infof("==>>TODO 312: %+v|%+v|%+v", max, per, min)
	return fmt.Sprintf("%.1f(1)-%.1f(2)-%.1f(3)", math.Floor(price*max), math.Floor(price*per), math.Floor(price*min))
}

func getGPRecommend(secucode string) *orm.GPRecommend {
	result, err := orm.GPRecommendMgr.FindOneBySecucodeDisabled(secucode, false)
	if err != nil || result == nil {
		result = orm.GPRecommendMgr.NewGPRecommend()
		result.Secucode = secucode
		result.CreateDate = time.Now().Unix()
	}
	return result
}
