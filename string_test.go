package hexago

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestString_Parse(t *testing.T) {
	str := String{}
	str.Parse("Hello")
	assert.Equal(t, "Hello", str.Value)
}
