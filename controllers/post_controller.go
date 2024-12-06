package controllers

import (
	"net/http"
	"social-media-api/database"
	"social-media-api/models"

	"github.com/gin-gonic/gin"
)

func GetPosts(c *gin.Context) {
	var posts []models.Post
	database.DB.Preload("Comments").Find(&posts)
	c.JSON(http.StatusOK, posts)
}

func CreatePost(c *gin.Context) {
	var post models.Post
	if err := c.ShouldBindJSON(&post); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	database.DB.Create(&post)
	c.JSON(http.StatusOK, post)
}
