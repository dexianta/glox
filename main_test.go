package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestRun(t *testing.T) {
	err := run("1 = 1")
	assert.Nil(t, err)
}