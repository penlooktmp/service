package penlook

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestApi(t *testing.T) {
	assert := assert.New(t)

	var a string = "Hello"
	var b string = "Hellox"

	assert.Equal(a, b)
}
