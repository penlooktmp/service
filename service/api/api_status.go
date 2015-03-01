package main

import (
	// "github.com/penlook/gin"
	"github.com/gin-gonic/gin"
)

func postStatus(c *gin.Context) {
	c.JSON(200, gin.H{
		"status": "Get One Status",
	})
}

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

func updateStatus(c *gin.Context) {
	c.JSON(200, gin.H{
		"status": "Get One Status",
	})
}

func deleteStatus(c *gin.Context) {
	c.JSON(200, gin.H{
		"status": "Get One Status",
	})
}
