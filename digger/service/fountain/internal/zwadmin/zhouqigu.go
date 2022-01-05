package zwadmin

import (
	"net/http"
	"sync"
	"time"

	orm "git.ezbuy.me/ezbuy/corsair/digger/service/internal/model"
	trpc "git.ezbuy.me/ezbuy/corsair/digger/service/internal/rpc"
	"git.ezbuy.me/ezbuy/lib/log"

	ezdb "github.com/ezbuy/ezorm/db"
	"github.com/gin-gonic/gin"
)

func AddGPZhouQi(in *gin.Context) {
	var req trpc.AddGPZhouQiReq
	resp := &trpc.AddGPZhouQiResp{
		Code: 21000,
	}
	if err := in.BindJSON(&req); err != nil {
		in.JSON(http.StatusBadRequest, resp)
		return
	}

	result := orm.GPZhouQiMgr.MustFindOneBySecucode(req.GetSecucode())
	result.Name = getName(req.GetSecucode())
	result.Secucode = req.GetSecucode()
	result.PresentPrice = getPresentPrice(req.GetSecucode())
	result.ExpectMin = req.GetExpectMin()
	result.ExpectMax = req.GetExpectMax()
	result.ExpectStart = req.GetExpectStart()
	result.CreateDate = time.Now().Unix()
	if result.CreateDate <= 0 {
		result.UpdateDate = time.Now().Unix()
	}

	if _, err := result.Save(); err != nil {
		log.Errorf("save zhou qi failed: %s|%q", req.GetSecucode(), err)
		in.JSON(http.StatusNotModified, resp)
		return
	}

	resp.Code = 20000
	in.JSON(http.StatusOK, resp)
}

func UpdateGPZhouQi(in *gin.Context) {
	var req trpc.UpdateGPZhouQiReq
	resp := &trpc.UpdateGPZhouQiResp{
		Code: 21000,
	}
	if err := in.BindJSON(&req); err != nil {
		in.JSON(http.StatusBadRequest, resp)
		return
	}
	log.Infof("==>>TODO 210: %+v", req)
	result := orm.GPZhouQiMgr.MustFindOneBySecucode(req.GetSecucode())
	result.Name = getName(req.GetSecucode())
	result.Secucode = req.GetSecucode()
	result.PresentPrice = getPresentPrice(req.GetSecucode())
	result.ExpectMin = req.GetExpectMin()
	result.ExpectMax = req.GetExpectMax()
	result.ExpectStart = req.GetExpectStart() / 1000
	result.ExpectEnd = req.GetExpectEnd() / 1000
	result.Disabled = req.GetDisabled()
	result.UpdateDate = time.Now().Unix()
	if result.CreateDate <= 0 {
		result.CreateDate = time.Now().Unix()
	}

	if _, err := result.Save(); err != nil {
		log.Errorf("save zhou qi failed: %s|%q", req.GetSecucode(), err)
		in.JSON(http.StatusNotModified, resp)
		return
	}

	resp.Code = 20000
	in.JSON(http.StatusOK, resp)
}

func GPZhouQiList(in *gin.Context) {
	var req trpc.GPZhouQiListReq
	resp := &trpc.GPZhouQiListResp{
		Code: 21000,
		Data: &trpc.GPZhouQiData{
			Items: make([]*trpc.GPZhouQiItem, 0),
		},
	}
	if err := in.BindJSON(&req); err != nil {
		log.Errorf("bind param failed: %q", err)
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
	if req.GetExpectStart() > 0 && req.GetExpectEnd() > 0 {
		query["ExpectStart"] = ezdb.M{"$gte": req.GetExpectStart() / 1000, "$lte": req.GetExpectEnd() / 1000}
	}

	results, err := orm.GPZhouQiMgr.Find(query, int(req.GetLimit()), int(req.GetOffset()), sortFields...)
	if err != nil {
		log.Errorf("get prompt buy failed: %q", err)
		in.JSON(http.StatusNotFound, resp)
		return
	}

	total := orm.GPZhouQiMgr.Count(query)
	resp.Data.Total = int32(total)

	wg := sync.WaitGroup{}
	wg.Add(len(results))

	for _, result := range results {
		result := result
		go func(wg *sync.WaitGroup) {
			item := &trpc.GPZhouQiItem{
				Name:         result.Name,
				Secucode:     result.Secucode,
				PresentPrice: getPresentPrice(result.Secucode),
				ExpectMin:    result.ExpectMin,
				ExpectMax:    result.ExpectMax,
				ExpectStart:  result.ExpectStart,
				ExpectEnd:    result.ExpectEnd,
				UpdateDate:   time.Unix(result.UpdateDate, 0).Format("2006-01-02"),
				Remarks:      make([]*trpc.GPZhouQiRemark, 0),
				State:        "result.State",
			}

			for _, val := range result.Remarks {
				item.Remarks = append(item.Remarks, &trpc.GPZhouQiRemark{
					Remark:     val.Remark,
					CreateDate: time.Unix(val.UpdateDate, 0).Format("2006-01-02"),
				})
			}

			resp.Data.Items = append(resp.Data.Items, item)
			wg.Done()
		}(&wg)
	}

	wg.Wait()

	resp.Code = 20000
	in.JSON(http.StatusOK, resp)

}

func AddGPZhouQiRemark(in *gin.Context) {
	var req trpc.AddGPZhouQiRemarkReq
	resp := &trpc.AddGPZhouQiRemarkResp{
		Code: 21000,
	}
	if err := in.BindJSON(&req); err != nil {
		in.JSON(http.StatusBadRequest, resp)
		return
	}

	result, err := orm.GPZhouQiMgr.FindOneBySecucode(req.GetSecucode())
	if err != nil {
		log.Errorf("get zhouqi failed: %q", err)
		in.JSON(http.StatusNotFound, resp)
		return
	}

	result.Remarks = append(result.Remarks, orm.GPZhouQiRemark{
		Remark:     req.GetRemark(),
		UpdateDate: time.Now().Unix(),
	})
	result.UpdateDate = time.Now().Unix()
	if _, err := result.Save(); err != nil {
		log.Errorf("save zhou qi failed: %s|%q", req.GetSecucode(), err)
		in.JSON(http.StatusNotModified, resp)
		return
	}

	resp.Code = 20000
	in.JSON(http.StatusOK, resp)
}