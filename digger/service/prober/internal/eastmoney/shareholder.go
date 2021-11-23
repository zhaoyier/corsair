package eastmoney

import (
	// "timevm

	"strings"
	"time"

	"git.ezbuy.me/ezbuy/corsair/digger/rpc/digger"
	"git.ezbuy.me/ezbuy/corsair/digger/service/internal/common/webapi"
	orm "git.ezbuy.me/ezbuy/corsair/digger/service/internal/model"
	log "github.com/Sirupsen/logrus"
	ezdb "github.com/ezbuy/ezorm/db"
	mgo "gopkg.in/mgo.v2"
)

var (
	timeLayout = "2006-01-02 15:04:05"
)

func GetShareholderTicker() {
	tk := time.NewTicker(time.Hour * 2)
	for range tk.C {
		weekday := time.Now().Weekday()
		nowHour := time.Now().Local().Hour()
		if weekday != time.Thursday || weekday != time.Friday { //周
			continue
		}

		if nowHour >= 18 && nowHour <= 24 {
			GetShareholder()
		}

	}
}

func GetShareholder() {
	sess, col := orm.CNSecucodeMgr.GetCol()
	defer sess.Close()

	var secucode *orm.CNSecucode
	iter := col.Find(ezdb.M{}).Batch(100).Prefetch(0.25).Iter()
	for iter.Next(&secucode) {

		shareholder := new(ShareholderResearch)
		code := strings.Replace(secucode.Secucode, ".", "", -1)
		if err := webapi.GetEastmoneyData(digger.EastMoneyType_EastMoneyTypeHolder, code, shareholder); err != nil {
			log.Infof("eastmoney get failed: %+v\n", err)
			continue
		}

		if err := applyShareholder(shareholder.Gdrs); err != nil {
			log.Errorf("apply share holder failed: %s|%q", secucode.Secucode, err)
		}

		if err := applyGDsdlt(shareholder.Sdltgd); err != nil {
			log.Errorf("apply share holder failed: %s|%q", secucode.Secucode, err)
		}

		log.Infof("%s succeed", secucode.Secucode)
	}
}

func applyShareholder(data []Holder) error {
	for _, holder := range data {
		loc, _ := time.LoadLocation("Local")
		tmp, _ := time.ParseInLocation(timeLayout, holder.ENDDATE, loc)
		result, err := orm.GDRenshuMgr.FindOneBySecucodeEndDate(holder.SECUCODE, tmp.Unix())
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

		result.EndDate = tmp.Unix()
		result.HolderTotalNum = int32(holder.HOLDERTOTALNUM)
		if int32(holder.TOTALNUMRATIO) < 100 && int32(holder.TOTALNUMRATIO) > -100 {
			result.TotalNumRatio = int32(holder.TOTALNUMRATIO)
		}
		result.AvgFreeShares = int32(holder.AVGFREESHARES)
		result.AvgFreesharesRatio = int32(holder.AVGFREESHARESRATIO)
		result.HoldFocus = holder.HOLDFOCUS
		result.Price = int32(holder.PRICE)
		result.AvgHoldAmt = int32(holder.AVGHOLDAMT)
		result.HoldRatioTotal = int32(holder.HOLDRATIOTOTAL)
		result.FreeholdRatioTotal = int32(holder.FREEHOLDRATIOTOTAL)
		result.UpdateDate = time.Now().Unix()
		result.CreateDate = time.Now().Unix()

		if _, err := result.Save(); err != nil {
			log.Errorf("save gd renshu failed: %s|%q", holder.SECUCODE, err)
			return err
		}
	}

	return nil
}

func applyGDsdlt(data []Sdltgd) error {
	for _, gd := range data {
		loc, _ := time.LoadLocation("Local")
		tmp, _ := time.ParseInLocation(timeLayout, gd.ENDDATE, loc)
		result, err := orm.GDsdltMgr.FindOneBySecucodeEndDateHolderName(gd.SECUCODE, tmp.Unix(), gd.HOLDERNAME)
		if err != nil && err != mgo.ErrNotFound {
			log.Errorf("find gd renshu failed: %s|%s", gd.SECUCODE, gd.ENDDATE)
			return err
		}
		if result != nil {
			return nil
		}

		result = orm.GDsdltMgr.NewGDsdlt()
		result.Secucode = gd.SECUCODE

		result.EndDate = tmp.Unix()
		result.HolderRank = int32(gd.HOLDERRANK)
		result.HolderName = gd.HOLDERNAME
		result.HolderType = gd.HOLDERTYPE
		result.HoldNum = int32(gd.HOLDNUM)
		result.FreeHoldnumRation = gd.FREEHOLDNUMRATIO
		result.HoldNumChange = gd.HOLDNUMCHANGE
		result.CreateDate = time.Now().Unix()

		if _, err := result.Save(); err != nil {
			log.Errorf("save gd renshu failed: %s|%q", gd.SECUCODE, err)
			return err
		}
	}

	return nil
}
