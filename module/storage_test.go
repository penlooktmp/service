// Copyright 2014 Penlook Development Team. All rights reserved.
// Use of this source code is governed by
// license that can be found in the LICENSE file.
// Author : Tim Nguyen <tinntt@penlook.com>
package penlook

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestStorage(t *testing.T) {
	assert := assert.New(t)

	var a string = "Hello"
	var b string = "Hello"

	assert.Equal(a, b)
}
