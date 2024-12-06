package main

import (
	"social-media-api/controllers"
	"social-media-api/database"
	"social-media-api/models"

	"github.com/gin-gonic/gin"
)

func main() {
	// Initialize Database
	database.InitDatabase()
	database.DB.AutoMigrate(&models.User{}, &models.Post{}, &models.Comment{})

	// Setup Router
	r := gin.Default()

	// Routes
	r.GET("/users", controllers.GetUsers)
	r.POST("/users", controllers.CreateUser)

	r.GET("/posts", controllers.GetPosts)
	r.POST("/posts", controllers.CreatePost)

	r.GET("/comments", controllers.GetComments)
	r.POST("/comments", controllers.CreateComment)

	// Start Server
	r.Run(":8080")
}
