package safecontainer

func NotifyAll[T any](value T, subs []chan T) {
	if subs == nil {
		return
	}

	for _, subscriber := range subs {
		go func(value T) {
			subscriber <- value
		}(value)
	}
}
