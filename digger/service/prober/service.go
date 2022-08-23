package prober

import (
	"os"
	"os/signal"
	"syscall"
	"time"

	"git.ezbuy.me/ezbuy/corsair/digger/service/internal/model"
	trpc "git.ezbuy.me/ezbuy/corsair/digger/service/internal/rpc"
	"git.ezbuy.me/ezbuy/corsair/digger/service/internal/utils"
	"github.com/ezbuy/ezorm/db"

	"git.ezbuy.me/ezbuy/corsair/digger/service/internal/dawdle"
	"git.ezbuy.me/ezbuy/corsair/digger/service/internal/eastmoney"
)

func Start() {
	tk := time.NewTicker(time.Minute * 1000) //TODOZ

	// go eastmoney.GetCodeListOnce()
	// go eastmoney.GetShareholderOnce()
	// go dawdle.GenLongLineOnce()
	// go dawdle.GenShortLineOnce()
	// go dawdle.GenRecommendOnce()

	// dawdle.GenRecommendTmp("SZ.002374")
	// dawdle.GenLongLineTmp("SH.603136")
	// dawdle.GenShortLineTmp("SZ.300741")
	// dawdle.GenZhouQiOnce()
	// eastmoney.GetFundFlowOnce()
	// eastmoney.GetFundFlowTemp("SZ.301100")
	// dawdle.GenWaterfallOnce()
	// dawdle.GenWaterfallTmp("SZ.300991")

	// dawdle.GenStatFundFlowTmp("SZ.300632")
	// dawdle.GenStatFundFlowTmp("SZ.003001")
	// dawdle.GenStatFundFlowTmp("SH.603212")
	// dawdle.GenStatFundFlowOnce()

	for range tk.C {
		if utils.CheckFuncValid(trpc.FunctionType_FunctionTypeCodeList) {
			eastmoney.GetCodeListTicker()
		}
		if utils.CheckFuncValid(trpc.FunctionType_FunctionTypeShareholder) {
			eastmoney.GetShareholderTicker()
		}
		if utils.CheckFuncValid(trpc.FunctionType_FunctionTypeFundFlow) {
			eastmoney.GetFundFlowTicker()
		}
		if utils.CheckFuncValid(trpc.FunctionType_FunctionTypeLongLine) {
			dawdle.GenLongLineTicker()
		}
		if utils.CheckFuncValid(trpc.FunctionType_FunctionTypeShortLine) {
			dawdle.GenShortLineTicker()
		}
		if utils.CheckFuncValid(trpc.FunctionType_FunctionTypeRecommend) {
			dawdle.GenRecommendTicker()
		}
		// 关注
		if utils.CheckFuncValid(trpc.FunctionType_FunctionTypeFocus) {
			dawdle.GenFocusStateTicker()
		}
		// 周期
		if utils.CheckFuncValid(trpc.FunctionType_FunctionTypeZhouQi) {
			dawdle.GenZhouQiTicker()
		}
		// 瀑布线
		if utils.CheckFuncValid(trpc.FunctionType_FunctionTypeWaterfall) {
			dawdle.GenWaterfallTicker()
		}
		// 统计流入
		if utils.CheckFuncValid(trpc.FunctionType_FunctionTypeStatFundFlow) {
			dawdle.GenStatFundFlowTicker()
		}

		dawdle.UpdatePresentPrice()
	}

	go func(tk *time.Ticker) {
		quit := make(chan os.Signal, 1)
		signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
		<-quit
		tk.Stop()
	}(tk)

}

func init() {
	model.MgoSetup(&db.MongoConfig{
		MongoDB: "mongodb://digger:digger_ppwd@81.69.250.236:12302/digger",
		DBName:  "digger",
	})

	dawdle.InitConf()
}
