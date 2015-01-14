// Copyright 2014 Penlook Development Team. All rights reserved.
// Use of this source code is governed by
// license that can be found in the LICENSE file.
// Author : Loi Nguyen <loint@penlook.com>

package sql

import (
	libsql "database/sql"
	_ "github.com/penlook/mysql"
	"strconv"
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
	Query func(cmd string, args ...interface {}) (*libsql.Rows, error)
}

func (sql *Sql) SetConnection(connection *libsql.DB) {
	sql.connection = connection
	sql.Query = connection.Query
}

func (sql Sql) Connect() Sql {
	
	connect_str := sql.User

	if sql.Password != "" {
		connect_str += ":" + sql.Password
	}

	connect_str += "@" + sql.Server + ":" + strconv.Itoa(sql.Port)

	if sql.Database != "" {
		connect_str += "/" + sql.Database
	}

	connection, err := libsql.Open("mysql", "root@localhost")

	if err != nil {
		panic(err)
	}

	sql.SetConnection(connection)

 	return sql
}

func (sql Sql) Run(args ... string) (*libsql.Rows, error) {

	if len(args) == 1 {
		return sql.Query(args[0])
	} 	

	for _, cmd := range args {
		sql.Query(cmd)
	}

	return nil, nil
}