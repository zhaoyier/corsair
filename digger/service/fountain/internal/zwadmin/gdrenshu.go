package zwadmin

import (
	"net/http"
	"sort"
	"strings"
	"sync"
	"time"

	orm "git.ezbuy.me/ezbuy/corsair/digger/service/internal/model"
	trpc "git.ezbuy.me/ezbuy/corsair/digger/service/internal/rpc"

	"git.ezbuy.me/ezbuy/corsair/digger/service/internal/utils"
	log "github.com/Sirupsen/logrus"
	ezdb "github.com/ezbuy/ezorm/db"
	"github.com/gin-gonic/gin"
)

func GDRenshuDetail(in *gin.Context) {
	var req trpc.GDRenshuDetailReq
	resp := &trpc.GDRenshuDetailResp{
		Code: 21000,
		Data: &trpc.GDRenshuDetailData{
			Items: make([]*trpc.GDRenshuItem, 0, 10),
		},
	}
	if err := in.BindJSON(&req); err != nil {
		in.JSON(http.StatusBadRequest, resp)
		return
	}

	query := ezdb.M{}
	sortFields := []string{"-EndDate"}
	if req.GetSecucode() != "" {
		query["Secucode"] = req.GetSecucode()
	}
	results, err := orm.GDRenshuMgr.Find(query, 20, 0, sortFields...)
	if err != nil {
		log.Errorf("save focus failed: %q", err)
		in.JSON(http.StatusForbidden, resp)
		return
	}
	wg := sync.WaitGroup{}
	wg.Add(len(results))

	for _, result := range results {
		result := result
		go func(wg *sync.WaitGroup) {
			resp.Data.Items = append(resp.Data.Items, &trpc.GDRenshuItem{
				Name:               getName(result.Secucode),
				Secucode:           result.Secucode,
				ReleaseDate:        time.Unix(result.EndDate, 0).Format("2006-01-02"),
				HolderTotalNum:     int32(result.HolderTotalNum),
				TotalNumRatio:      result.TotalNumRatio,
				HoldFocus:          result.HoldFocus,
				Price:              result.Price,
				HoldRatioTotal:     utils.TruncateFloat(result.HoldRatioTotal),
				FreeholdRatioTotal: utils.TruncateFloat(result.FreeholdRatioTotal),
			})
			wg.Done()
		}(&wg)
	}

	wg.Wait()
	sort.Slice(resp.Data.Items, func(i, j int) bool {
		return resp.Data.Items[i].ReleaseDate > resp.Data.Items[j].ReleaseDate
	})

	resp.Code = 20000
	in.JSON(http.StatusOK, resp)
}

func GDRenshuList(in *gin.Context) {
	var req trpc.GDRenshuListReq
	resp := &trpc.GDRenshuListResp{
		Code: 21000,
		Data: &trpc.GDRenshuData{
			Items: make([]*trpc.GDRenshuItem, 0, 10),
		},
	}
	if err := in.BindJSON(&req); err != nil {
		in.JSON(http.StatusBadRequest, resp)
		return
	}

	query := ezdb.M{}
	sortFields := []string{}
	if req.GetSecucode() != "" {
		query["Secucode"] = req.GetSecucode()
	}
	if req.GetName() != "" {
		// sortFields = append(sortFields, "-EndDate")
		query["Secucode"] = getSecucode(req.GetName())
	}

	if req.GetReleaseStart() > 0 && req.GetReleaseEnd() > 0 {
		query["EndDate"] = ezdb.M{"$gte": req.GetReleaseStart() / 1000, "$lte": req.GetReleaseEnd() / 1000}
	}

	if req.GetTotalRatio() > 0 {
		query["TotalNumRatio"] = ezdb.M{"$gte": req.GetTotalRatio()}
	}
	if req.GetTotalRatio() < 0 {
		query["TotalNumRatio"] = ezdb.M{"$lte": req.GetTotalRatio()}
	}
	switch trpc.GDRenshuSortType(req.GetSortTyp()) {
	case trpc.GDRenshuSortType_GDRenshuSortTypeDecrease:
		sortFields = append(sortFields, "TotalNumRatio")
	case trpc.GDRenshuSortType_GDRenshuSortTypeIncrease:
		sortFields = append(sortFields, "-TotalNumRatio")
	}
	sortFields = append(sortFields, "-EndDate")
	results, err := orm.GDRenshuMgr.Find(query, int(req.GetLimit()), int(req.GetOffset()), sortFields...)
	if err != nil {
		log.Errorf("save focus failed: %q", err)
		in.JSON(http.StatusForbidden, resp)
		return
	}

	total := orm.GDRenshuMgr.Count(query)
	resp.Data.Total = int32(total)

	wg := sync.WaitGroup{}
	wg.Add(len(results))

	for _, result := range results {
		result := result
		go func(wg *sync.WaitGroup) {
			resp.Data.Items = append(resp.Data.Items, &trpc.GDRenshuItem{
				Name:               getName(result.Secucode),
				Secucode:           result.Secucode,
				ReleaseDate:        time.Unix(result.EndDate, 0).Format("2006-01-02"),
				HolderTotalNum:     int32(result.HolderTotalNum),
				TotalNumRatio:      utils.TruncateFloat(result.TotalNumRatio),
				HoldFocus:          result.HoldFocus,
				Price:              result.Price,
				HoldRatioTotal:     utils.TruncateFloat(result.HoldRatioTotal),
				FreeholdRatioTotal: utils.TruncateFloat(result.FreeholdRatioTotal),
				PresentPrice:       getPresentPrice(result.Secucode),
				Focused:            getFocusBySecucode(result.Secucode),
			})
			wg.Done()
		}(&wg)
	}

	wg.Wait()

	if req.GetSortTyp() == int32(trpc.GDRenshuSortType_GDRenshuSortTypeDecrease) {
		sort.Slice(resp.Data.Items, func(i, j int) bool {
			return resp.Data.Items[i].TotalNumRatio < resp.Data.Items[j].TotalNumRatio
		})
	} else {
		sort.Slice(resp.Data.Items, func(i, j int) bool {
			return resp.Data.Items[i].TotalNumRatio > resp.Data.Items[j].TotalNumRatio
		})
	}

	resp.Code = 20000
	in.JSON(http.StatusOK, resp)
}

func GDAggregationReset(in *gin.Context) {
	var req trpc.GDAggregationResetReq
	resp := &trpc.GDAggregationResetResp{
		Code: 21000,
	}
	if err := in.BindJSON(&req); err != nil {
		in.JSON(http.StatusBadRequest, resp)
		return
	}

	query := ezdb.M{}
	if req.GetSecucode() != "" {
		query["Secucode"] = req.GetSecucode()
	}

	// 删除
	_, err := orm.GDAggregationMgr.RemoveAll(query)
	if err != nil {
		in.JSON(http.StatusNotExtended, resp)
		return
	}

	go resetGDAggregation(&req)

	resp.Code = 20000
	in.JSON(http.StatusOK, resp)
}

func GDAggregationList(in *gin.Context) {
	var req trpc.GDAggregationListReq
	resp := &trpc.GDAggregationListResp{
		Code: 21000,
		Data: &trpc.GDAggregationData{
			Items: make([]*trpc.GDAggregationItem, 0, 10),
		},
	}
	if err := in.BindJSON(&req); err != nil {
		in.JSON(http.StatusBadRequest, resp)
		return
	}

	query := ezdb.M{}
	sortFileds := []string{}
	if req.GetTotalNumRatio() > 0 {
		sortFileds = append(sortFileds, "-TotalRatioAccum")
		query["TotalRatioAccum"] = ezdb.M{"$gte": req.GetTotalNumRatio()}
	}
	if req.GetTotalNumRatio() < 0 {
		sortFileds = append(sortFileds, "TotalRatioAccum")
		query["TotalRatioAccum"] = ezdb.M{"$lte": req.GetTotalNumRatio()}
	}

	if req.GetReleaseStart() > 0 && req.GetReleaseEnd() > 0 {
		query["EndDate"] = ezdb.M{"$gte": req.GetReleaseStart(), "$lte": req.GetReleaseEnd()}
	}
	if req.GetPriceRatio() > 0 {
		sortFileds = append(sortFileds, "-PriceRatio")
		query["PriceRatio"] = ezdb.M{"$gte": req.GetPriceRatio()}
	}
	if req.GetPriceRatio() < 0 {
		sortFileds = append(sortFileds, "PriceRatio")
		query["PriceRatio"] = ezdb.M{"$lte": req.GetPriceRatio()}
	}

	if req.GetHoldRatio() > 0 {
		sortFileds = append(sortFileds, "-HoldRatioTotal")
		query["HoldRatioTotal"] = ezdb.M{"$gte": req.GetHoldRatio()}
	}
	if req.GetHoldRatio() < 0 {
		sortFileds = append(sortFileds, "HoldRatioTotal")
		query["HoldRatioTotal"] = ezdb.M{"$lte": req.GetHoldRatio()}
	}

	results, err := orm.GDAggregationMgr.Find(query, int(req.GetLimit()), int(req.GetOffset()), sortFileds...)
	if err != nil {
		in.JSON(http.StatusBadRequest, resp)
		return
	}

	total := orm.GDAggregationMgr.Count(query)
	resp.Data.Total = int32(total)

	for _, result := range results {
		resp.Data.Items = append(resp.Data.Items, &trpc.GDAggregationItem{
			Name:          result.Name,
			Secucode:      result.Secucode,
			EndDate:       result.EndDate,
			HolderRatio:   result.TotalRatioAccum,
			PriceRatio:    result.PriceRatio,
			PriceMax:      result.PriceMax,
			PriceMin:      result.PriceMin,
			HoldFocus:     result.HoldFocus,
			HoldRatio:     result.HoldRatioTotal,
			HolderNum:     result.HolderTotalNum,
			FreeholdRatio: result.FreeholdRatioTotal,
			UpdateDate:    result.UpdateDate,
			Focused:       getFocusBySecucode(result.Secucode),
		})
	}

	resp.Code = 20000
	in.JSON(http.StatusOK, resp)
}

func getSecucode(name string) string {
	result, err := orm.CNSecucodeMgr.FindOneByName(name)
	if err != nil {
		return ""
	}
	return result.Secucode
}

func getPresentPrice(secucode string) float64 {
	secucode = strings.Split(secucode, ".")[1]
	result, err := orm.GPDailyMgr.FindOne(ezdb.M{"Secucode": secucode}, "-CreateDate")
	if err != nil {
		return 0
	}
	return utils.TruncateFloat(result.Closing)
}

func resetGDAggregation(in *trpc.GDAggregationResetReq) error {
	// 重新生成
	sess, col := orm.CNSecucodeMgr.GetCol()
	defer sess.Close()

	query := ezdb.M{"Disabled": false}
	if in.GetSecucode() != "" {
		query["Secucode"] = in.GetSecucode()
	}

	//查询最近6个月的数据
	var secucode *orm.CNSecucode
	iter := col.Find(query).Batch(100).Prefetch(0.25).Iter()

	for iter.Next(&secucode) {
		resetSecucodeGD(secucode.Secucode, in)
	}

	return nil
}

func resetSecucodeGD(secucode string, in *trpc.GDAggregationResetReq) error {
	query := ezdb.M{"Secucode": secucode}
	if in.GetPeriodStart() > 0 && in.GetPeriodEnd() > 0 {
		query["EndDate"] = ezdb.M{"$gte": in.GetPeriodStart() / 1000, "$lte": in.GetPeriodEnd() / 1000}
	} else {
		return nil
	}

	results, err := orm.GDRenshuMgr.Find(query, 0, 0, "-EndDate")
	if err != nil {
		log.Errorf("query gdrenshu failed: %s|%q", secucode, err)
		return err
	}
	if len(results) <= 0 {
		return nil
	}

	data, _ := orm.GDAggregationMgr.FindOneBySecucode(secucode)
	if data == nil {
		data = orm.GDAggregationMgr.NewGDAggregation()
	}

	for idx, result := range results {
		if idx == 0 {
			data.Secucode = secucode
			data.Name = getName(secucode)
			data.EndDate = result.EndDate
			data.HolderTotalNum = int32(result.HolderTotalNum)
			data.HoldFocus = result.HoldFocus
			data.HoldRatioTotal = int32(result.HoldRatioTotal)
			data.FreeholdRatioTotal = int32(result.FreeholdRatioTotal)
			data.UpdateDate = time.Now().Unix()
			data.PriceMin = result.Price
		} else {
			if data.PriceMax < result.Price {
				data.PriceMax = result.Price
			}
			if data.PriceMin > result.Price {
				data.PriceMin = result.Price
			}
		}

		data.TotalRatioAccum += int32(result.TotalNumRatio)
	}

	if data.Name == "" {
		return nil
	}

	if data.PriceMin <= 0 || data.PriceMax <= 0 {
		return nil
	}

	data.PriceMin = getPresentPrice(secucode)
	data.PriceRatio = int32(utils.GetPercent((data.PriceMax - data.PriceMin), data.PriceMax))
	if _, err := data.Save(); err != nil {
		log.Errorf("update aggregation failed: %s|%q", secucode, err)
		return err
	}

	return nil
}
