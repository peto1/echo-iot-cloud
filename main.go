package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
	ginSwagger "github.com/swaggo/gin-swagger"
	"github.com/swaggo/gin-swagger/swaggerFiles"

	"github.com/airren/echo-iot-backend/config"
	"github.com/airren/echo-iot-backend/dal"
	"github.com/airren/echo-iot-backend/docs"
	"github.com/airren/echo-iot-backend/middleware"
	"github.com/airren/echo-iot-backend/router"
)

// @title Echo IoT Cloud
// @version 0.0.1
// @description PLC manager on Cloud
// @contact.name Airren
// @contact.email renqiang@outlook.com
// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html
// @host localhost:8080
// @BasePath /
func main() {

	config.AuthInit()
	//r := gin.Default()
	docs.SwaggerInfo.BasePath = "/api/v1"

	r := gin.New()
	r.Use(gin.Logger(), gin.Recovery())
	router.UserAPI(r)
	r.Use(middleware.AuthMiddleware())
	err := dal.InitMySQL()
	if err != nil {
		panic(err)
	}

	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	r.GET("/health", func(c *gin.Context) {
		val := c.Query("val")
		c.String(http.StatusOK, "ok: "+val)
	})

	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	router.ApiV_1(r)
	_ = r.Run()

}
