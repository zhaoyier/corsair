package zwadmin

import (
	"net/http"
	"time"

	orm "git.ezbuy.me/ezbuy/corsair/digger/service/internal/model"
	trpc "git.ezbuy.me/ezbuy/corsair/digger/service/internal/rpc"
	"git.ezbuy.me/ezbuy/corsair/digger/service/internal/utils"

	log "github.com/Sirupsen/logrus"
	ezdb "github.com/ezbuy/ezorm/db"
	"github.com/gin-gonic/gin"
)

func GetDailyList(in *gin.Context) {
	var req trpc.GetDailyListReq
	resp := &trpc.GetDailyListResp{
		Code: 21000,
		Data: &trpc.GPDailyData{
			Items: make([]*trpc.GPDailyItem, 0),
		},
	}
	if err := in.BindJSON(&req); err != nil {
		in.JSON(http.StatusBadRequest, resp)
		return
	}

	query := ezdb.M{}
	var sortField string
	if req.GetName() != "" {
		query["Name"] = req.GetName()
	} else if req.GetSecucode() != "" {
		query["Secucode"] = utils.GetSecucode(req.GetSecucode())
	} else if req.GetStartDate() > 0 && req.GetEndDate() > 0 {
		sortField = "CreateDate"
		query["CreateDate"] = ezdb.M{"$gte": req.GetStartDate() / 1000, "$lte": req.GetEndDate() / 1000}
	} else if req.GetDecrease() > 0 {
		sortField = "-PRise"
		query["PRise"] = ezdb.M{"$gte": req.GetDecrease()}
	} else if req.GetDecrease() < 0 {
		sortField = "PRise"
		query["PRise"] = ezdb.M{"$lte": req.GetDecrease()}
	} else {
		query["CreateDate"] = getCreateDate()
	}

	results, err := orm.GPDailyMgr.Find(query, int(req.GetLimit()), int(req.GetOffset()), sortField)
	if err != nil {
		log.Errorf("get prompt buy failed: %q", err)
		in.JSON(http.StatusForbidden, resp)
		return
	}

	total := orm.GPDailyMgr.Count(query)
	resp.Data.Total = int32(total)

	for _, result := range results {
		resp.Data.Items = append(resp.Data.Items, &trpc.GPDailyItem{
			Name:       result.Name,
			Secucode:   result.Secucode,
			Opening:    result.Opening,
			Closing:    result.Closing,
			Prise:      utils.TruncateFloat(result.Rise),
			Turnover:   result.Turnover,
			Business:   result.Business,
			Liangbi:    result.Liangbi,
			MaxPrice:   result.MaxPrice,
			MinPrice:   result.MinPrice,
			Market:     result.Market / 100000000,
			Traded:     result.Traded / 100000000,
			BookRatio:  result.BookRatio,
			CreateDate: time.Unix(result.CreateDate, 0).Format("2006-01-02"),
		})
	}

	resp.Code = 20000
	in.JSON(http.StatusOK, resp)
}

func getCreateDate() int64 {
	now := time.Now()

	if now.Weekday() == time.Saturday {
		now = now.AddDate(0, 0, -1)
	} else if now.Weekday() == time.Sunday {
		now = now.AddDate(0, 0, -2)
	} else if now.Weekday() == time.Monday && now.Hour() <= 17 {
		now = now.AddDate(0, 0, -3)
	} else if now.Hour() <= 17 {
		now = now.AddDate(0, 0, -1)
	}

	timeStr := now.Format("2006-01-02")
	t, _ := time.ParseInLocation("2006-01-02", timeStr, time.Local)
	return t.Unix()
}
