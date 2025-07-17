package routes

import (
	"Marketplace/internal/controllers"
	"Marketplace/internal/middleware"

	"github.com/gin-gonic/gin"
)

func SetupRoutes(router *gin.Engine) {
	router.POST("/register", controllers.Register)
	router.POST("/login", controllers.Login)

	protected := router.Group("/")
	protected.Use(middleware.AuthMiddleware())
	{
		protected.POST("/ads", controllers.CreateAd)
	}

	router.GET("/ads", controllers.GetAds)
}
