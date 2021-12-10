package recommend

import (
	"net/http"
	"strconv"

	orm "git.ezbuy.me/ezbuy/corsair/digger/service/internal/model"
	trpc "git.ezbuy.me/ezbuy/corsair/digger/service/internal/rpc"
	"git.ezbuy.me/ezbuy/corsair/digger/service/internal/utils"
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
			Id:         int32(idx + 1),
			Secucode:   result.Secucode,
			Name:       result.Name, //getName(result.Secucode),
			RMIndex:    result.RMIndex,
			PDecrease:  result.PDecrease,
			MaxPrice:   result.MaxPrice,
			MaxPDay:    utils.TS2Date(result.MaxPDay),
			RMPrice:    result.RMPrice,
			GDDecrease: result.GDDecrease,
			State:      getState(result.State),
			UpdateDate: utils.TS2Date(result.UpdateDate),
		})
	}
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
		return "进行中"
	case 4:
		return "结束"
	default:
		return "准备"
	}

}
