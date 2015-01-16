// Copyright 2014 Penlook Development Team. All rights reserved.
// Use of this source code is governed by
// license that can be found in the LICENSE file.
// Author : Loi Nguyen <loint@penlook.com>

package sql

import (
	"fmt"
	"github.com/penlook/gorm"
	_ "github.com/penlook/mysql"
)

type Sql struct {
	Name     string
	Server   string
	Port     int
	Database string
	Charset  string
	Username string
	Password string
}

func (sql Sql) Connect() gorm.DB {

	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=%s&parseTime=True",
		sql.Username,
		sql.Password,
		sql.Server,
		sql.Port,
		sql.Database,
		sql.Charset,
	)

	connection, err := gorm.Open("mysql", dsn)

	if err != nil {
		panic(err)
	}

	connection.DB()
	err = connection.DB().Ping()

	if err != nil {
		panic(err)
	}

	connection.DB().SetMaxIdleConns(10)
	connection.DB().SetMaxOpenConns(100)
	connection.SingularTable(true)

	return connection
}
