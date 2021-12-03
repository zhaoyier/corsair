package recommend

import (
	"net/http"

	orm "git.ezbuy.me/ezbuy/corsair/digger/service/internal/model"
	trpc "git.ezbuy.me/ezbuy/corsair/digger/service/internal/rpc"
	"git.ezbuy.me/ezbuy/corsair/digger/service/internal/utils"
	log "github.com/Sirupsen/logrus"
	ezdb "github.com/ezbuy/ezorm/db"
	"github.com/gin-gonic/gin"
)

func GPRecommendList(in *gin.Context) {
	query := ezdb.M{
		"Disabled": false,
		"State":    ezdb.M{"$gte": 1},
	}

	resp := &trpc.GPRecommendListResp{
		Rows: make([]*trpc.GPRecommend, 0),
	}
	results, err := orm.GPRecommendMgr.FindAll(query, "-HDecrease", "-MDecrease", "-TDecrease")
	if err != nil {
		log.Errorf("query recommend failed: %q", err)
	}
	for idx, result := range results {
		resp.Rows = append(resp.Rows, &trpc.GPRecommend{
			Id:         int32(idx + 1),
			Secucode:   result.Secucode,
			Name:       getName(result.Secucode),
			HDecrease:  result.Decrease,
			MaxPrice:   result.MaxPrice,
			MaxDay:     result.MaxDay,
			RMPrice:    result.RMPrice,
			GDDecrease: getGDDecrease(result.Secucode),
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

func getGDDecrease(secucode string) string {
	query := ezdb.M{
		"Secucode": secucode,
	}
	result, err := orm.GDLongLineMgr.FindOne(query, "-CreateDate")
	if err != nil {
		return ""
	}
	return result.GDReduceRatio
}
