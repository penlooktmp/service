// Copyright 2014 Penlook Development Team. All rights reserved.
// Use of this source code is governed by
// license that can be found in the LICENSE file.
// Author : Loi Nguyen <loint@penlook.com>

package doc

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

type Person struct {
	Name  string
	Phone string
}

func TestDoc(t *testing.T) {
	assert := assert.New(t)

	doc := Doc{
		Name: "Penlook",
		Server: []string{
			"localhost:27017",
		},
	}

	// Establish new connection
	doc = doc.Connect()
	penlook := 	doc.Database("penlook")
	person  := 	penlook.Collection("person")

	// Insert some documents
	err1 	:=  person.Insert(
					Person{"Ale", "+55 53 8116 9639"},
					Person{"Cla", "+55 53 8402 8510"},
			    )
	assert.Nil(err1)

	// Retrieve result
	count, err2 := person.Find(Json{}).Count()
	assert.Nil(err2)
	assert.Equal(2, count)

	// Clean up
	_, err3 := person.RemoveAll(Json{})
	assert.Nil(err3)

	// Close connection
	doc.Close()
}
