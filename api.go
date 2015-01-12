package main

import (
	"./penlook"
	"github.com/penlook/daemon"
)

func main() {

	service := daemon.Service{
		Name:        "api",
		Description: "Penlook API Service",
		Process:     penlook.Api,
	}

	service.Initialize()
}
