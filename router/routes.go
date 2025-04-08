package router

import (
	"github.com/gin-gonic/gin"
	"github.com/mathleite/gin-boilerplate-api/handlers"
	"github.com/mathleite/gin-boilerplate-api/router/middlewares"
)

func initializeRoutes(router *gin.Engine) {
	router.GET("/health", handlers.HealthCheck)

	v1 := router.Group("/api/v1")
	{
		v1.POST("/register", handlers.Register)
		v1.POST("/auth/signin", handlers.Signin)

		transactions := v1.Group("/transactions").Use(middlewares.AuthMiddleware())
		{
			transactions.GET("/", func(c *gin.Context) {
				c.JSON(200, gin.H{"message": "The middleware is working!"})
			})
		}
	}

}
