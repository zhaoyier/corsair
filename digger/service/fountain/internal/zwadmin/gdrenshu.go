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
				HoldRatioTotal:     result.HoldRatioTotal,
				FreeholdRatioTotal: result.FreeholdRatioTotal,
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
		query["Secucode"] = getSecucode(req.GetName())
	}

	if req.GetReleaseStart() > 0 && req.GetReleaseEnd() > 0 {
		query["EndDate"] = ezdb.M{"$gte": req.GetReleaseStart() / 1000, "$lte": req.GetReleaseEnd() / 1000}
	}

	if req.GetTotalRatio() > 0 {
		query["TotalNumRatio"] = ezdb.M{"$gte": req.GetTotalRatio()}
	} else {
		query["TotalNumRatio"] = ezdb.M{"$lte": req.GetTotalRatio()}
	}
	log.Infof("==>>TODO 123: %+v|%+v", query, sortFields)
	switch trpc.GDRenshuSortType(req.GetSortTyp()) {
	case trpc.GDRenshuSortType_GDRenshuSortTypeDecrease:
		sortFields = append(sortFields, "TotalNumRatio")
	case trpc.GDRenshuSortType_GDRenshuSortTypeIncrease:
		sortFields = append(sortFields, "-TotalNumRatio")
	}
	// sortFields = append(sortFields, "-EndDate")
	log.Infof("==>>TODO 123: %+v|%+v", query, sortFields)
	results, err := orm.GDRenshuMgr.Find(query, int(req.GetLimit()), int(req.GetOffset()), sortFields...)
	// results, err := orm.GDRenshuMgr.Find(query, int(req.GetLimit()), int(req.GetOffset()), "TotalNumRatio")
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
				HoldRatioTotal:     result.HoldRatioTotal,
				FreeholdRatioTotal: result.FreeholdRatioTotal,
				PresentPrice:       getPresentPrice(result.Secucode),
				Focused:            getFocus(result.Secucode),
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

func getName(secucode string) string {
	result, err := orm.CNSecucodeMgr.FindOneBySecucode(secucode)
	if err != nil {
		return ""
	}
	return result.Name
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

func getFocus(secucode string) string {
	result, err := orm.GPFocusMgr.FindOneBySecucodeDisabled(secucode, false)
	if err != nil {
		return "关注"
	}
	if result == nil {
		return "关注"
	}
	return "取消关注"
}
