package middlewares

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/usmonali4/jwt_auth/utils/token"
)

func JwtAuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		err := token.TokenValid(c)
		if err != nil {
			c.String(http.StatusUnauthorized, "Unauthorized")
			c.Abort()
			return
		}
		c.Next()
	}
}