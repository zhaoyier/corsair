package fountain

import (
	"log"
	"net/http"
	"os"
	"os/signal"

	"git.ezbuy.me/ezbuy/corsair/digger/common"
	"github.com/gin-gonic/gin"
)

func Start() {
	router := gin.Default()
	router.Use(common.Cors())
	startRoute(router)

	server := &http.Server{
		Addr:    ":12303",
		Handler: router,
	}

	quit := make(chan os.Signal)
	signal.Notify(quit, os.Interrupt)

	go func() {
		<-quit
		log.Println("receive interrupt signal")
		if err := server.Close(); err != nil {
			log.Fatal("Server Close:", err)
		}
	}()

	if err := server.ListenAndServe(); err != nil {
		if err == http.ErrServerClosed {
			log.Println("Server closed under request")
		} else {
			log.Fatal("Server closed unexpect")
		}
	}

	log.Println("Server exiting")
}

func Start2() {
	router := gin.Default()
	router.Use(common.Cors())
	go startRoute(router)

	pem := "/home/gopath/crt/ibet/ibet.sale_bundle.pem"
	key := "/home/gopath/crt/ibet/ibet.sale.key"
	router.RunTLS(":8080", pem, key)
}
