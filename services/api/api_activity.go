package main

import (
	"github.com/penlook/gin"
)

func getAllActivities(c *gin.Context) {
	c.JSON(200, gin.H{
		"status": "Get All Activities",
	})
}

func getActivity(c *gin.Context) {
	c.JSON(200, gin.H{
		"status": "Get One Activity",
	})
}
