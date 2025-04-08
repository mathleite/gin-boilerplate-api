package middlewares

import (
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	utils "github.com/mathleite/gin-boilerplate-api/pkg/jwt"
)

func AuthMiddleware() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		authHeader := ctx.GetHeader("Authorization")
		if authHeader == "" {
			ctx.JSON(http.StatusUnauthorized, gin.H{"error": "Authorization header is required"})
			ctx.Abort()
			return
		}

		tokenString := strings.TrimPrefix(authHeader, "Bearer ")
		if tokenString == authHeader {
			ctx.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid authorization header format"})
			ctx.Abort()
			return
		}

		token, err := utils.ParseToken(tokenString)

		if err != nil {
			ctx.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid or expired token"})
			ctx.Abort()
			return
		}

		claims := *token
		ctx.Set("user_id", claims["user_id"])
		ctx.Next()
	}
}
