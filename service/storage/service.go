// Copyright 2014 Penlook Development Team. All rights reserved.
// Use of this source code is governed by
// license that can be found in the LICENSE file.
// Author : Tin Nguyen <tinntt@penlook.com>

package main

import (
	"github.com/penlook/daemon"
)

func main() {

	service := daemon.Service{
		Name:        "storage",
		Description: "Penlook Storage Service",
		Process:     Storage,
	}

	service.Initialize()
}
