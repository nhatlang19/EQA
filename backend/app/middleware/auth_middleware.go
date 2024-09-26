package middleware

import (
	"net/http"
	"os"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
)

func JWTMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		jwtKey := os.Getenv("SECRET_KEY")
		authHeader := c.GetHeader("Authorization")

		if c.Request.URL.Path == "/api/auth/login" {
			c.Next()
			return
		}

		if authHeader == "" {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"Error": "Authorization header is missing"})
			return
		}

		const prefix = "Bearer "
		if !strings.HasPrefix(authHeader, prefix) {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"Error": "Invalid Authorization header format"})
			return
		}

		token, err := jwt.ParseWithClaims(authHeader[len(prefix):], &jwt.MapClaims{}, func(token *jwt.Token) (interface{}, error) {
			return []byte(jwtKey), nil
		})

		if err != nil || !token.Valid {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"Error": "Invalid token"})
			return
		}

		claims := token.Claims.(*jwt.MapClaims)
		c.Set("user", claims)

		c.Next()
	}
}
