package main

import (
	"bytes"
	"testing"

	"gotest.tools/assert"
)

func TestGreet(t *testing.T) {

	buf := bytes.Buffer{}
	Greet(&buf, "Lorenzo")

	assert.Equal(t, "Hello Lorenzo", buf.String())

}
