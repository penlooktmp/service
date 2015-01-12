package main

import (
	"github.com/penlook/daemon"
	"github.com/penlook/service/penlook"
)

func main() {

	service := daemon.Service{
		Name:        "socket",
		Description: "Penlook Socket Service",
		Process:     penlook.Socket,
	}

	service.Initialize()
}
