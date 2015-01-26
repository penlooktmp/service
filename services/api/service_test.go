package main

import (
	//"bufio"
	//"fmt"
	"github.com/stretchr/testify/assert"
	"net"
	"testing"
)

// TCP Client
func TestTcpClient(t *testing.T) {
	assert := assert.New(t)

	conn, err := net.Dial("tcp", ":80")
	assert.Nil(err)

	conn.Write([]byte("Hello Server"))
	conn.Close()

	//fmt.Println(status)
}
