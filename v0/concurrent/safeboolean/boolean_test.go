package safebolean_test

import (
	"fmt"
	"sync"
	"testing"

	safebolean "github.com/fengdotdev/golibs-datacontainer/v0/concurrent/safeboolean"
	"github.com/fengdotdev/golibs-testing/assert"
)

func TestSafeBolean(t *testing.T) {

	t.Run("New", func(t *testing.T) {

		// Create a new SafeBoolean instance
		sb := safebolean.New()

		assert.FalseWithMessage(t, sb.Get(), "Expected initial value to be false")
		assert.FalseWithMessage(t, sb.IsTrue(), "Expected IsTrue to return false initially")
		assert.TrueWithMessage(t, sb.IsFalse(), "Expected IsFalse to return true initially")

		// Set value to true
		sb.Set(true)
		assert.TrueWithMessage(t, sb.Get(), "Expected Get to return true after setting to true")
		assert.TrueWithMessage(t, sb.IsTrue(), "Expected value to be true after setting to true")
		assert.FalseWithMessage(t, sb.IsFalse(), "Expected IsFalse to return false after setting to true")

		// Toggle value
		sb.Toggle()
		assert.FalseWithMessage(t, sb.Get(), "Expected Get to return false after toggling from true")

		// Clean up
		sb.Clean()

	})

	t.Run("concurrent", func(t *testing.T) {

		sb := safebolean.True()
		assert.TrueWithMessage(t, sb.Get(), "Expected initial value to be true")
		var wg sync.WaitGroup
		for i := 0; i < 100; i++ {
			wg.Add(1)
			go func() {
				sb.Toggle()
				t.Log(fmt.Sprintf("Toggled value: %v", sb.Get()))
				wg.Done()
			}()
		}
		for i := 0; i < 100; i++ {
			wg.Add(1)
			go func() {
				defer wg.Done()
				sb.Operate(func(value bool) bool {
					result := !value
					t.Log(fmt.Sprintf("Toggled value: %v", result))
					return result
				})
			}()
		}

		wg.Wait()

	})
}
