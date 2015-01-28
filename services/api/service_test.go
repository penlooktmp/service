package main

import (
	"github.com/stretchr/testify/assert"
	"net"
	"os/exec"
	"testing"
)

// TCP Client
func TestService(t *testing.T) {
	exec.Command("service", "api", "restart")
	assert := assert.New(t)
	_, err := net.Dial("tcp", ":8080")
	assert.Nil(err)
}

func TestStatus(t *testing.T) {

}
