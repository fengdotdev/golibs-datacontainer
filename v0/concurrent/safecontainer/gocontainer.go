package safecontainer

import "sync"

type GoContainer[T any] struct {
	isInitialized bool
	value         T
	mu            sync.RWMutex
}

type GoContainerWithStream[T any] struct {
	isInitialized bool
	value         T
	mu            sync.RWMutex
	subscribers   []chan T
}
