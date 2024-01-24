package routes

import (
	"myapp/http/controllers"
	"myapp/http/middlewares"

	"github.com/gin-gonic/gin"
)

func userRoutes(r *gin.Engine) {
	group := r.Group("/api/account")
	group.POST("/register", controllers.UserRegister)
	group.POST("/login", controllers.UserLogin)

	authGroup := r.Group("/api/account")
	authGroup.Use(middlewares.AuthMiddleware())

	authGroup.POST("/logout", controllers.UserLogout)
	authGroup.GET("/profile", controllers.UserDetail)
	authGroup.PUT("/profile", controllers.UserUpdate)
}
