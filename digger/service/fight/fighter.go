package fight

import (
	"os"
	"os/signal"
	"syscall"

	"git.ezbuy.me/ezbuy/corsair/digger/service/internal/model"
	"github.com/ezbuy/ezorm/db"

	"git.ezbuy.me/ezbuy/corsair/digger/service/internal/dawdle"
)

func Start() {
	// dawdle.GenStatFundFlowTmp("SZ.300632")
	// dawdle.GenStatFundFlowTmp("SZ.003001")
	// dawdle.GenStatFundFlowTmp("SH.603212")
	dawdle.GenStatFundFlowOnce()

	go func() {
		quit := make(chan os.Signal, 1)
		signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
		<-quit
	}()

}

func init() {
	model.MgoSetup(&db.MongoConfig{
		MongoDB: "mongodb://digger:digger_ppwd@81.69.250.236:12302/digger",
		DBName:  "digger",
	})

	dawdle.InitConf()
}
