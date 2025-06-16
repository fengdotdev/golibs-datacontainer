package gocounter

func (c *GOCounter) Increment() {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.counter++
}
func (c *GOCounter) Decrement() {
	c.mu.Lock()
	defer c.mu.Unlock()
	if c.counter > 0 {
		c.counter--
	}
}

func (c *GOCounter) Value() uint64 {
	c.mu.Lock()
	defer c.mu.Unlock()
	return c.counter
}

func (c *GOCounter) EqualTo(value uint64) bool {
	c.mu.Lock()
	defer c.mu.Unlock()
	return c.counter == value
}

func (c *GOCounter) GreaterOrEqualTo(value uint64) bool {
	c.mu.Lock()
	defer c.mu.Unlock()
	return c.counter >= value
}

func (c *GOCounter) LessOrEqualTo(value uint64) bool {
	c.mu.Lock()
	defer c.mu.Unlock()
	return c.counter <= value
}

func (c *GOCounter) Reset() {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.counter = 0
}

func (c *GOCounter) Set(value uint64) {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.counter = value
}

func (c *GOCounter) On(condition func(value uint64) bool, action func()) {
	c.mu.Lock()
	defer c.mu.Unlock()
	if condition(c.counter) {
		action()
	}
}

func (c *GOCounter) OnEqualTo(value uint64, action func()) {
	c.mu.Lock()
	defer c.mu.Unlock()
	if c.counter == value {
		action()
	}
}

func (c *GOCounter) OnGreaterOrEqualTo(value uint64, action func()) {
	c.mu.Lock()
	defer c.mu.Unlock()
	if c.counter >= value {
		action()
	}
}
func (c *GOCounter) DoneOnGreaterOrEqualTo(value uint64) chan struct{} {
	done := make(chan struct{})
	c.mu.Lock()
	defer c.mu.Unlock()
	if c.counter >= value {
		close(done)
	}

	return done
}
