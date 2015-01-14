// Copyright 2014 Penlook Development Team. All rights reserved.
// Use of this source code is governed by
// license that can be found in the LICENSE file.
// Author : Loi Nguyen <loint@penlook.com>

package sql

import (
	libsql "database/sql"
	_ "github.com/penlook/mysql"
	//"fmt"
)

type Sql struct {
	Name string
	Server string
	Port int
	Database string
	User string
	Password string
	connection *libsql.DB
	Run func(cmd string, args ...interface {}) (*libsql.Rows, error)
}

func (sql *Sql) SetConnection(connection *libsql.DB) {
	sql.connection = connection
	sql.Run = connection.Query
}

func (sql Sql) Connect() Sql {
	connection_string := sql.
	connection, err := libsql.Open("mysql", "root@localhost")

	if err != nil {
		panic(err)
	}

	sql.SetConnection(connection)
 	return sql
}