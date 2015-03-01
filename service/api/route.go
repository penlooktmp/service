package main

import (
	// "github.com/penlook/gin"
	"github.com/gin-gonic/gin"
)

type Router struct {
	Handler *gin.Engine
}

func (router Router) GetHandler() *gin.Engine {
	return router.Handler
}

func (router Router) Register() {
	router.Status("/status")
	router.Comment("/comment")
	router.Activity("/activity")
}

func (router Router) Status(root string) {
	route := router.Handler
	route.POST(root, postStatus)
	route.GET(root, getAllStatus)
	route.GET(root+"/:id", getStatus)
	route.PUT(root+"/:id", updateStatus)
	route.DELETE(root+"/:id", deleteStatus)
}

func (router Router) Comment(root string) {
	route := router.Handler
	route.GET(root, getAllComments)
	route.GET(root+"/:id", getComment)
}

func (router Router) Activity(root string) {
	route := router.Handler
	route.GET(root, getAllActivities)
	route.GET(root+"/:id", getActivity)
}
