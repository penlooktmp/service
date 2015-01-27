package main

import (
	"github.com/penlook/gin"
)

type Router struct {
	Handler *gin.Engine
}

func (router Router) GetHandler() *gin.Engine {
	return router.Handler
}

func (router Router) Register() {
	router.Status()
	router.Comment()
	router.Activity()
}

func (router Router) Status() {
	route := router.Handler
	route.GET("/status", getAllStatus)
	route.GET("/status/:id", getStatus)
}

func (router Router) Comment() {
	route := router.Handler
	route.GET("/comment", getAllComments)
	route.GET("/comment/:id", getComment)
}

func (router Router) Activity() {
	route := router.Handler
	route.GET("/activity", getAllActivities)
	route.GET("/activity/:id", getActivity)
}
