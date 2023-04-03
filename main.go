package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/RDelg/compare-encrypt-time/src/routers"
	setting "github.com/RDelg/compare-encrypt-time/src/settings"
	"github.com/gin-gonic/gin"
)

func init() {
	setting.Setup()
}

func main() {
	gin.SetMode(setting.AppSetting.ServerSetting.RunMode)

	routersInit := routers.InitRouter()
	endPoint := fmt.Sprintf(":%d", setting.AppSetting.ServerSetting.HttpPort)
	maxHeaderBytes := 1 << 20

	server := &http.Server{
		Addr:           endPoint,
		Handler:        routersInit,
		MaxHeaderBytes: maxHeaderBytes,
	}

	log.Printf("[info] start http server listening %s", endPoint)

	server.ListenAndServe()

}
