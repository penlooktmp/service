package main

import (
	"./module/daemon"
	"./module/penlook"
)

func main() {

	service := daemon.Service{
		Name:        "api",
		Description: "Penlook API Service",
		Process:     penlook.Api,
	}

	service.Initialize()
}
