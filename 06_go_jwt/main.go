package main

import (
	"log"
	"os"

	"github.com/JagdeepSingh13/go_jwt/middleware"
	"github.com/JagdeepSingh13/go_jwt/routes"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	port := os.Getenv("PORT")
	if port == "" {
		port = "8000"
	}

	router := gin.New()
	router.Use(gin.Logger())

	routes.AuthRoutes(router)
	// routes.UserRoutes(router)

	router.GET("/api-1", func(c *gin.Context) {
		c.JSON(200, gin.H{"success": "Access granted for api-1"})
	})

	router.GET("/api-2", func(c *gin.Context) {
		c.JSON(200, gin.H{"success": "Access granted for api-2"})
	})

	// Protected routes
	protected := router.Group("/")
	protected.Use(middleware.Authenticate())
	routes.UserRoutes(protected)

	router.Run(":" + port)
}
