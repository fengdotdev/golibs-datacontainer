package gocounter

import "sync"

func NewGOCounter[T Number]() *GOCounter[T] {
	return &GOCounter[T]{
		counter:      0, // Assuming the zero value of T is 0
		initialValue: 0, // Assuming the zero value of T is 0
		mu:           sync.Mutex{},
	}
}

func NewGOCounterWithValue[T Number](initialValue T) *GOCounter[T] {
	return &GOCounter[T]{
		counter:      initialValue,
		initialValue: initialValue,
		mu:           sync.Mutex{},
	}
}
