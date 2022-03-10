package zwadmin

import (
	"fmt"
	"net/http"
	"sort"
	"sync"

	orm "git.ezbuy.me/ezbuy/corsair/digger/service/internal/model"
	"git.ezbuy.me/ezbuy/corsair/digger/service/internal/request"
	trpc "git.ezbuy.me/ezbuy/corsair/digger/service/internal/rpc"
	log "github.com/Sirupsen/logrus"
	ezdb "github.com/ezbuy/ezorm/db"
	"github.com/gin-gonic/gin"
)

func GetWaterfallList(in *gin.Context) {
	var req trpc.GetWaterfallReq
	resp := &trpc.GetWaterfallResp{
		Code: 21000,
		Data: &trpc.WaterfallData{
			Items: make([]*trpc.WaterfallItem, 0),
		},
	}
	if err := in.BindJSON(&req); err != nil {
		in.JSON(http.StatusBadRequest, resp)
		return
	}

	query := ezdb.M{
		"Disabled": false,
	}

	sortFields := []string{}

	if req.GetName() != "" {
		query["Name"] = req.GetName()
	}

	if req.GetSecucode() != "" {
		query["Secucode"] = req.GetSecucode()
	}

	if req.GetDecrease() > 0 {
		sortFields = append(sortFields, "-Decrease")
		query["Decrease"] = ezdb.M{"$gte": req.GetDecrease()}
	}

	if req.GetDecrease() < 0 {
		sortFields = append(sortFields, "Decrease")
		query["Decrease"] = ezdb.M{"$lte": req.GetDecrease()}
	}

	if req.GetLimit() <= 0 {
		req.Limit = 10
	}

	if len(sortFields) <= 0 {
		sortFields = append(sortFields, "-Decrease")
	}

	fmt.Printf("==>>TODO 221: %+v|%+v\n", query, sortFields)
	results, err := orm.GPWaterfallLineMgr.Find(query, int(req.Limit), int(req.Offset), sortFields...)
	if err != nil {
		log.Errorf("query recommend failed: %q", err)
	}

	resp.Data.Total = int32(orm.GPWaterfallLineMgr.Count(query))

	wg := sync.WaitGroup{}
	wg.Add(len(results))

	for _, result := range results {
		result := result
		go func(wg *sync.WaitGroup) {
			resp.Data.Items = append(resp.Data.Items, &trpc.WaterfallItem{
				Name:           getName(result.Secucode),
				Secucode:       result.Secucode,
				Decrease:       result.Decrease,
				MaxPDay:        result.MaxPDay,
				MaxPrice:       result.MaxPrice,
				MinPrice:       result.MinPrice,
				Closing:        result.PresentPrice,
				PresentPrice:   request.GetSinaDayPrice(result.Secucode),
				State:          result.State,
				Traded:         getTraded(result.Secucode) * 100,
				InflowRatioStr: getSecucodeInflowRatioStr(result.Secucode, 5),
				CreateDate:     result.CreateDate,
				UpdateDate:     result.UpdateDate,
			})
			wg.Done()
		}(&wg)
	}

	wg.Wait()
	if req.GetDecrease() > 0 {
		sort.Slice(resp.Data.Items, func(i, j int) bool {
			return resp.Data.Items[i].Decrease > resp.Data.Items[j].Decrease
		})
	} else {
		sort.Slice(resp.Data.Items, func(i, j int) bool {
			return resp.Data.Items[i].Decrease < resp.Data.Items[j].Decrease
		})
	}

	resp.Code = 20000
	in.JSON(http.StatusOK, resp)
}

func UpdateWaterfallState(in *gin.Context) {
	var req trpc.GetGDSDLTReq
	resp := &trpc.UpdateWaterfallStateResp{
		Code: 21000,
	}
	if err := in.BindJSON(&req); err != nil {
		in.JSON(http.StatusBadRequest, resp)
		return
	}

	resp.Code = 20000
	in.JSON(http.StatusOK, resp)
}
