package main

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"os"
	"testing"
)

var s3 = S3{}

func TestS3Config(t *testing.T) {
	assert := assert.New(t)
	s3.GetConfig()
	assert.NotNil(s3.ID)
	assert.NotNil(s3.Secret)

	assert.Equal(20, len(s3.ID))
	assert.Equal(40, len(s3.Secret))
	fmt.Println()
}

func TestS3Create(t *testing.T) {
	assert := assert.New(t)
	s3.Create()
	assert.NotNil(s3.s3cli)
}

func TestS3CheckExistFile(t *testing.T) {
	filename := "test"
	assert := assert.New(t)
	os.Mkdir(filename, 0644)
	fi, fd := s3.CheckExistFile(filename)
	os.Remove(filename)
	assert.NotNil(fi)
	assert.NotNil(fd)
}
