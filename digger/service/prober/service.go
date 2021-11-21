package prober

import (
	"git.ezbuy.me/ezbuy/corsair/digger/service/internal/model"
	"git.ezbuy.me/ezbuy/corsair/digger/service/prober/internal/eastmoney"
	"github.com/ezbuy/ezorm/db"
)

func Start() {
	model.MgoSetup(&db.MongoConfig{
		MongoDB: "mongodb://digger:digger_ppwd@81.69.250.236:12302/digger",
		DBName:  "digger",
	})

	eastmoney.GetCodeList()
}
