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
	now := time.Now().Unix()
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

	result.Name = getName(req.GetSecucode())
	if result.ExpectPrice <= 0 {
		result.ExpectPrice = req.GetExpectPrice()
	}

	result.PresentPrice = 0
	result.CreateDate = now
	result.UpdateDate = now
	result.Remarks = append(result.Remarks, orm.GPRemark{
		Content:    req.GetRemark(),
		UpdateDate: now,
	})

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
	sortFields := []string{}
	if req.GetName() != "" {
		query["Name"] = req.GetName()
	}
	if req.GetSecucode() != "" {
		query["Secucode"] = req.GetSecucode()
	}
	if req.GetState() > 0 {
		sortFields = append(sortFields, "-State")
		query["State"] = ezdb.M{"$gte": req.GetState()}
	}

	if req.GetDisabled() == trpc.DisabledType_DisabledTypeValid {
		query["Disabled"] = false
	} else if req.GetDisabled() == trpc.DisabledType_DisabledTypeInvalid {
		query["Disabled"] = true
	}

	if req.GetExpectDateStart() > 0 && req.GetExpectDateEnd() > 0 {
		query["ExpectDate"] = ezdb.M{"$gte": req.GetExpectDateStart() / 1000, "$lte": req.GetExpectDateEnd() / 1000}
	}

	sortFields = append(sortFields, "-CreateDate")
	results, err := orm.GPFocusMgr.Find(query, int(req.GetLimit()), int(req.GetOffset()), sortFields...)
	if err != nil {
		log.Errorf("get prompt buy failed: %q", err)
		in.JSON(http.StatusForbidden, resp)
		return
	}

	total := orm.GPFocusMgr.Count(query)
	resp.Data.Total = int32(total)

	for _, result := range results {
		item := &trpc.GPFocusItem{
			Name:          result.Name,
			Secucode:      result.Secucode,
			PresentPrice:  getPresentPrice(result.Secucode),
			ExpectPrice:   result.ExpectPrice,
			Focused:       getFocused(result.Disabled),
			CreateDate:    time.Unix(result.CreateDate, 0).Format("2006-01-02"),
			UpdateDate:    time.Unix(result.UpdateDate, 0).Format("2006-01-02"),
			HoldFocus:     getHoldFocus(result.Secucode),
			TotalNumRatio: getTotalNumRatio(result.Secucode),
			Traded:        getTraded(result.Secucode),
			State:         getFocusState(result.State),
			ExpectDate:    result.ExpectDate,
		}

		item.DiffPrice = getDiffPrice(item.PresentPrice, item.ExpectPrice)
		for _, remark := range result.Remarks {
			item.Remarks = append(item.Remarks, &trpc.GPRemark{
				Content:    remark.Content,
				UpdateDate: remark.UpdateDate,
			})
			item.Remark = remark.Content
		}

		resp.Data.Items = append(resp.Data.Items, item)
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
		in.JSON(http.StatusBadRequest, resp)
		return
	}

	if req.GetSecucode() == "" {
		in.JSON(http.StatusBadRequest, resp)
		return
	}

	// query := ezdb.M{"Secucode": req.GetSecucode()}
	result, err := orm.GPFocusMgr.FindOneBySecucode(req.GetSecucode())
	if err != nil {
		in.JSON(http.StatusNotFound, resp)
		return
	}

	result.Disabled = true
	result.UpdateDate = time.Now().Unix()
	if _, err := result.Save(); err != nil {
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
		in.JSON(http.StatusBadRequest, resp)
		return
	}

	if req.GetSecucode() == "" {
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
	result.ExpectDate = req.GetExpectDate() / 1000
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

func getDiffPrice(present, expect float64) float64 {
	// log.Infof("==>>441: %+v|%+v", data.PresentPrice, data.ExpectPrice)
	if expect > 0 {
		return utils.TruncateFloat(present - expect)
	}
	return utils.TruncateFloat(present - expect)
}

func getFocusState(state int32) string {
	switch trpc.GPFocusState(state) {
	case trpc.GPFocusState_FocusStatePrepare:
		return "准备"
	case trpc.GPFocusState_FocusStateStart:
		return "开始"
	case trpc.GPFocusState_FocusStateProgress:
		return "进行中"
	default:
		return "待定"
	}
}
