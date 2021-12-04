package fountain

import (
	"git.ezbuy.me/ezbuy/corsair/digger/common"
	"github.com/gin-gonic/gin"
)

func Start() {
	router := gin.Default()
	router.Use(common.Cors())
	go startRoute(router)

	pem := "/home/gopath/crt/ibet/ibet.sale_bundle.pem"
	key := "/home/gopath/crt/ibet/ibet.sale.key"
	router.RunTLS(":8080", pem, key)
}
