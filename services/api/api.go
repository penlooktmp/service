package main

import (
	"github.com/penlook/gin"
	"log"
	"net/http"
	"time"
)

func Api() {

	log.Println("Register API")

	//gin.SetMode(gin.ReleaseMode)
	gin.SetMode(gin.DebugMode)

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
