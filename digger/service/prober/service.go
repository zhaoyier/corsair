package prober

import (
	"git.ezbuy.me/ezbuy/corsair/digger/service/internal/model"
	"git.ezbuy.me/ezbuy/corsair/digger/service/prober/internal/dawdle"
	"git.ezbuy.me/ezbuy/corsair/digger/service/prober/internal/eastmoney"
	"git.ezbuy.me/ezbuy/corsair/digger/service/prober/internal/sina"
	"github.com/ezbuy/ezorm/db"
)

func Start() {
	model.MgoSetup(&db.MongoConfig{
		MongoDB: "mongodb://digger:digger_ppwd@81.69.250.236:12302/digger",
		DBName:  "digger",
	})
	// 300897/300897
	//
	go eastmoney.GetCodeListTicker()
	//go eastmoney.GetCodeListOnce()
	go eastmoney.GetShareholderTicker()
	//go eastmoney.GetShareholderOnce()
	go sina.GetDailyDataTicker()
	// sina.GetDailyDataTmp("SZ.300897")
	// go sina.GetDailyDataOnce()
	go dawdle.GenShareholderTicker()
	// go dawdle.GenShareholderOnce()
	// dawdle.GenShareholderTmp("SZ.003005")

	for {

	}
}
