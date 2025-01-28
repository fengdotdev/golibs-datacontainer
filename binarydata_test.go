package golibsdatacontainer_test

import (
	"testing"

	container "github.com/fengdotdev/golibs-datacontainer"
	"github.com/fengdotdev/golibs-testing/assert"
)

func TestBinaryInterface(t *testing.T) {
	var b container.DataContainer = container.NewBinaryData([]byte{1, 2, 3})
	assert.TrueWithMessage(t, b != nil, "b should not be nil")
}
