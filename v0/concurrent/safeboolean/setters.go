package safebolean

func (sb *GoBoolean) Set(value bool) {
	sb.mu.Lock()
	defer sb.mu.Unlock()
	sb.value = value
	NotifyAll(sb.value, sb.subscribers)
}

func (sb *GoBoolean) Toggle() {
	sb.mu.Lock()
	defer sb.mu.Unlock()
	sb.value = !sb.value
	NotifyAll(sb.value, sb.subscribers)
}

func (sb *GoBoolean) Clean() {
	sb.mu.Lock()
	defer sb.mu.Unlock()

	if sb.subscribers != nil {
		for _, subscriber := range sb.subscribers {
			close(subscriber)
		}
		sb.subscribers = nil
	}
}

// internal methods for GoBoolean

func (sb *GoBoolean) Subscribe(boolChan chan bool) {
	sb.mu.Lock()
	defer sb.mu.Unlock()

	if sb.subscribers == nil {
		sb.subscribers = make([]chan bool, 2)
	}

	sb.subscribers = append(sb.subscribers, boolChan)
}



func (sb *GoBoolean) Operate(op func(bool) bool) {
	sb.mu.Lock()
	defer sb.mu.Unlock()

	sb.value = op(sb.value)
	NotifyAll(sb.value, sb.subscribers)
}
