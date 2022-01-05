package zwadmin

import (
	"net/http"
	"strings"
	"time"

	orm "git.ezbuy.me/ezbuy/corsair/digger/service/internal/model"
	trpc "git.ezbuy.me/ezbuy/corsair/digger/service/internal/rpc"
	"git.ezbuy.me/ezbuy/corsair/digger/service/internal/utils"

	log "github.com/Sirupsen/logrus"
	ezdb "github.com/ezbuy/ezorm/db"
	"github.com/gin-gonic/gin"
)

func ConfirmFocus(in *gin.Context) {
	var req trpc.FocusConfirmReq
	resp := &trpc.FocusConfirmResp{
		Code: 21000,
	}
	if err := in.BindJSON(&req); err != nil {
		in.JSON(http.StatusBadRequest, resp)
		return
	}

	query := ezdb.M{
		"Secucode": req.GetSecucode(),
	}

	result, _ := orm.GPFocusMgr.FindOne(query)
	if result == nil {
		result = orm.GPFocusMgr.NewGPFocus()
		result.Secucode = req.GetSecucode()
	} else {
		result.Disabled = !result.Disabled
	}

	result.Name = req.GetName()
	if result.FocusPrice <= 0 {
		result.FocusPrice = req.GetPresentPrice()
	}

	result.PresentPrice = 0
	result.CreateDate = time.Now().Unix()
	result.UpdateDate = time.Now().Unix()

	if _, err := result.Save(); err != nil {
		log.Errorf("save focus failed: %q", err)
		in.JSON(http.StatusForbidden, resp)
		return
	}

	resp.Code = 20000
	in.JSON(http.StatusOK, resp)
}

func GetFocusList(in *gin.Context) {
	var req trpc.FocusListReq
	resp := &trpc.FocusListResp{
		Code: 21000,
		Data: &trpc.GPFocusData{
			Items: make([]*trpc.GPFocusItem, 0),
		},
	}
	if err := in.BindJSON(&req); err != nil {
		in.JSON(http.StatusBadRequest, resp)
		return
	}

	query := ezdb.M{}
	sortField := "-CreateDate"
	if req.GetName() != "" {
		query["Name"] = req.GetName()
	}
	if req.GetSecucode() != "" {
		query["Secucode"] = req.GetSecucode()
	}
	if req.GetDisabled() == trpc.DisabledType_DisabledTypeValid {
		query["Disabled"] = false
	} else if req.GetDisabled() == trpc.DisabledType_DisabledTypeInvalid {
		query["Disabled"] = true
	}

	results, err := orm.GPFocusMgr.Find(query, int(req.GetLimit()), int(req.GetOffset()), sortField)
	if err != nil {
		log.Errorf("get prompt buy failed: %q", err)
		in.JSON(http.StatusForbidden, resp)
		return
	}

	total := orm.GPFocusMgr.Count(query)
	resp.Data.Total = int32(total)

	for _, result := range results {
		resp.Data.Items = append(resp.Data.Items, &trpc.GPFocusItem{
			Name:          result.Name,
			Secucode:      result.Secucode,
			FocusPrice:    result.FocusPrice,
			PresentPrice:  getPresentPrice(result.Secucode),
			ExpectPrice:   result.ExpectPrice,
			DiffPrice:     getDiffPrice(result),
			Focused:       getFocused(result.Disabled),
			CreateDate:    time.Unix(result.CreateDate, 0).Format("2006-01-02"),
			UpdateDate:    time.Unix(result.UpdateDate, 0).Format("2006-01-02"),
			HoldFocus:     getHoldFocus(result.Secucode),
			TotalNumRatio: getTotalNumRatio(result.Secucode),
			Traded:        getTraded(result.Secucode),
		})
	}

	resp.Code = 20000
	in.JSON(http.StatusOK, resp)
}

func CancelFocus(in *gin.Context) {
	var req trpc.CancelFocusReq
	resp := &trpc.CancelFocusResp{
		Code: 21000,
	}
	if err := in.BindJSON(&req); err != nil {
		log.Infof("==>>TODO 121: %+v", nil)
		in.JSON(http.StatusBadRequest, resp)
		return
	}

	if req.GetSecucode() == "" {
		log.Infof("==>>TODO 122: %+v", nil)
		in.JSON(http.StatusBadRequest, resp)
		return
	}

	// query := ezdb.M{"Secucode": req.GetSecucode()}
	result, err := orm.GPFocusMgr.FindOneBySecucodeDisabled(req.GetSecucode(), false)
	if err != nil {
		log.Infof("==>>TODO 123: %+v", nil)
		in.JSON(http.StatusNotFound, resp)
		return
	}

	result.Disabled = true
	result.UpdateDate = time.Now().Unix()
	if _, err := result.Save(); err != nil {
		log.Infof("==>>TODO 124: %+v", nil)
		in.JSON(http.StatusNotFound, resp)
		return
	}

	resp.Code = 20000
	in.JSON(http.StatusOK, resp)
}

func UpdateFocus(in *gin.Context) {
	var req trpc.UpdateFocusReq
	resp := &trpc.UpdateFocusResp{
		Code: 21000,
	}
	if err := in.BindJSON(&req); err != nil {
		log.Infof("==>>TODO 121: %+v", nil)
		in.JSON(http.StatusBadRequest, resp)
		return
	}

	if req.GetSecucode() == "" {
		log.Infof("==>>TODO 122: %+v", nil)
		in.JSON(http.StatusBadRequest, resp)
		return
	}

	query := ezdb.M{"Secucode": req.GetSecucode()}
	result, err := orm.GPFocusMgr.FindOne(query)
	if err != nil {
		in.JSON(http.StatusNotFound, resp)
		return
	}
	result.ExpectPrice = req.GetExpectPrice()
	result.UpdateDate = time.Now().Unix()
	if _, err := result.Save(); err != nil {
		in.JSON(http.StatusConflict, resp)
		return
	}

	resp.Code = 20000
	in.JSON(http.StatusOK, resp)
}

func getFocused(disabled bool) string {
	if disabled {
		return "关注"
	}
	return "取消关注"
}

func getHoldFocus(secucode string) string {
	result, err := orm.GDRenshuMgr.FindOne(ezdb.M{"Secucode": secucode}, "-EndDate")
	if err != nil {
		return ""
	}
	return result.HoldFocus
}

func getTotalNumRatio(secucode string) float64 {
	var ratio float64
	start := time.Now().AddDate(0, -1, 0).Unix()
	query := ezdb.M{
		"Secucode": secucode,
		"EndDate":  ezdb.M{"$gte": start},
	}

	results, err := orm.GDRenshuMgr.FindAll(query, "-EndDate")
	if err != nil {
		return ratio
	}
	for _, result := range results {
		ratio += result.TotalNumRatio
	}

	return utils.TruncateFloat(ratio)
}

func getTraded(secucode string) int64 {
	secucode = strings.Split(secucode, ".")[1]
	result, err := orm.GPDailyMgr.FindOne(ezdb.M{"Secucode": secucode}, "-CreateDate")
	if err != nil {
		return 0
	}
	return result.Traded / 100000000
}

func getDiffPrice(data *orm.GPFocus) float64 {
	if data.ExpectPrice > 0 {
		return utils.TruncateFloat(data.PresentPrice - data.ExpectPrice)
	}
	return utils.TruncateFloat(data.PresentPrice - data.FocusPrice)
}
