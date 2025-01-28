package golibsdatacontainer_test

import (
	"testing"

	container "github.com/fengdotdev/golibs-datacontainer"
	"github.com/fengdotdev/golibs-testing/assert"
)

func TestStringInterface(t *testing.T) {
	var s container.DataContainer = container.NewStringData("Hello")
	assert.TrueWithMessage(t, s != nil, "b should not be nil")
}
