package dawdle

import (
	"fmt"
	"math"
	"strings"
	"sync"
	"time"

	orm "git.ezbuy.me/ezbuy/corsair/digger/service/internal/model"
	trpc "git.ezbuy.me/ezbuy/corsair/digger/service/internal/rpc"
	"git.ezbuy.me/ezbuy/corsair/digger/service/internal/utils"
	log "github.com/Sirupsen/logrus"
	ezdb "github.com/ezbuy/ezorm/db"
)

var (
	genShortLineOnce sync.Once
)

func GenShortLineTicker() {
	tk := time.NewTicker(time.Minute * 10)
	for range tk.C {
		if utils.CheckFuncValid(trpc.FunctionType_FunctionTypeShortLine) {
			GenShortLineData()

		}
	}
}

func GenShortLineOnce() {
	genShortLineOnce.Do(func() {
		GenShortLineData()
	})
}

func GenShortLineTmp(secucode string) {
	getShortLineData(secucode)
}

func GenShortLineData() {
	sess, col := orm.CNSecucodeMgr.GetCol()
	defer sess.Close()

	var secucode *orm.CNSecucode
	iter := col.Find(ezdb.M{"Disabled": false}).Batch(100).Prefetch(0.25).Iter()
	for iter.Next(&secucode) {
		if err := getShortLineData(secucode.Secucode); err != nil {
			log.Infof("gen short line failed: %s|%q", secucode.Secucode, err)
		}
	}
	// 更新任务
	utils.UpdateFunction(trpc.FunctionType_FunctionTypeShortLine)
}

// 最近60日数据
func getShortLineData(secucode string) error {
	codes := strings.Split(secucode, ".")
	data := getGPRecommend(secucode)
	if err := getLastDecrease(data); err != nil {
		return err
	}
	data.DecreaseTag = getDecreaseValue(codes[1])
	if data.Decrease < data.DecreaseTag {
		return nil
	}
	data.RMPrice = calRecommendPrice(data)
	data.RMType = int32(trpc.RMType_RmTypeShort)

	if err := applyGPRecommend(data); err != nil {
		log.Errorf("apply recommend failed: %s|%q", secucode, err)
		return err
	}
	return nil
}

func getLastDecrease(data *orm.GPRecommend) error {
	secucode := utils.GetSecucode(data.Secucode)
	tm := time.Now().AddDate(0, -2, 0).Unix()
	query := ezdb.M{
		"Secucode":   secucode,
		"CreateDate": ezdb.M{"$gte": tm},
	}

	results, err := orm.GPDailyMgr.FindAll(query, "-CreateDate")
	if err != nil {
		log.Errorf("query daily failed: %s|%q", secucode, err)
		return err
	}

	var createDate int64
	for idx, result := range results {
		if idx == 0 {
			data.PresentPrice = math.Min(result.Closing, result.MinPrice)
		}

		if result.MaxPrice > data.MaxPrice {
			data.MaxDay++
			data.MaxPrice = result.MaxPrice
			createDate = result.CreateDate
		}
	}
	dateStr := time.Unix(createDate, 0).Format("2006-01-02")
	if data.MaxDay == 1 {
		return fmt.Errorf("invalid data:%s", data.Secucode)
	}
	data.MaxPrice = utils.Decimal(data.MaxPrice)
	data.PresentPrice = utils.Decimal(data.PresentPrice)
	data.Decrease = utils.DecreasePercent(data.MaxPrice, data.PresentPrice)
	data.DecreaseDay = fmt.Sprintf("%d&%s", data.MaxDay, dateStr)
	return nil
}

func calRecommendPrice(data *orm.GPRecommend) string {
	price := data.MaxPrice
	tag := utils.Decimal(1 - utils.GetPercentum(data.DecreaseTag))
	max, min := utils.Decimal(tag+0.05), utils.Decimal(tag-0.05)
	return fmt.Sprintf("%.1f(1)-%.1f(2)-%.1f(3)", math.Floor(price*max), math.Ceil(price*tag), math.Floor(price*min))
}

func getDecreaseValue(secucode string) int32 {
	query := ezdb.M{
		"Secucode": secucode,
	}
	result, err := orm.GPDailyMgr.FindOne(query, "-CreateDate")
	if err != nil {
		return GPShortDecrease
	}
	if result.Market > int64(math.Pow10(11)*2) {
		return GPShortDecrease - 8
	} else if result.Market > int64(math.Pow10(11)) {
		return GPShortDecrease - 6
	} else if result.Market > int64(math.Pow10(10)*5) {
		return GPShortDecrease - 4
	} else if result.Market > int64(math.Pow10(10)) {
		return GPShortDecrease - 3
	} else if result.Market > int64(math.Pow10(9)) {
		return GPShortDecrease - 3
	}
	return GPShortDecrease
}
