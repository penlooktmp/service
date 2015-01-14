// Copyright 2014 Penlook Development Team. All rights reserved.
// Use of this source code is governed by
// license that can be found in the LICENSE file.
// Author : Loi Nguyen <loint@penlook.com>

package sql

import (
	"crypto/md5"
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

type User struct {
	Id       int64
	Username string `sql:"type:varchar(100);"`
	Email    string `sql:"type:varchar(100);"`
	Password string `sql:"type:varchar(200);"`
}

func TestSql(t *testing.T) {

	assert := assert.New(t)

	sql := Sql{
		Name: "Penlook",
	}.Connect()

	sql.DropTableIfExists(&User{})
	sql.CreateTable(&User{})

	for i := 0; i < 100; i++ {
		sql.Create(User{
			Username: "loint",
			Email:    "loint@penlook.com",
			Password: fmt.Sprintf("%x", md5.Sum([]byte("12345"))),
		})
	}

	var users []User
	var count int

	sql.Find(&users).Count(&count)
	assert.Equal(100, count)

	sql.DropTableIfExists(&User{})
}
