package main

import (
	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"golang-test-task/meta/docs"
	"golang-test-task/meta/internal/crawler"
	"golang-test-task/meta/pkg/logging"
	"golang-test-task/meta/pkg/requests"
	_ "net/http/pprof"
)

// @title Do API
// @version 1.0.0
// @description API server for application
// @BasePath /

func main() {
	p := crawler.NewParser()
	serviceLogger := logging.NewToFile("serviceLogs")
	requestSender := requests.New()
	service := crawler.NewService(requestSender, p, serviceLogger)
	httpDelivery := crawler.NewHttp(service)

	middleware := crawler.NewMiddleware()

	router := gin.Default()
	router.Use(middleware.CORSMiddleware)
	router.Use(middleware.JSONLogMiddleware)

	docs.SwaggerInfo.BasePath = "/"
	router.GET("swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))
	router.POST("/walk", httpDelivery.ParseDataByUrls)

	log.Fatal(router.Run("0.0.0.0:8080"))
}
