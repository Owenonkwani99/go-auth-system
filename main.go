package main

import (
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	// Public route
	r.POST("/login", LoginHandler)

	r.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{"message": "Fresh start working!"})
	})

	protected := r.Group("/api")
	protected.Use(AuthMiddleware())
	{
		protected.GET("/profile", func(c *gin.Context) {
			c.JSON(200, gin.H{
				"message": "Welcome to your profile!",
			})
		})
	}

	r.Run(":8080")
}
