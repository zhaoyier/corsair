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

	//
	go eastmoney.GetCodeListTicker()
	//
	go eastmoney.GetShareholderTicker()
	//
	// go sina.GetDailyDataTicker()
	go sina.GetDailyData()
	//
	go dawdle.GenShareholderTicker()

	for {

	}
}
