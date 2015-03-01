package main

import (
	// "github.com/penlook/gin"
	"github.com/gin-gonic/gin"
)

func getAllComments(c *gin.Context) {
	c.JSON(200, gin.H{
		"status": "Get All Comment",
	})
}

func getComment(c *gin.Context) {
	c.JSON(200, gin.H{
		"status": "Get One Comment",
	})
}
