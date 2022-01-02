package longline

import (
	"net/http"
	"sort"
	"strings"
	"sync"
	"time"

	orm "git.ezbuy.me/ezbuy/corsair/digger/service/internal/model"
	trpc "git.ezbuy.me/ezbuy/corsair/digger/service/internal/rpc"

	log "github.com/Sirupsen/logrus"
	ezdb "github.com/ezbuy/ezorm/db"
	"github.com/gin-gonic/gin"
)

func GetLongLineList(in *gin.Context) {
	var req trpc.GetLongLineListReq
	resp := &trpc.GetLongLineListResp{
		Code: 21000,
		Data: &trpc.LongLineData{
			Items: make([]*trpc.LongLineItem, 0),
		},
	}
	if err := in.BindJSON(&req); err != nil {
		in.JSON(http.StatusBadRequest, resp)
		return
	}

	query := ezdb.M{
		"Disabled": false,
	}

	if req.GetSecucode() != "" {
		query["Secucode"] = req.GetSecucode()
	}

	sortField := "-GDReduceRatio"
	if req.GetGdRatio() > 0 {
		query["GDReduceRatio"] = ezdb.M{"$gte": req.GetGdRatio()}
	}

	if req.GetGdRatio() < 0 {
		sortField = "GDReduceRatio"
		query["GDReduceRatio"] = ezdb.M{"$lte": req.GetGdRatio()}
	}

	results, err := orm.GDLongLineMgr.Find(query, int(req.GetLimit()), int(req.GetOffset()), sortField)
	if err != nil {
		log.Errorf("get long line failed: %q", err)
		in.JSON(http.StatusForbidden, resp)
		return
	}

	total := orm.GDLongLineMgr.Count(query)
	resp.Data.Total = int32(total)

	wg := sync.WaitGroup{}
	wg.Add(len(results))

	for _, result := range results {
		result := result
		go func(wg *sync.WaitGroup) {
			resp.Data.Items = append(resp.Data.Items, &trpc.LongLineItem{
				Secucode:     result.Secucode,
				Name:         result.Name,
				Price:        result.CumulantPrice,
				Focus:        result.CumulantFocus,
				Date:         result.CumulantDate,
				GdRatio:      result.GDReduceRatio,
				RatioStr:     result.CumulantRatio,
				CreateDate:   time.Unix(result.CreateDate, 0).Format("2006-01-02 15:04:05"),
				PresentPrice: getPresentPrice(result.Secucode),
				ValueIndex:   result.ValueIndex,
			})
			wg.Done()
		}(&wg)
	}

	wg.Wait()

	if req.GetGdRatio() >= 0 {
		sort.Slice(resp.Data.Items, func(i, j int) bool {
			return resp.Data.Items[i].GdRatio > resp.Data.Items[j].GdRatio
		})
	} else {
		sort.Slice(resp.Data.Items, func(i, j int) bool {
			return resp.Data.Items[i].GdRatio < resp.Data.Items[j].GdRatio
		})
	}

	resp.Code = 20000
	in.JSON(http.StatusOK, resp)
}

func getPresentPrice(secucode string) float64 {
	secucode = strings.Split(secucode, ".")[1]
	result, err := orm.GPDailyMgr.FindOne(ezdb.M{"Secucode": secucode}, "-CreateDate")
	if err != nil {
		return 0
	}
	return result.Closing
}
