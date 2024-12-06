package controllers

import (
	"net/http"
	"social-media-api/database"
	"social-media-api/models"

	"github.com/gin-gonic/gin"
)

func GetComments(c *gin.Context) {
	var comments []models.Comment
	database.DB.Find(&comments)
	c.JSON(http.StatusOK, comments)
}

func CreateComment(c *gin.Context) {
	var comment models.Comment
	if err := c.ShouldBindJSON(&comment); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	database.DB.Create(&comment)
	c.JSON(http.StatusOK, comment)
}
