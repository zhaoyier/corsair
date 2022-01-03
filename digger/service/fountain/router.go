package fountain

import (
	"net/http"
	"time"

	admin "git.ezbuy.me/ezbuy/corsair/digger/service/fountain/internal/zwadmin"
	"git.ezbuy.me/ezbuy/corsair/digger/service/fountain/longline"
	"git.ezbuy.me/ezbuy/corsair/digger/service/fountain/prompt"
	"git.ezbuy.me/ezbuy/corsair/digger/service/fountain/recommend"
	"git.ezbuy.me/ezbuy/corsair/digger/service/internal/model"
	"github.com/ezbuy/ezorm/db"
	"github.com/gin-gonic/gin"
)

func startRoute(router *gin.Engine) {
	model.MgoSetup(&db.MongoConfig{
		MongoDB: "mongodb://digger:digger_ppwd@81.69.250.236:12302/digger",
		DBName:  "digger",
	})

	router.GET("/", func(c *gin.Context) {
		time.Sleep(time.Second)
		c.String(http.StatusOK, "Welcome Gin Server")
	})
	// 推荐列表
	router.GET("/api/digger.Fountain/GPRecommendList", recommend.GPRecommendList)
	adminGroup := router.Group("/api/backend")
	{
		adminGroup.POST("/login", admin.Login)
		adminGroup.POST("/logout", admin.Logout)
		adminGroup.GET("/info", admin.UserInfo)
		adminGroup.POST("/updateCNConfig", admin.UpdateCNConfig)
	}
	stockGroup := router.Group("/api/stock")
	{
		stockGroup.POST("/GetRecommend", recommend.GetRecommend)
		stockGroup.POST("/UpdateRecommend", recommend.UpdateRecommend)
		stockGroup.POST("/PromptBuyList", prompt.PromptBuyList)
		stockGroup.POST("/GetLongLineList", longline.GetLongLineList)
		stockGroup.POST("/GetDailyList", admin.GetDailyList)
		stockGroup.POST("/ManualDecreaseList", admin.ManualDecreaseList)
		stockGroup.POST("/GetFocusList", admin.GetFocusList)
		stockGroup.POST("/ConfirmFocus", admin.ConfirmFocus)
		stockGroup.POST("/CancelFocus", admin.CancelFocus)
		stockGroup.POST("/updateFocus", admin.UpdateFocus)

	}
}
