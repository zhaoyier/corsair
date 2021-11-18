package fountain

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

func startRoute(router *gin.Engine) {
	router.GET("/", func(c *gin.Context) {
		time.Sleep(time.Second)
		c.String(http.StatusOK, "Welcome Gin Server")
	})
}
