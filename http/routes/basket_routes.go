package routes

import (
	"myapp/http/controllers"
	"myapp/http/middlewares"

	"github.com/gin-gonic/gin"
)

func basketRoutes(r *gin.Engine) {
	authGroup := r.Group("/api/basket")
	authGroup.Use(middlewares.AuthMiddleware())

	authGroup.GET("", controllers.BasketByUser)
	authGroup.POST("/dish/:dishId", controllers.BasketAddDish)
	authGroup.DELETE("/dish/:dishId", controllers.BasketDeleteDish)
}
