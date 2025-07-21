package main

import (
	"Marketplace/internal/config"
	"Marketplace/internal/routes"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	godotenv.Load()
	config.Connect()

	r := gin.Default()
	routes.SetupRoutes(r)

	r.Run(":8080")
}
