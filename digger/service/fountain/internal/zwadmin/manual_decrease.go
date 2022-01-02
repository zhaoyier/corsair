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

func ManualDecreaseList(in *gin.Context) {
	var req trpc.ManualDecreaseListReq
	resp := &trpc.ManualDecreaseListResp{
		Code: 21000,
		Data: &trpc.ManualDecreaseData{
			Items: make([]*trpc.ManualDecreaseItem, 0),
		},
	}
	if err := in.BindJSON(&req); err != nil {
		in.JSON(http.StatusBadRequest, resp)
		return
	}

	var sortField string
	query := ezdb.M{}

	if req.GetName() != "" {
		query["Name"] = req.GetName()
	} else if req.GetSecucode() != "" {
		query["Secucode"] = utils.GetSecucode(req.GetSecucode())
	}

	results, err := orm.GPManualDecreaseMgr.Find(query, int(req.GetLimit()), int(req.GetOffset()), sortField)
	if err != nil {
		log.Errorf("get manual decrease failed: %q", err)
		in.JSON(http.StatusForbidden, resp)
		return
	}
	for _, result := range results {
		resp.Data.Items = append(resp.Data.Items, &trpc.ManualDecreaseItem{
			Name:        result.Name,
			Secucode:    result.Secucode,
			DecreaseTag: result.DecreaseTag,
			CreateDate:  time.Unix(result.CreateDate, 0).Format("2006-01-02"),
			UpdateDate:  time.Unix(result.UpdateDate, 0).Format("2006-01-02"),
		})
	}

	resp.Code = 20000
	in.JSON(http.StatusOK, resp)
}
