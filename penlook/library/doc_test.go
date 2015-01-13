package library

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestDoc(t *testing.T) {
	assert := assert.New(t)

	var a string = "Hello"
	var b string = "Hello"

	assert.Equal(a, b)
}
