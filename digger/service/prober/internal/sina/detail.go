package sina

import (
	"fmt"
	"strings"
	"time"

	"git.ezbuy.me/ezbuy/corsair/digger/service/internal/common/webapi"
	orm "git.ezbuy.me/ezbuy/corsair/digger/service/internal/model"
	log "github.com/Sirupsen/logrus"
	ezdb "github.com/ezbuy/ezorm/db"
	mgo "gopkg.in/mgo.v2"
)

func GetDailyDataTicker() {
	tk := time.NewTicker(time.Hour * 2)
	for range tk.C {
		nowHour := time.Now().Local().Hour()
		if nowHour >= 18 && nowHour <= 24 { //周
			GetDailyData()
		}

	}
}

func GetDailyData() {
	now := time.Now()
	date := fmt.Sprintf("%d-%d-%d", now.Year(), now.Month(), now.Day())
	sess, col := orm.CNSecucodeMgr.GetCol()
	defer sess.Close()

	var secucode *orm.CNSecucode
	iter := col.Find(ezdb.M{}).Batch(100).Prefetch(0.25).Iter()
	for iter.Next(&secucode) {
		code := strings.Replace(secucode.Secucode, ".", "", -1)
		result, err := webapi.GetSinaDayDetail(code)
		if err != nil {
			log.Errorf("get sina daily failed: %s|%q", secucode.Secucode, err)
			continue
		}

		if err := applyDaily(secucode.Secucode, date, result); err != nil {
			log.Errorf("get sina daily failed: %s|%q", secucode.Secucode, err)
			continue
		}
	}
}

func applyDaily(secucode, date string, data []string) error {
	result, err := orm.DailyMgr.FindOneBySecucodeEndDate(secucode, date)
	if err != nil && err != mgo.ErrNotFound {
		log.Errorf("find gd renshu failed: %s|%s", secucode, date)
		return err
	}
	if result != nil {
		return nil
	}

	result = orm.DailyMgr.NewDaily()
	result.Secucode = secucode
	result.EndDate = date
	result.Price = data[3]
	result.CreateDate = time.Now().Unix()

	if _, err := result.Save(); err != nil {
		log.Errorf("save gd renshu failed: %s|%q", secucode, err)
		return err
	}

	return nil
}
