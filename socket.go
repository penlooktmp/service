package main

import (
	"./penlook"
	"github.com/penlook/daemon"
)

func main() {

	service := daemon.Service{
		Name:        "socket",
		Description: "Penlook Socket Service",
		Process:     penlook.Socket,
	}

	service.Initialize()
}
