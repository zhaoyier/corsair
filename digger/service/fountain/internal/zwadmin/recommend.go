package zwadmin

import (
	"net/http"
	"sort"
	"strconv"
	"strings"
	"sync"
	"time"

	orm "git.ezbuy.me/ezbuy/corsair/digger/service/internal/model"
	trpc "git.ezbuy.me/ezbuy/corsair/digger/service/internal/rpc"
	"git.ezbuy.me/ezbuy/corsair/digger/service/internal/utils"

	"git.ezbuy.me/ezbuy/corsair/digger/service/internal/dawdle"
	log "github.com/Sirupsen/logrus"
	ezdb "github.com/ezbuy/ezorm/db"
	"github.com/gin-gonic/gin"
)

func GPRecommendList(in *gin.Context) {
	limit, _ := strconv.Atoi(in.Query("limit"))
	offset, _ := strconv.Atoi(in.Query("offset"))
	resp := &trpc.GPRecommendListResp{
		Rows: make([]*trpc.GPRecommend, 0),
	}

	query := ezdb.M{
		"Disabled": false,
	}

	if limit <= 0 {
		limit = 10
	}

	results, err := orm.GPRecommendMgr.Find(query, limit, offset, "-PDecrease", "-State", "GDDecrease")
	if err != nil {
		log.Errorf("query recommend failed: %q", err)
	}

	resp.Total = int32(orm.GPRecommendMgr.Count(query))

	for idx, result := range results {
		resp.Rows = append(resp.Rows, &trpc.GPRecommend{
			Id:           int32(idx + 1),
			Secucode:     result.Secucode,
			Name:         result.Name, //getName(result.Secucode),
			RMIndex:      result.RMIndex,
			PDecrease:    result.PDecrease,
			MaxPrice:     result.MaxPrice,
			MaxPDay:      utils.TS2Date(result.MaxPDay),
			RMPrice:      getRMPrice(result.RMPrice),
			GDDecrease:   result.GDDecrease,
			State:        getState(result.State),
			UpdateDate:   utils.TS2Date(result.UpdateDate),
			PresentPrice: result.PresentPrice,
		})
	}

	resp.Rows = sortRecommend(resp.Rows)

	in.JSON(http.StatusOK, resp)
}

func GetRecommend(in *gin.Context) {
	var req trpc.GetRecommendReq
	resp := &trpc.GetRecommendResp{
		Code: 21000,
		Data: &trpc.RecommendData{
			Items: make([]*trpc.RecommendItem, 0),
		},
	}
	if err := in.BindJSON(&req); err != nil {
		in.JSON(http.StatusBadRequest, resp)
		return
	}

	query := ezdb.M{
		"Disabled": false,
	}

	if req.GetPDecrease() > 0 {
		query["PDecrease"] = ezdb.M{"$gte": req.GetPDecrease()}
	}
	if req.GetState() > 0 {
		query["State"] = ezdb.M{"$gte": req.GetState()}
	}
	if req.GetSecucode() != "" {
		query["Secucode"] = req.GetSecucode()
	}
	if req.GetRmIndex() > 0 {
		query["RMIndex"] = ezdb.M{"$gte": req.GetRmIndex()}
	}

	if req.GetLimit() <= 0 {
		req.Limit = 20
	}

	results, err := orm.GPRecommendMgr.Find(query, int(req.Limit), int(req.GetOffset()), "-PDecrease", "-State", "GDDecrease")
	if err != nil {
		log.Errorf("query recommend failed: %q", err)
	}

	resp.Data.Total = int32(orm.GPRecommendMgr.Count(query))

	wg := sync.WaitGroup{}
	wg.Add(len(results))

	for idx, result := range results {
		idx, result := idx, result
		go func(wg *sync.WaitGroup) {
			item := &trpc.RecommendItem{
				Id:            int32(idx + 1),
				Secucode:      result.Secucode,
				Name:          result.Name,
				RMIndex:       result.RMIndex,
				PDecrease:     result.PDecrease,
				MaxPrice:      result.MaxPrice,
				MaxPDay:       utils.TS2Date(result.MaxPDay),
				RMPrice:       getRMPrice(result.RMPrice),
				GDDecrease:    result.GDDecrease,
				State:         getState(result.State),
				UpdateNum:     result.UpdateNum,
				UpdateDate:    utils.TS2Date(result.UpdateDate),
				PresentPrice:  result.PresentPrice,
				ReferDecrease: getReferDecrease(result.Secucode),
				Focused:       getFocusBySecucode(result.Secucode),
			}

			item.Traded = getTraded(item.Secucode) * 100
			item.FundFive, item.FundTwenty = getFundFlow(item.Secucode)
			item.FundPercent = int32(utils.GetPercent(float64(item.FundTwenty), float64(item.Traded*100)))
			resp.Data.Items = append(resp.Data.Items, item)

			wg.Done()
		}(&wg)
	}

	wg.Wait()

	resp.Code = 20000
	resp.Data.Items = sortRecommend2(resp.Data.Items)

	in.JSON(http.StatusOK, resp)
}

func UpdateRecommend(in *gin.Context) {
	var req trpc.UpdateRecommendReq
	resp := &trpc.UpdateRecommendResp{
		Code: 21000,
	}
	if err := in.BindJSON(&req); err != nil {
		in.JSON(http.StatusBadRequest, resp)
		return
	}
	result := orm.GPManualDecreaseMgr.MustFindOneBySecucodeDisabled(req.GetSecucode(), false)
	result.Name = req.GetName()
	result.DecreaseTag = req.GetPriceDecrease()
	result.UpdateDate = time.Now().Unix()
	if _, err := result.Save(); err != nil {
		in.JSON(http.StatusNotFound, resp)
		return
	}

	dawdle.GenRecommendTmp(req.GetSecucode())

	resp.Code = 20000
	in.JSON(http.StatusOK, resp)
}

func getName(secucode string) string {
	result, err := orm.CNSecucodeMgr.FindOneBySecucode(secucode)
	if err != nil {
		return ""
	}
	return result.Name
}

func getGDDecrease(secucode string) int32 {
	query := ezdb.M{
		"Secucode": secucode,
	}
	result, err := orm.GDLongLineMgr.FindOne(query, "-CreateDate")
	if err != nil {
		return 0
	}
	return result.GDReduceRatio
}

func getState(state int32) string {
	switch state {
	case 1:
		return "准备"
	case 2:
		return "开始"
	case 3:
		return "第一阶段"
	case 4:
		return "第二阶段"
	case 5:
		return "第三阶段"
	case 6:
		return "结束"
	case 7:
		return "放弃"
	default:
		return "准备"
	}
}

func sortRecommend(rows []*trpc.GPRecommend) []*trpc.GPRecommend {
	list := make([]*trpc.GPRecommend, 0, len(rows))
	results := make([]*trpc.RecommendSort, 0, len(rows))

	for idx, row := range rows {
		results = append(results, &trpc.RecommendSort{
			Idx:      int32(idx),
			Secucode: row.Secucode,
			Ratio:    getRMPriceRation(row.RMPrice, row.PresentPrice),
		})
	}
	sort.Slice(results, func(i, j int) bool {
		return results[i].Ratio < results[j].Ratio
	})

	for _, result := range results {
		list = append(list, rows[result.Idx])
	}
	return list
}

func sortRecommend2(rows []*trpc.RecommendItem) []*trpc.RecommendItem {
	list := make([]*trpc.RecommendItem, 0, len(rows))
	results := make([]*trpc.RecommendSort, 0, len(rows))

	for idx, row := range rows {
		results = append(results, &trpc.RecommendSort{
			Idx:      int32(idx),
			Secucode: row.Secucode,
			Ratio:    getRMPriceRation(row.RMPrice, row.PresentPrice),
		})
	}
	sort.Slice(results, func(i, j int) bool {
		return results[i].Ratio < results[j].Ratio
	})

	for _, result := range results {
		list = append(list, rows[result.Idx])
	}
	return list
}

func getRMPriceRation(price string, present float64) int32 {
	results := strings.Split(price, "(")
	rmPrice := utils.String2Float64(results[0])
	if rmPrice == 0 {
		return 100
	}
	return int32(((present - rmPrice) / rmPrice) * 100)
}

func getFundFlow(secucode string) (int32, int32) {
	result, err := orm.GPFundFlowMgr.FindOneBySecucode(secucode)
	if err != nil {
		return 0, 0
	}
	// 百万级别
	return result.Five / 1000000, result.Twenty / 1000000
}
