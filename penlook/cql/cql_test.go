// Copyright 2014 Penlook Development Team. All rights reserved.
// Use of this source code is governed by
// license that can be found in the LICENSE file.
// Author : Loi Nguyen <loint@penlook.com>

package cql

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestCql(t *testing.T) {
	assert.New(t)
	cql := Cql {
		Name:  "Penlook"
		Server: "Abc"
	}
}
