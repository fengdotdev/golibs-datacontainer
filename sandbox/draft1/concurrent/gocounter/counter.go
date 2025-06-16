package gocounter

import (
	"sync"
)

type GOCounter[T Number] struct {
	counter      T
	initialValue T
	mu           sync.Mutex
}
