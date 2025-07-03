package safebolean

import "sync"

type GoBoolean struct {
	value       bool
	mu          sync.Mutex
	subscribers []chan bool
}
