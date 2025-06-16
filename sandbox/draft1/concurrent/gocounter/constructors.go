package gocounter

import "sync"

func NewGOCounter() *GOCounter {
	return &GOCounter{
		counter: 0,
		mu:      sync.Mutex{},
	}
}
