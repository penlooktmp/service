package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
	//"fmt"
	"net"
)

func TestApi(t *testing.T) {
	assert := assert.New(t)

	_, err := net.Dial("tcp", ":80")
	assert.Nil(err)
}
