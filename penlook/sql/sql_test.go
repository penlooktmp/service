// Copyright 2014 Penlook Development Team. All rights reserved.
// Use of this source code is governed by
// license that can be found in the LICENSE file.
// Author : Loi Nguyen <loint@penlook.com>

package sql

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestSql(t *testing.T) {
	assert.New(t)
	sql := Sql {
		Name : "Penlook",
		Server: "localhost",
		Port: 3306,
		Database: "penlook",
		User: "root",
		Password: "123",
	}.Connect()

	sql.Run (
		`DROP TABLE IF EXISTS test`,
		`CREATE TABLE test (
			id INT(6) UNSIGNED AUTO_INCREMENT PRIMARY KEY,
			firstname VARCHAR(30) NOT NULL,
			lastname VARCHAR(30) NOT NULL,
			email VARCHAR(50)
		)`,
	)
}
