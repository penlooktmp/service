// Copyright 2014 Penlook Development Team. All rights reserved.
// Use of this source code is governed by
// license that can be found in the LICENSE file.
// Author : Loi Nguyen <loint@penlook.com>

package penlook

import (
	"github.com/penlook/gin"
)

func Api() {
	router := gin.Default()
	router.GET("/", func(c *gin.Context) {
		c.String(200, "hello world")
	})
	router.GET("/ping", func(c *gin.Context) {
		c.String(200, "pong")
	})
	router.POST("/submit", func(c *gin.Context) {
		c.String(401, "not authorized")
	})
	router.PUT("/error", func(c *gin.Context) {
		c.String(500, "and error hapenned :(")
	})
	router.Run(":8080")
}
