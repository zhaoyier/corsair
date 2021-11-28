package prober

import (
	"git.ezbuy.me/ezbuy/corsair/digger/service/internal/model"
	"git.ezbuy.me/ezbuy/corsair/digger/service/prober/internal/dawdle"
	"git.ezbuy.me/ezbuy/corsair/digger/service/prober/internal/eastmoney"

	// "git.ezbuy.me/ezbuy/corsair/digger/service/prober/internal/sina"
	"github.com/ezbuy/ezorm/db"
)

func Start() {
	model.MgoSetup(&db.MongoConfig{
		MongoDB: "mongodb://digger:digger_ppwd@81.69.250.236:12302/digger",
		DBName:  "digger",
	})

	go eastmoney.GetCodeListTicker()
	// go eastmoney.GetCodeListOnce()
	go eastmoney.GetShareholderTicker()
	//go eastmoney.GetShareholderOnce()
	go dawdle.GenShareholderTicker()
	go dawdle.GenShareholderOnce()

	// dawdle.GenShareholderTmp("SZ.300469")
	// go sina.GetDailyDataTicker()
	// go sina.GetDailyDataOnce()
	// sina.GetDailyDataTmp("SZ.300897")

	for {

	}
}
