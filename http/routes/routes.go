package routes

import "github.com/gin-gonic/gin"

func InitRoutes(r *gin.Engine) {
	basketRoutes(r)
	dishRoutes(r)
	orderRoutes(r)
	userRoutes(r)
}
