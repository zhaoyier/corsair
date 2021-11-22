package eastmoney

import (
	// "timevm

	"fmt"
	"strings"
	"time"

	"git.ezbuy.me/ezbuy/corsair/digger/rpc/digger"
	"git.ezbuy.me/ezbuy/corsair/digger/service/internal/common/webapi"
	orm "git.ezbuy.me/ezbuy/corsair/digger/service/internal/model"
	log "github.com/Sirupsen/logrus"
	ezdb "github.com/ezbuy/ezorm/db"
	mgo "gopkg.in/mgo.v2"
)

func GetShareholderTicker() {
	tk := time.NewTicker(time.Hour * 2)
	for range tk.C {
		if time.Now().Local().Hour() != 18 { //周
			continue
		}

		GetShareholder()
	}

	// resp := new(ShareholderResearch)
	// if err := webapi.GetEastmoneyData(digger.EastMoneyTypeEnum.Holder, "SZ002202", resp); err != nil {
	// 	fmt.Printf("eastmoney get failed: %+v\n", err)
	// 	return
	// }
	// fmt.Printf("==>>TODO: %+v\n", resp)

	// resp := new(StockList)
	// if err := webapi.GetEastmoneyCode(1, 20, resp); err != nil {
	// 	fmt.Printf("eastmoney get failed: %+v\n", err)
	// 	return
	// }

	// fmt.Printf("==>>TODO: %+v\n", resp)

}

func GetShareholder() {
	sess, col := orm.CNSecucodeMgr.GetCol()
	defer sess.Close()

	var secucode *orm.CNSecucode
	var shareholder ShareholderResearch
	iter := col.Find(ezdb.M{}).Batch(100).Prefetch(0.25).Iter()
	for iter.Next(&secucode) {
		code := strings.Replace(secucode.Secucode, ".", "", -1)
		if err := webapi.GetEastmoneyData(digger.EastMoneyTypeEnum.Holder, code, shareholder); err != nil {
			fmt.Printf("eastmoney get failed: %+v\n", err)
			continue
		}

		if err := applyShareholder(shareholder.Gdrs); err != nil {
			log.Errorf("apply share holder failed: %s|%q", secucode.Secucode, err)
		}
	}
}

func applyShareholder(data []Holder) error {
	for _, holder := range data {
		result, err := orm.GDRenshuMgr.FindOneBySecucodeEndDate(holder.SECUCODE, holder.ENDDATE)
		if err != nil && err != mgo.ErrNotFound {
			log.Errorf("find gd renshu failed: %s|%s", holder.SECUCODE, holder.ENDDATE)
			return err
		}

		if result != nil {
			return nil
		}

		result = orm.GDRenshuMgr.NewGDRenshu()
		result.Secucode = holder.SECUCODE
		result.SecurityCode = holder.SECURITYCODE
		result.EndDate = holder.ENDDATE
		result.HolderTotalNum = holder.HOLDERTOTALNUM
		result.TotalNumRatio = holder.TOTALNUMRATIO
		result.AvgFreeShares = holder.AVGFREESHARES
		result.AvgFreesharesRatio = holder.AVGFREESHARESRATIO
		result.HoldFocus = holder.HOLDFOCUS
		result.Price = holder.PRICE
		result.AvgHoldAmt = holder.AVGHOLDAMT
		result.HoldRatioTotal = holder.HOLDRATIOTOTAL
		result.FreeholdRatioTotal = holder.FREEHOLDRATIOTOTAL
		result.UpdateDate = time.Now().Unix()
		result.CreateDate = time.Now().Unix()

		if _, err := result.Save(); err != nil {
			log.Errorf("save gd renshu failed: %s|%q", holder.SECUCODE, err)
			return err
		}
	}

	return nil
}
