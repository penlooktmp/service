package main

import (
	"./module/daemon"
	"./module/penlook"
)

func main() {

	service := daemon.Service{
		Name:        "socket",
		Description: "Penlook Socket Service",
		Process:     penlook.Socket,
	}

	service.Initialize()
}
