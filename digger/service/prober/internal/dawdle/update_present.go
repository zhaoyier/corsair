package dawdle

import (
	"strings"
	"time"

	"git.ezbuy.me/ezbuy/corsair/digger/service/internal/common/webapi"
	orm "git.ezbuy.me/ezbuy/corsair/digger/service/internal/model"
	"git.ezbuy.me/ezbuy/corsair/digger/service/internal/utils"
	log "github.com/Sirupsen/logrus"
	ezdb "github.com/ezbuy/ezorm/db"
)

func UpdatePresentPrice() {
	now := time.Now()
	if now.Hour() < 9 || now.Hour() > 15 {
		return
	}

	iterGPRecommend()
}

func iterGPRecommend() {
	query := ezdb.M{"Disabled": false}

	results, err := orm.GPRecommendMgr.FindAll(query)
	if err != nil {
		log.Infof("get recommend failed: %q", err)
		return
	}

	for _, result := range results {
		price := getPresentPrice(result.Secucode)
		updateRecommendPrice(result.Secucode, price)
	}
}

func getPresentPrice(secucode string) float64 {
	secucode = strings.ToLower(secucode)
	secucode = strings.Replace(secucode, ".", "", -1)

	results, err := webapi.GetSinaDayDetail(secucode)
	if err != nil {
		log.Errorf("get present price failed: %s|%q", secucode, err)
		return 0
	}
	price := results[3]
	return utils.String2Float64(price)
}

func updateRecommendPrice(secucode string, price float64) error {
	if price == 0 {
		log.Infof("update price invalid: %s|%.1f", secucode, price)
		return nil
	}

	result, err := orm.GPRecommendMgr.FindOneBySecucodeDisabled(secucode, false)
	if err != nil {
		log.Infof("get recommend failed: %q", err)
		return err
	}
	result.PresentPrice = price
	result.UpdateBy = "update price"
	result.UpdateDate = time.Now().Unix()

	if _, err := result.Save(); err != nil {
		log.Infof("update recommend failed: %q", err)
		return err
	}
	return nil
}
