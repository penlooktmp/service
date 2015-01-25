// Copyright 2014 Penlook Development Team. All rights reserved.
// Use of this source code is governed by
// license that can be found in the LICENSE file.
// Author : Loi Nguyen <loint@penlook.com>

package main

import (
	"github.com/penlook/daemon"
	sql "github.com/penlook/service/modules/sql"
)

func main() {

	service := daemon.Service{
		Name:        "api",
		Description: "Penlook API Service",
		Process:     Api,
	}

	service.Initialize()
}

type User struct {
	Id       int64
	Username string `sql:"type:varchar(100);"`
	Email    string `sql:"type:varchar(100);"`
	Password string `sql:"type:varchar(200);"`
}

func Api() {

	sql := sql.Sql{
		Name:     "Penlook",
		Server:   "localhost",
		Port:     3306,
		Database: "test",
		Charset:  "utf8",
		Username: "root",
	}.Connect()

	sql.DropTableIfExists(&User{})
	// Api service in here
}
