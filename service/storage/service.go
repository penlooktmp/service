// Copyright 2014 Penlook Development Team. All rights reserved.
// Use of this source code is governed by
// license that can be found in the LICENSE file.
// Author : Loi Nguyen <loint@penlook.com>

package main

import (
	// "github.com/penlook/socket"
	"github.com/penlook/daemon"
)

func main() {

	service := daemon.Service{
		Process: S3Service,
	}

	service.GetInfo("s3")
	service.Initialize()
}
