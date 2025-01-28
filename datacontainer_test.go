package golibsdatacontainer_test

import (
	"bytes"
	"testing"

	datacontainer "github.com/fengdotdev/golibs-datacontainer"
	"github.com/fengdotdev/golibs-testing/assert"
)

func TestString(t *testing.T) {
	s := "Hello"
	container := datacontainer.NewDataContainer(datacontainer.STRING, s)
	assert.EqualWithMessage(t, container.Kind(), datacontainer.STRING, "container.Kind() should be STRING")
	assert.EqualWithMessage(t, container.Get(), s, "container.Get() should be s")

	var sprime string = container.Get().(string)
	assert.EqualWithMessage(t, sprime, s, "sprime should be s")
}

func TestBinary(t *testing.T) {
	b := []byte{1, 2, 3}
	container := datacontainer.NewDataContainer(datacontainer.BINARY, b)
	assert.EqualWithMessage(t, container.Kind(), datacontainer.BINARY, "container.Kind() should be BINARY")
	assert.TrueWithMessage(t, bytes.Equal(container.Get().([]byte), b), "container.Get() should be b")

	var bprime []byte = container.Get().([]byte)
	assert.TrueWithMessage(t, bytes.Equal(b, bprime), "bprime should be b")

}
