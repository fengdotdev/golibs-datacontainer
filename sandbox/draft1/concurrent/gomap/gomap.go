package gomap

type GoMap[K comparable, V any] struct {
	mapdata map[K]V
	jobs    chan func()
	onsaver func(data []byte) error
	limiter chan struct{}
}
