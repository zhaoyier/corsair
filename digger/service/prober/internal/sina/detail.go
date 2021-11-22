package sina

import (
	"git.ezbuy.me/ezbuy/corsair/digger/service/internal/common/webapi"
	log "github.com/Sirupsen/logrus"
)

func GetDailyDataTicker() {

}

func GetDailyData() {
	result, err := webapi.GetSinaDayDetail("sz002307")
	if err != nil {
		log.Errorf("daily data failed: %q", err)
		return
	}
	log.Infof("==>>TODO 101: %+v", result)

}
