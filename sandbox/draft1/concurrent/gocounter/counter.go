package gocounter

import (
	"sync"
)

type GOCounter struct {
	counter uint64
	mu      sync.Mutex
}
