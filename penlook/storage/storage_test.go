// Copyright 2014 Penlook Development Team. All rights reserved.
// Use of this source code is governed by
// license that can be found in the LICENSE file.
// Author : Tin Nguyen <tinntt@penlook.com>
package storage

import (
	// gstorage "code.google.com/p/google-api-go-client/storage/v1"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestStore(t *testing.T) {
	// assert := assert.New(t)

	//storage := CreateStorage("config.json")
	//storage.ListAllBucket()
	assert := assert.New(t)

	var a string = "Hello"
	var b string = "Hello"

	assert.Equal(a, b)
}
