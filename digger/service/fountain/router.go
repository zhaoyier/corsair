package fountain

import (
	"net/http"
	"time"

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
}
