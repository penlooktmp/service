// Copyright 2014 Penlook Development Team. All rights reserved.
// Use of this source code is governed by
// license that can be found in the LICENSE file.
// Author : Loi Nguyen <loint@penlook.com>

package rabit

import (
	"github.com/streadway/amqp"
	"github.com/stretchr/testify/assert"
	"log"
	"testing"
)

func TestRabit(t *testing.T) {

	assert.New(t)

	rabit := Rabit{
		Name:   "Penlook",
		Server: "amqp://guest:guest@localhost:5672/",
	}

	rabit.Connect()

}
