// Copyright 2014 Penlook Development Team. All rights reserved.
// Use of this source code is governed by
// license that can be found in the LICENSE file.
// Author : Loi Nguyen <loint@penlook.com>

package redis

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestRedis(t *testing.T) {
	assert := assert.New(t)

	redis := Redis{
		Name:   "Penlook",
		Server: "localhost:6379",
	}.Connect()

	redis.Do("SET", "key", "value")

	result, err := String(redis.Do("GET", "key"))
	assert.Nil(err)
	assert.Equal(result, "valuex")
	redis.Close()
}
