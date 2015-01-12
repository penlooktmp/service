package main

import (
	"./module/daemon"
	"./module/penlook"
)

func main() {

	service := daemon.Service{
		Name:        "socket",
		Description: "Penlook Socket Service",
		Port:        9876,
		Process:     penlook.Socket,
	}

	// Run service in daemon
	service.Daemon()
}
