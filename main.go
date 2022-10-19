package main

import (
	"github.com/Jayleonc/golang-jwt-project/routes"
	"github.com/gin-gonic/gin"
	"os"
)

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "8000"
	}
	r := gin.New()
	r.Use(gin.Logger())

	routes.AuthRoutes(r)
	routes.UserRoutes(r)

	r.GET("/api-1", func(c *gin.Context) {
		c.JSON(200, gin.H{"success": "api-1"})
	})

	r.GET("/api-2", func(c *gin.Context) {
		c.JSON(200, gin.H{"success": "api-2"})
	})

	r.Run(":" + port)

}
