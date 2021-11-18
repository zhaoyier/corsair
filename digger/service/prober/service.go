package prober

import (
	"git.ezbuy.me/ezbuy/corsair/digger/service/prober/internal/eastmoney"
)

func Start() {
	eastmoney.GetShareholder()
}
