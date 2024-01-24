package routes

import (
	"myapp/http/controllers"
	"myapp/http/middlewares"

	"github.com/gin-gonic/gin"
)

func orderRoutes(r *gin.Engine) {
	authGroup := r.Group("/api/order")
	authGroup.Use(middlewares.AuthMiddleware())

	authGroup.GET("/:id", controllers.OrderDetail)
	authGroup.GET("", controllers.OrderByUser)
	authGroup.POST("", controllers.OrderCreate)
	authGroup.POST("/:id/status", controllers.OrderConfirm)
}
