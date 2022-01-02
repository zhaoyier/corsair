package zwadmin

import (
	"net/http"
	"time"

	orm "git.ezbuy.me/ezbuy/corsair/digger/service/internal/model"
	trpc "git.ezbuy.me/ezbuy/corsair/digger/service/internal/rpc"

	log "github.com/Sirupsen/logrus"
	ezdb "github.com/ezbuy/ezorm/db"
	"github.com/gin-gonic/gin"
)

func FocusConfirm(in *gin.Context) {
	var req trpc.FocusConfirmReq
	resp := &trpc.FocusConfirmResp{
		Code: 21000,
	}
	if err := in.BindJSON(&req); err != nil {
		in.JSON(http.StatusBadRequest, resp)
		return
	}

	data := orm.GPFocusMgr.MustFindOneBySecucodeDisabled(req.GetSecucode(), false)
	if data.CreateDate <= 0 {
		data.CreateDate = time.Now().Unix()
	}

	data.Disabled = false
	data.Name = req.GetName()
	if data.FocusPrice <= 0 {
		data.FocusPrice = req.GetPresentPrice()
	}

	data.PresentPrice = 0
	data.UpdateDate = time.Now().Unix()

	if _, err := data.Save(); err != nil {
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
	var sortField string

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
			Name:         result.Name,
			Secucode:     result.Secucode,
			FocusPrice:   result.FocusPrice,
			PresentPrice: result.PresentPrice,
			CreateDate:   time.Unix(result.CreateDate, 0).Format("2006-01-02"),
			UpdateDate:   time.Unix(result.UpdateDate, 0).Format("2006-01-02"),
		})
	}

	resp.Code = 20000
	in.JSON(http.StatusOK, resp)
}

// func getFocusPrice(secucode string) float64 {
// 	codes := strings.Split(secucode, ".")
// 	secucode = codes[len(codes)-1]
// 	orm.
// }

// func getFocusName(secucode string) string {
// 	result, err := orm.CNSecucodeMgr.FindOneBySecucode(secucode)
// 	if err != nil {
// 		log.Errorf("get secucode failed: %s|%q", secucode, err)
// 		return ""
// 	}

// 	return result.Name
// }
