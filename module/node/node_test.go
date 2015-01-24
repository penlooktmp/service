// Copyright 2014 Penlook Development Team. All rights reserved.
// Use of this source code is governed by
// license that can be found in the LICENSE file.
// Author : Loi Nguyen <loint@penlook.com>

package node

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNode(t *testing.T) {
	assert := assert.New(t)
	node := Node{
		Name: "Test Node",
	}

	result, err := node.Abc(6, 2)
	assert.Nil(err)
	assert.Equal(3, result)
}
