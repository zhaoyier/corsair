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

func UpdateCNConfig(in *gin.Context) {
	var req trpc.UpdateCNConfigReq
	resp := &trpc.UpdateCNConfigResp{
		Code: 21000,
	}
	if err := in.BindJSON(&req); err != nil {
		in.JSON(http.StatusBadRequest, resp)
		return
	}

	result, _ := orm.CNConfigMgr.FindOne(ezdb.M{})
	if result == nil {
		result = orm.CNConfigMgr.NewCNConfig()
	}
	if req.GetDecreaseTag() > 0 {
		result.DecreaseTag = req.GetDecreaseTag()
	}
	if req.GetDecreasePeriod() > 0 {
		result.DecreasePeriod = req.GetDecreasePeriod()
	}

	result.UpdateDate = time.Now().Unix()
	if _, err := result.Save(); err != nil {
		log.Errorf("save config failed: %q", err)
		in.JSON(http.StatusFailedDependency, resp)
		return
	}

	if req.GetDecreaseTag() > 0 || req.GetDecreasePeriod() > 0 {
		go resetGPRecommend()
	}

	resp.Code = 20000
	in.JSON(http.StatusOK, resp)
}

// TODO 待完善
func resetGPRecommend() {

}
