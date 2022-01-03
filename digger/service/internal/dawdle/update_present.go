package dawdle

import (
	"time"

	orm "git.ezbuy.me/ezbuy/corsair/digger/service/internal/model"
	"git.ezbuy.me/ezbuy/corsair/digger/service/internal/request"
	"git.ezbuy.me/ezbuy/corsair/digger/service/internal/utils"
	log "github.com/Sirupsen/logrus"
	ezdb "github.com/ezbuy/ezorm/db"
)

func UpdatePresentPrice() {
	now := time.Now()
	if now.Hour() < 9 || now.Hour() > 15 {
		return
	}

	// 更新推荐的价格
	iterGPRecommend()
	// 更新关注的价格
	iterGPFocusPresentPrice()
}

func iterGPRecommend() {
	query := ezdb.M{"Disabled": false}

	results, err := orm.GPRecommendMgr.FindAll(query)
	if err != nil {
		log.Infof("get recommend failed: %q", err)
		return
	}

	for _, result := range results {
		price := request.GetSinaDayPrice(result.Secucode)
		updateRecommendPrice(result.Secucode, price)
	}
}

func iterGPFocusPresentPrice() {
	query := ezdb.M{}

	results, err := orm.GPFocusMgr.FindAll(query)
	if err != nil {
		log.Infof("get recommend failed: %q", err)
		return
	}

	for _, result := range results {
		price := request.GetSinaDayPrice(result.Secucode)
		updateFocusPrice(result, price)
	}
}

func updateRecommendPrice(secucode string, price float64) error {
	if price == 0 {
		log.Infof("update price invalid: %s|%.1f", secucode, price)
		return nil
	}
	result, err := orm.GPShortLineMgr.FindOne(ezdb.M{"Secucode": secucode}, "-CreateDate")
	if err != nil {
		log.Infof("get short line failed: %q", err)
		return err
	}
	result.MDecrease = utils.DecreasePercent(result.MaxPrice, price)
	result.TDecrease = utils.DecreasePercent(result.MaxPrice, price)
	if result.MinPrice < price {
		return nil
	} else {
		result.MinPrice = price
	}

	if err := genShortRecommendedData(result, false); err != nil {
		log.Infof("gen recommend failed: %s|%q", secucode, err)
		return err
	}

	return nil
}

func updateFocusPrice(data *orm.GPFocus, price float64) error {
	data.PresentPrice = price
	data.UpdateDate = time.Now().Unix()
	if _, err := data.Save(); err != nil {
		log.Errorf("update gp focus failed: %s|%q", data.Secucode, price)
		return err
	}
	return nil
}
