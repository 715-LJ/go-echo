package web

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"go-echo/internal/check"
	"go-echo/internal/conf"
	"log/slog"
)

// Gui - start web server
func Gui(dirPath, nodePath string) {

	confPath := dirPath + "/config_v1.yaml"
	check.Path(confPath)

	appConfig := conf.Get(confPath)
	fmt.Println(appConfig)

	address := appConfig.Host + ":" + appConfig.Port

	slog.Info("=================================== ")
	slog.Info("Web API at http://" + address)
	slog.Info("=================================== ")

	gin.SetMode(gin.ReleaseMode)

	router := gin.New()
	router.Use(gin.Recovery())

	router.POST("/api/word2pdf/", word2pdf)
	router.POST("/api/merge2pdf/", merge2pdf)

	err := router.Run(address)
	check.IfError(err)
}
