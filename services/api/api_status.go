package main

import (
	"github.com/penlook/gin"
)

func getAllStatus(c *gin.Context) {
	c.JSON(200, gin.H{
		"status": "Get All Status",
	})
}

func getStatus(c *gin.Context) {
	c.JSON(200, gin.H{
		"status": "Get One Status",
	})
}
