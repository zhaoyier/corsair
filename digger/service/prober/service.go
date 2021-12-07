package prober

import (
	"os"
	"os/signal"
	"syscall"
	"time"

	"git.ezbuy.me/ezbuy/corsair/digger/service/internal/model"
	"git.ezbuy.me/ezbuy/corsair/digger/service/internal/utils"
	"git.ezbuy.me/ezbuy/corsair/digger/service/prober/internal/dawdle"
	"git.ezbuy.me/ezbuy/corsair/digger/service/prober/internal/eastmoney"
	"github.com/ezbuy/ezorm/db"
)

func Start() {
	tk := time.NewTicker(time.Hour * 1)

	// go dawdle.GenRecommendOnce()
	// go eastmoney.GetCodeListOnce()
	// go eastmoney.GetShareholderOnce()
	// go dawdle.GenLongLineOnce()
	// go dawdle.GenShortLineOnce()
	// dawdle.RecommendedLongTmp("SZ.300614")
	// dawdle.GenLongLineTmp("SZ.300614")
	// dawdle.GenShortLineTmp("SZ.002923")

	for range tk.C {
		if utils.CheckFuncValid() {
			eastmoney.GetCodeListTicker()
			eastmoney.GetShareholderTicker()
			dawdle.GenLongLineTicker()
			dawdle.GenShortLineTicker()
			dawdle.GenRecommendTicker()
		}
	}

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	tk.Stop()
}

func init() {
	model.MgoSetup(&db.MongoConfig{
		MongoDB: "mongodb://digger:digger_ppwd@81.69.250.236:12302/digger",
		DBName:  "digger",
	})
}
