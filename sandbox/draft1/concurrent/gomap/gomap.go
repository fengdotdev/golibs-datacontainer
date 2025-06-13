package gomap

import "sync/atomic"

type GoMap[K comparable, V any] struct {
	mapdata map[K]V
	jobs    chan func()
	counter   atomic.Uint64
	limit   int64
}
