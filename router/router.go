package router

import (
	"github.com/gin-gonic/gin"

	"github.com/airren/echo-iot-backend/controller"
)

func ApiV_1(r *gin.Engine) {

	task := r.Group("api/task")

	task.GET("/:id", controller.GetOrderById)
	task.POST("/create", controller.CreateOrder)
	task.PUT("/update", controller.UpdateOrder)
	task.POST("/list", controller.QueryOrders)

}

func UserAPI(r *gin.Engine) {
	user := r.Group("api/user")
	user.GET("/login", controller.UserLogin)
	user.GET("/info", controller.UserInfo)
	user.PUT("/logout", controller.UserLogout)
}
