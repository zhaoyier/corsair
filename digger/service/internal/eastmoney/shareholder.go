package eastmoney

import (
	// "timevm

	"strings"
	"sync"
	"time"

	"git.ezbuy.me/ezbuy/corsair/digger/service/internal/common/webapi"
	"git.ezbuy.me/ezbuy/corsair/digger/service/internal/job"
	orm "git.ezbuy.me/ezbuy/corsair/digger/service/internal/model"
	trpc "git.ezbuy.me/ezbuy/corsair/digger/service/internal/rpc"
	"git.ezbuy.me/ezbuy/corsair/digger/service/internal/utils"
	log "github.com/Sirupsen/logrus"
	ezdb "github.com/ezbuy/ezorm/db"
	mgo "gopkg.in/mgo.v2"
)

var (
	getCodeListOnce    sync.Once
	getShareholderOnce sync.Once
	timeLayout         = "2006-01-02 15:04:05"
)

func GetShareholderTicker() {
	GetShareholder()
	job.UpdateJob(trpc.FunctionType_FunctionTypeShareholder)
}

func GetShareholderOnce() {
	getShareholderOnce.Do(func() {
		GetShareholder()
	})
}

func GetShareholder() {
	sess, col := orm.CNSecucodeMgr.GetCol()
	defer sess.Close()

	var secucode *orm.CNSecucode
	iter := col.Find(ezdb.M{"Disabled": false}).Batch(100).Prefetch(0.25).Iter()
	for iter.Next(&secucode) {
		shareholder := new(ShareholderResearch)
		code := strings.Replace(secucode.Secucode, ".", "", -1)
		if err := webapi.GetEastmoneyData(trpc.EastMoneyType_EastMoneyTypeHolder, code, shareholder); err != nil {
			log.Infof("eastmoney get failed: %s|%+v\n", code, err)
			continue
		}

		if err := applyShareholder(shareholder.Gdrs); err != nil && mgo.IsDup(err) {
			log.Errorf("apply share holder failed: %s|%q", secucode.Secucode, err)
			return
		}

		if err := applyGDsdlt(shareholder.Sdltgd); err != nil && mgo.IsDup(err) {
			log.Errorf("apply share holder failed: %s|%q", secucode.Secucode, err)
			return
		}

		log.Infof("%s succeed", secucode.Secucode)
	}
}

func applyShareholder(data []Holder) error {
	for _, holder := range data {
		loc, _ := time.LoadLocation("Local")
		tmp, _ := time.ParseInLocation(timeLayout, holder.ENDDATE, loc)
		codes := strings.Split(holder.SECUCODE, ".")
		if len(codes) < 2 {
			continue
		}
		secucode := codes[1] + "." + codes[0]
		result, err := orm.GDRenshuMgr.FindOneBySecucodeEndDate(secucode, tmp.Unix())
		if err != nil && err != mgo.ErrNotFound {
			log.Errorf("find gd renshu failed: %s|%s", secucode, holder.ENDDATE)
			return err
		}
		if result != nil {
			return nil
		}

		result = orm.GDRenshuMgr.NewGDRenshu()
		result.Secucode = secucode
		result.SecurityCode = holder.SECURITYCODE

		result.EndDate = tmp.Unix()
		result.HolderTotalNum = holder.HOLDERTOTALNUM
		if int32(holder.TOTALNUMRATIO) < 100 && int32(holder.TOTALNUMRATIO) > -100 {
			result.TotalNumRatio = holder.TOTALNUMRATIO
		}
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
			log.Errorf("save gd renshu failed: %s|%q", secucode, err)
			return err
		}
	}

	return nil
}

func applyGDsdlt(data []Sdltgd) error {
	for _, gd := range data {
		loc, _ := time.LoadLocation("Local")
		codes := strings.Split(gd.SECUCODE, ".")
		if len(codes) < 2 {
			continue
		}
		secucode := codes[1] + "." + codes[0]
		tmp, _ := time.ParseInLocation(timeLayout, gd.ENDDATE, loc)
		result, err := orm.GDsdltMgr.FindOneBySecucodeEndDateHolderName(secucode, tmp.Unix(), gd.HOLDERNAME)
		if err != nil && err != mgo.ErrNotFound {
			log.Errorf("find gd renshu failed: %s|%s", secucode, gd.ENDDATE)
			return err
		}
		if result != nil {
			return nil
		}

		result = orm.GDsdltMgr.NewGDsdlt()
		result.Secucode = secucode
		result.EndDate = tmp.Unix()
		result.HolderRank = int32(gd.HOLDERRANK)
		result.HolderName = gd.HOLDERNAME
		result.HolderType = gd.HOLDERTYPE
		result.HoldNum = int32(gd.HOLDNUM)
		result.FreeHoldnumRation = utils.Decimal(gd.FREEHOLDNUMRATIO)
		result.HoldNumChange = gd.HOLDNUMCHANGE
		result.CreateDate = time.Now().Unix()

		if _, err := result.Save(); err != nil {
			log.Errorf("save gd renshu failed: %s|%q", gd.SECUCODE, err)
			return err
		}
	}

	return nil
}
