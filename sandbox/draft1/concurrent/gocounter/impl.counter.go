package gocounter

import "github.com/fengdotdev/golibs-datacontainer/sandbox/draft1/interfaces"

var _ interfaces.Counter[int] = (*GOCounter[int])(nil)

func (c *GOCounter[T]) Increment() {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.counter++
}
func (c *GOCounter[T]) Decrement() {
	c.mu.Lock()
	defer c.mu.Unlock()
	if c.counter > c.initialValue {
		c.counter--
	}
}

func (c *GOCounter[T]) Value() T {
	c.mu.Lock()
	defer c.mu.Unlock()
	return c.counter
}

func (c *GOCounter[T]) Reset() {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.counter = c.initialValue
}

func (c *GOCounter[T]) Set(value T) {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.counter = value
}
