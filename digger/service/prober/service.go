package prober

import (
	"time"

	"git.ezbuy.me/ezbuy/corsair/digger/service/internal/model"
	"git.ezbuy.me/ezbuy/corsair/digger/service/prober/internal/dawdle"
	"git.ezbuy.me/ezbuy/corsair/digger/service/prober/internal/eastmoney"

	// "git.ezbuy.me/ezbuy/corsair/digger/service/prober/internal/sina"

	"github.com/ezbuy/ezorm/db"
)

func Start() {
	tk := time.NewTicker(time.Minute * 2)

	model.MgoSetup(&db.MongoConfig{
		MongoDB: "mongodb://digger:digger_ppwd@81.69.250.236:12302/digger",
		DBName:  "digger",
	})

	go eastmoney.GetCodeListTicker()
	go eastmoney.GetShareholderTicker()
	go dawdle.GenLongLineTicker()
	go dawdle.GenShortLineTicker()
	go dawdle.GenRecommendedTicker()
	// go dawdle.GenRecommendOnce()
	// go eastmoney.GetCodeListOnce()
	// go eastmoney.GetShareholderOnce()
	// go dawdle.GenLongLineOnce()
	// go dawdle.GenShortLineOnce()
	// go sina.GetDailyDataTicker()
	// go sina.GetDailyDataOnce()
	// sina.GetDailyDataTmp("SZ.300897")
	// dawdle.RecommendedLongTmp("SZ.300614")
	// dawdle.GenLongLineTmp("SZ.300614")
	// dawdle.GenShortLineTmp("SZ.002923")
	// utils.Decimal("ddd")

	for range tk.C {

	}
}
