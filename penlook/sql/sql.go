// Copyright 2014 Penlook Development Team. All rights reserved.
// Use of this source code is governed by
// license that can be found in the LICENSE file.
// Author : Loi Nguyen <loint@penlook.com>

package sql

import (
	"github.com/penlook/gorm"
	_ "github.com/penlook/mysql"
)

type Sql struct {
	Name string
}

func (sql Sql) Connect() gorm.DB {
	connection, err := gorm.Open("mysql", "root@tcp(localhost:3306)/test?charset=utf8&parseTime=True")

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
