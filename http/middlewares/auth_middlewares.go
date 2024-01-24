package middlewares

import (
	"fmt"
	"myapp/helpers"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
)

func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		tokenString := c.GetHeader("Authorization")
		if tokenString == "" {
			fmt.Println("aaa")
			c.JSON(http.StatusUnauthorized, nil)
			c.Abort()
			return
		}

		authTokens := strings.Split(tokenString, " ")
		if authTokens[0] != "Bearer" {
			fmt.Println("bbb")
			c.JSON(http.StatusUnauthorized, nil)
			c.Abort()
			return
		}

		token, err := helpers.ParseJWTToken(authTokens[1])
		if err != nil || !token.Valid {
			fmt.Println("ccc")
			c.JSON(http.StatusUnauthorized, nil)
			c.Abort()
			return
		}

		claims, ok := token.Claims.(jwt.MapClaims)
		if !ok {
			fmt.Println("ddd")
			c.JSON(http.StatusUnauthorized, nil)
			c.Abort()
			return
		}

		userID, exists := claims["user_id"].(string)
		if !exists {
			fmt.Println("eee")
			c.JSON(http.StatusUnauthorized, nil)
			c.Abort()
			return
		}

		c.Set("user_id", userID)
		c.Next()
	}
}
