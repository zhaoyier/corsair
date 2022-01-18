package zwadmin

import (
	"net/http"
	"sort"
	"sync"

	orm "git.ezbuy.me/ezbuy/corsair/digger/service/internal/model"
	trpc "git.ezbuy.me/ezbuy/corsair/digger/service/internal/rpc"

	log "github.com/Sirupsen/logrus"
	ezdb "github.com/ezbuy/ezorm/db"
	"github.com/gin-gonic/gin"
)

func GetFundFlowList(in *gin.Context) {
	var req trpc.GetFundFlowListReq
	resp := &trpc.GetFundFlowListResp{
		Code: 21000,
		Data: &trpc.GPFundFlowData{
			Items: make([]*trpc.GPFundFlowItem, 0),
		},
	}
	if err := in.BindJSON(&req); err != nil {
		in.JSON(http.StatusBadRequest, resp)
		return
	}

	query := ezdb.M{}
	sortField := ""
	limit := int(req.GetLimit())
	offset := int(req.GetOffset())
	if req.GetSecucode() != "" {
		query["Secucode"] = req.GetSecucode()
	} else {
		query["FundDate"] = getCreateDate()
	}

	if req.GetName() != "" {
		query["Secucode"] = getSecucode(req.GetName())
	}

	if req.GetFundStart() > 0 && req.GetFundEnd() > 0 {
		sortField = "-FundDate"
		query["FundDate"] = ezdb.M{"$gte": req.GetFundStart(), "$lte": req.GetFundEnd()}
	}

	if req.GetInflowRatio() > 0 {
		sortField = "-InflowRatio"
		query["InflowRatio"] = ezdb.M{"$gte": req.GetInflowRatio()}
	}
	if req.GetInflowRatio() < 0 {
		sortField = "InflowRatio"
		query["InflowRatio"] = ezdb.M{"$gte": req.GetInflowRatio()}
	}

	if req.GetInflow() > 0 {
		query["Inflow"] = ezdb.M{"$gte": req.GetInflow() * 1000000}
	}

	log.Infof("==>>TODO 202: %+v", query)
	results, err := orm.GPFundFlowMgr.Find(query, limit, offset, sortField)
	if err != nil {
		log.Errorf("get fund flow failed: %q", err)
		in.JSON(http.StatusForbidden, resp)
		return
	}
	total := orm.GPFundFlowMgr.Count(query)
	resp.Data.Total = int32(total)

	for _, result := range results {
		resp.Data.Items = append(resp.Data.Items, &trpc.GPFundFlowItem{
			Name:          getName(result.Secucode),
			Secucode:      result.Secucode,
			FundDate:      result.FundDate,
			Inflow:        result.Inflow / 1000000,
			InflowRatio:   result.InflowRatio,
			PresentPrice:  result.PresentPrice,
			IncreaseRatio: result.IncreaseRatio,
			CreateDate:    result.CreateDate,
			Focused:       getFocusByName(result.Name),
			Traded:        getTraded(result.Secucode) * 100,
		})
	}

	resp.Code = 20000
	in.JSON(http.StatusOK, resp)
}

func GetFundDetailList(in *gin.Context) {
	var req trpc.GetFundDetailListReq
	resp := &trpc.GetFundDetailListResp{
		Code: 21000,
		Data: &trpc.FundDetailData{
			Items: make([]*trpc.FundDetailItem, 0),
		},
	}
	if err := in.BindJSON(&req); err != nil {
		in.JSON(http.StatusBadRequest, resp)
		return
	}

	wg := sync.WaitGroup{}
	wg.Add(len(req.GetSecucodes()))

	for _, secucode := range req.GetSecucodes() {
		secucode := secucode
		go func(wg *sync.WaitGroup) {
			name := getName(secucode)
			results := getSecucodeFund(secucode)
			item := &trpc.FundDetailItem{
				Name:  name,
				Type:  "line",
				Stack: "Total",
			}
			sort.Slice(results, func(i, j int) bool {
				return results[i].FundDate < results[j].FundDate
			})

			for _, result := range results {
				item.Data = append(item.Data, result.InflowRatio)
			}
			resp.Data.Items = append(resp.Data.Items, item)
			resp.Data.LegendData = append(resp.Data.LegendData, name)
			wg.Done()
		}(&wg)
	}

	wg.Wait()

	resp.Code = 20000
	in.JSON(http.StatusOK, resp)
}

func getSecucodeFund(secucode string) []*trpc.GPFundFlowItem {
	query := ezdb.M{"Secucode": secucode}
	items := make([]*trpc.GPFundFlowItem, 0, 8)

	results, err := orm.GPFundFlowMgr.Find(query, 20, 0, "-FundDate")
	if err != nil {
		log.Errorf("get fund flow failed: %q", err)
		return nil
	}

	for _, result := range results {
		items = append(items, &trpc.GPFundFlowItem{
			Name:          getName(result.Secucode),
			Secucode:      result.Secucode,
			FundDate:      result.FundDate,
			Inflow:        result.Inflow,
			InflowRatio:   result.InflowRatio,
			PresentPrice:  result.PresentPrice,
			IncreaseRatio: result.IncreaseRatio,
			CreateDate:    result.CreateDate,
		})
	}
	return items
}
