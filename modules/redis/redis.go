// Copyright 2014 Penlook Development Team. All rights reserved.
// Use of this source code is governed by
// license that can be found in the LICENSE file.
// Author : Loi Nguyen <loint@penlook.com>

package redis

import (
	//"fmt"
	goredis "github.com/penlook/redis/redis"
)

type Redis struct {
	Name       string
	Server     string
	Do         func(cmd string, args ...interface{}) (interface{}, error)
	connection goredis.Conn
}

// Get server
func (redis Redis) GetServer() string {
	return redis.Server
}

// Set connection
func (redis *Redis) SetConnection(connection goredis.Conn, err error) {

	if err != nil {
		panic(err)
	}

	redis.connection = connection

	// Alias function
	redis.Do = redis.connection.Do
}

// Establish new connection
func (redis Redis) Connect() Redis {
	redis.SetConnection(goredis.Dial("tcp", redis.GetServer()))
	return redis
}

// Close connection
func (redis Redis) Close() {
	redis.connection.Close()
}
