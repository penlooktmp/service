package main

import (
	"./module/daemon"
	"./module/penlook"
)

func main() {

	service := daemon.Service{
		Name:        "penlook-api",
		Description: "Penlook API Service",
		Port:        9876,
		Process:     penlook.Api,
	}

	// Run service in daemon
	service.Daemon()
}
