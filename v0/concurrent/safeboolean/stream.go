package safebolean

func (sb *GoBoolean) Stream() chan bool {

	boolChan := make(chan bool)
	sb.Subscribe(boolChan)

	return boolChan
}

func (sb *GoBoolean) Await() bool {
	boolChan := sb.Stream()
	defer close(boolChan)

	select {
	case value := <-boolChan:
		return value
	}
}

func (sb *GoBoolean) GetSubscribers() []chan bool {
	sb.mu.Lock()
	defer sb.mu.Unlock()

	return sb.subscribers
}

func NotifyAll(value bool, subs []chan bool) {

	if subs == nil {
		return
	}

	for _, subscriber := range subs {
		go func(value bool) {
			subscriber <- value
		}(value)
	}
}
