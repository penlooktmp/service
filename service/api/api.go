package main

import (
	// "github.com/penlook/gin"
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)

func Api() {

	//gin.SetMode(gin.ReleaseMode)
	gin.SetMode(gin.ReleaseMode)

	router := Router{
		Handler: gin.Default(),
	}

	router.Register()

	s := &http.Server{
		Addr:           ":8080",
		Handler:        router.GetHandler(),
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}

	s.ListenAndServe()
}
