package middlewares

import (
	"strings"

	"github.com/fireab/goapi2/utils"
	"github.com/gin-gonic/gin"
)

// AuthMiddleware
func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		tokenString := c.GetHeader("Authorization")

		if tokenString == "" || !strings.HasPrefix(tokenString, "Bearer ") {
			c.JSON(401, gin.H{"message": "Token not Found!"})
			c.Abort()
			return
		}
		decode, err := utils.ValidateJWT(tokenString[7:])

		if err != nil {
			c.JSON(401, gin.H{"message": "Unauthorized"})
			c.Abort()

			return
		}
		c.Set("user_id", decode["id"])
		c.Next()
	}
}
