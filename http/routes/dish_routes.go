package routes

import (
	"myapp/http/controllers"
	"myapp/http/middlewares"

	"github.com/gin-gonic/gin"
)

func dishRoutes(r *gin.Engine) {
	authGroup := r.Group("/api/dish")
	authGroup.Use(middlewares.AuthMiddleware())

	authGroup.GET("", controllers.DishList)
	authGroup.GET("/:id", controllers.DishDetail)
	authGroup.GET("/:id/rating/check", controllers.DishRatingCheck)
	authGroup.POST("/:id/rating", controllers.DishRating)
}
