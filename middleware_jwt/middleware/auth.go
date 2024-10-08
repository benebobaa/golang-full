package middleware

import (
	"middleware_jwt/pkg"
	"net/http"

	"github.com/gin-gonic/gin"
)

const ClaimsKey = "claims"

func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		tokenString := c.GetHeader("Authorization")
		if tokenString == "" {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Authorization header is required"})
			c.Abort()
			return
		}

		if len(tokenString) > 7 && tokenString[:7] == "Bearer " {
			tokenString = tokenString[7:]
		}

		claims, err := pkg.ValidateToken(tokenString)
		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid or expired token"})
			c.Abort()
			return
		}

		payload := pkg.UserInfo{
			ID:    claims.User.ID,
			Email: claims.User.Email,
		}

		c.Set(ClaimsKey, payload)
		c.Next()
	}
}
