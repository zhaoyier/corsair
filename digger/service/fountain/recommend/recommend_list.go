package recommend

import (
	"net/http"
	"strconv"

	"git.ezbuy.me/ezbuy/corsair/digger/rpc/digger"
	orm "git.ezbuy.me/ezbuy/corsair/digger/service/internal/model"
	log "github.com/Sirupsen/logrus"
	ezdb "github.com/ezbuy/ezorm/db"
	"github.com/gin-gonic/gin"
)

func GetRecommendList(in *gin.Context) {
	resp := digger.GetRecommendListResp{
		Rows: make([]*digger.RecommendData, 0, 4),
	}

	limit, _ := strconv.Atoi(in.Query("limit"))
	offset, _ := strconv.Atoi(in.Query("offset"))
	resp.Total = int32(orm.GDHoldRecommendMgr.Count(ezdb.M{}))

	results, err := orm.GDHoldRecommendMgr.Find(ezdb.M{}, limit, offset, "-Level", "-EndDate")
	if err != nil {
		log.Errorf("query recommend failed: %q", err)
		in.JSON(http.StatusFailedDependency, gin.H{"status": "!ok"})
		return
	}

	for idx, item := range results {
		resp.Rows = append(resp.Rows, &digger.RecommendData{
			Id:            int32(idx + 1),
			Secucode:      item.Secucode,
			Level:         float32(item.Level),
			CumulantFocus: item.CumulantFocus,
			CumulantPrice: item.CumulantPrice,
			CumulantDate:  item.CumulantDate,
			EndDate:       item.EndDate,
			CreateDate:    item.CreateDate,
		})
	}

	in.JSON(http.StatusOK, resp)
}
