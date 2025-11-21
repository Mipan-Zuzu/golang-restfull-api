package main

import (
	"fmt"
	"golang-new/controllers"
	"golang-new/models"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	router.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{"message": "pong"})
	})

	for _, r := range router.Routes() {
		fmt.Println(r.Method, r.Path)
	}

	models.ConnectDatabase()

	router.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "hai broe",
		})
	})
	router.Run(":3000")

	router.GET("/api/post", controllers.GetPost)
}
