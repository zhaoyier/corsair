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

func AddGPRemark(in *gin.Context) {
	var req trpc.AddGPRemarkReq
	resp := &trpc.AddGPRemarkResp{
		Code: 21000,
	}
	if err := in.BindJSON(&req); err != nil {
		in.JSON(http.StatusBadRequest, resp)
		return
	}

	data := orm.GPRemarkMgr.NewGPRemark()
	data.Name = getName(req.GetSecucode())
	data.Secucode = req.GetSecucode()
	data.Remark = req.GetRemark()
	data.CreateDate = time.Now().Unix()

	if _, err := data.Save(); err != nil {
		in.JSON(http.StatusNotFound, resp)
		return
	}

	resp.Code = 20000
	in.JSON(http.StatusOK, resp)
}

func GetGPRemark(in *gin.Context) {
	var req trpc.GetGPRemarkReq
	resp := &trpc.GetGPRemarkResp{
		Code: 21000,
		Data: &trpc.GPRemarkData{
			Items: make([]*trpc.GPRemarkItem, 0),
		},
	}
	if err := in.BindJSON(&req); err != nil {
		in.JSON(http.StatusBadRequest, resp)
		return
	}
	limit := int(req.GetLimit())
	offset := int(req.GetOffset())
	query := ezdb.M{"Disabled": false}
	results, err := orm.GPRemarkMgr.Find(query, limit, offset, "-CreateDate")
	if err != nil {
		log.Errorf("query recommend failed: %q", err)
		in.JSON(http.StatusBadRequest, resp)
		return
	}

	for _, result := range results {
		resp.Data.Items = append(resp.Data.Items, &trpc.GPRemarkItem{
			Name:       result.Name,
			Secucode:   result.Secucode,
			Remark:     result.Remark,
			CreateDate: result.CreateDate,
		})
	}

	resp.Code = 20000
	in.JSON(http.StatusOK, resp)
}

func getGPRemarks(secucode string, limit, offset int) []*trpc.GPRemarkItem {
	items := make([]*trpc.GPRemarkItem, 0, 8)
	query := ezdb.M{"Disabled": false, "Secucode": secucode}
	results, err := orm.GPRemarkMgr.Find(query, limit, offset, "-CreateDate")
	if err != nil {
		return items
	}

	for _, result := range results {
		items = append(items, &trpc.GPRemarkItem{
			Name:       result.Name,
			Secucode:   result.Secucode,
			Remark:     result.Remark,
			CreateDate: result.CreateDate,
		})
	}
	return items
}
