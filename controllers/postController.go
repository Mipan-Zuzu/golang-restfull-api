package controllers

import (
	"golang-new/models"

	"github.com/gin-gonic/gin"
)

func GetPost(c *gin.Context) {
	var posts []models.Post
	models.DB.Find(&posts)

	c.JSON(200, gin.H{
		"success": true,
		"message": "success",
		"data":    posts,
	})
}
