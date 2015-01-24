// Copyright 2014 Penlook Development Team. All rights reserved.
// Use of this source code is governed by
// license that can be found in the LICENSE file.
// Author : Loi Nguyen <loint@penlook.com>

package main

import (
	"github.com/penlook/daemon"
	"github.com/penlook/service/module"
)

func main() {

	service := daemon.Service{
		Name:        "api",
		Description: "Penlook API Service",
		Process:     Api,
	}

	service.Initialize()
}

func Api() {
	// Api service in here
}
