package eastmoney

import (
	// "timevm

	"fmt"

	"git.ezbuy.me/ezbuy/corsair/digger/service/internal/common/webapi"
)

func GetShareholder() {
	// tk := time.NewTicker(time.Hour * 2)
	// for range tk.C {

	// }

	// resp := new(ShareholderResearch)
	// if err := webapi.GetEastmoneyData(digger.EastMoneyTypeEnum.Holder, "SZ002202", resp); err != nil {
	// 	fmt.Printf("eastmoney get failed: %+v\n", err)
	// 	return
	// }
	// fmt.Printf("==>>TODO: %+v\n", resp.Gdrs)

	resp := new(StockList)
	if err := webapi.GetEastmoneyCode(1, 20, resp); err != nil {
		fmt.Printf("eastmoney get failed: %+v\n", err)
		return
	}

	fmt.Printf("==>>TODO: %+v\n", resp)

}
