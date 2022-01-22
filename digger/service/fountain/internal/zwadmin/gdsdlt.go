package zwadmin

import (
	"net/http"
	"sort"
	"sync"

	orm "git.ezbuy.me/ezbuy/corsair/digger/service/internal/model"
	trpc "git.ezbuy.me/ezbuy/corsair/digger/service/internal/rpc"
	"git.ezbuy.me/ezbuy/corsair/digger/service/internal/utils"
	log "github.com/Sirupsen/logrus"
	ezdb "github.com/ezbuy/ezorm/db"
	"github.com/gin-gonic/gin"
	"gopkg.in/mgo.v2/bson"
)

func GetGDSDLT(in *gin.Context) {
	var req trpc.GetGDSDLTReq
	resp := &trpc.GetGDSDLTResp{
		Code: 21000,
		Data: &trpc.GDSDLTData{
			Items: make([]*trpc.GDSDLTItem, 0),
		},
	}
	if err := in.BindJSON(&req); err != nil {
		in.JSON(http.StatusBadRequest, resp)
		return
	}

	match := ezdb.M{}
	limit := int(req.GetLimit())
	offset := int(req.GetOffset())
	group := ezdb.M{"_id": "$Secucode", "total": ezdb.M{"$sum": "$FreeHoldnumRation"}}

	if req.GetKeyword() != "" {
		match["HolderName"] = bson.RegEx{Pattern: req.GetKeyword(), Options: "i"}
	}
	if req.GetReleaseStart() > 0 && req.GetReleaseEnd() > 0 {
		match["EndDate"] = ezdb.M{"$gte": req.GetReleaseStart() / 1000, "$lte": req.GetReleaseEnd() / 1000}
	}

	sess, col := orm.GDTopTenMgr.GetCol()
	defer sess.Close()

	results := []GDSDLTItemMsg{}
	pipe := col.Pipe([]bson.M{
		{"$match": match},
		{"$group": group},
		{"$sort": bson.M{"total": -1}},
	})

	if err := pipe.All(&results); err != nil {
		log.Errorf("get prompt buy failed: %q", err)
		in.JSON(http.StatusNotFound, resp)
		return
	}

	max := offset + limit
	resp.Data.Total = int32(len(results))

	if len(results) > max {
		results = results[offset:max]
	} else if len(results) > offset {
		results = results[offset:]
	} else {
		results = results[:0]
	}

	wg := sync.WaitGroup{}
	wg.Add(len(results))

	for _, result := range results {
		result := result
		go func(wg *sync.WaitGroup) {
			resp.Data.Items = append(resp.Data.Items, &trpc.GDSDLTItem{
				Name:       getName(result.Secucode),
				Secucode:   result.Secucode,
				FreeRation: utils.TruncateFloat(result.FreeRation),
				Focused:    getFocusBySecucode(result.Secucode),
			})
			wg.Done()
		}(&wg)
	}

	wg.Wait()

	sort.Slice(resp.Data.Items, func(i, j int) bool {
		return resp.Data.Items[i].FreeRation > resp.Data.Items[j].FreeRation
	})

	resp.Code = 20000
	in.JSON(http.StatusOK, resp)

}

func GetGDSDLTDetail(in *gin.Context) {
	var req trpc.GetGDSDLTDetailReq
	resp := &trpc.GetGDSDLTDetailResp{
		Code: 21000,
		Data: &trpc.GDSDLTDetailData{
			Items: make([]*trpc.GDSDLTDetailItem, 0),
		},
	}
	if err := in.BindJSON(&req); err != nil {
		in.JSON(http.StatusBadRequest, resp)
		return
	}

	if req.GetName() == "" && req.GetSecucode() == "" {
		in.JSON(http.StatusBadRequest, resp)
		return
	}

	query := ezdb.M{}
	limit := int(req.GetLimit())
	offset := int(req.GetOffset())
	if req.GetName() != "" {
		query["Secucode"] = getSecucode(req.GetName())
	}
	if req.GetSecucode() != "" {
		query["Secucode"] = req.GetSecucode()
	}

	results, err := orm.GDTopTenMgr.Find(query, limit, offset, "-EndDate")
	if err != nil {
		log.Errorf("get prompt buy failed: %q", err)
		in.JSON(http.StatusNotFound, resp)
		return
	}
	wg := sync.WaitGroup{}
	wg.Add(len(results))

	for _, result := range results {
		result := result
		go func(wg *sync.WaitGroup) {
			resp.Data.Items = append(resp.Data.Items, &trpc.GDSDLTDetailItem{
				Name:          getName(result.Secucode),
				Secucode:      result.Secucode,
				EndDate:       result.EndDate,
				HolderRank:    result.HolderRank,
				HolderName:    result.HolderName,
				HolderType:    result.HolderType,
				HoldNum:       result.HoldNum,
				HoldnumRation: result.HoldnumRation,
				HoldNumChange: utils.String2I32(result.HoldNumChange),
				CreateDate:    result.CreateDate,
			})
			wg.Done()
		}(&wg)
	}

	wg.Wait()

	sort.Slice(resp.Data.Items, func(i, j int) bool {
		return resp.Data.Items[i].HolderRank < resp.Data.Items[j].HolderRank
	})

	resp.Code = 20000
	in.JSON(http.StatusOK, resp)
}

type GDSDLTItemMsg struct {
	Secucode   string  `bson:"_id"`
	FreeRation float64 `bson:"total"`
}
