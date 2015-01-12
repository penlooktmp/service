package main

import (
	"github.com/penlook/daemon"
	"github.com/penlook/service/penlook"
)

func main() {

	service := daemon.Service{
		Name:        "api",
		Description: "Penlook API Service",
		Process:     penlook.Api,
	}

	service.Initialize()
}
