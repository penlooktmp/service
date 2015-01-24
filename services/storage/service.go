// Copyright 2014 Penlook Development Team. All rights reserved.
// Use of this source code is governed by
// license that can be found in the LICENSE file.
// Author : Tin Nguyen <tinntt@penlook.com>

package main

import (
	"github.com/penlook/daemon"
	"github.com/penlook/service/module"
)

func main() {

	service := daemon.Service{
		Name:        "storage",
		Description: "Penlook API Service",
		Process:     module.Storage,
	}

	service.Initialize()
}

func Storage() {
	// Storage service in here
}
