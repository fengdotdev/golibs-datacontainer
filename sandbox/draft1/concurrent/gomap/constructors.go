package gomap

import "sync/atomic"

const (
	limit = 1000000 // Default limit for the number of jobs
)

func NewGoMap[K comparable, V any]() *GoMap[K, V] {
	m := &GoMap[K, V]{
		mapdata: make(map[K]V),
		jobs:    make(chan func(), 100), // Buffered channel for jobs
		counter: atomic.Uint64{},
		limit:   limit, // Set the default limit for the number of jobs
	}
	m.Start() // Start the goroutine to process jobs

	return m
}

func NewGoMapWithBuffer[K comparable, V any](buffer int) *GoMap[K, V] {
	m := &GoMap[K, V]{
		mapdata: make(map[K]V),
		jobs:    make(chan func(), buffer), // Buffered channel for jobs
		counter: atomic.Uint64{},
		limit:   limit, // Set the default limit for the number of jobs
	}
	m.Start() // Start the goroutine to process jobs

	return m
}

func NewGoMapFromMap[K comparable, V any](initialMap map[K]V) *GoMap[K, V] {
	m := &GoMap[K, V]{
		mapdata: make(map[K]V),
		jobs:    make(chan func(), 100), // Buffered channel for jobs
		counter: atomic.Uint64{},
		limit:   limit, // Set the default limit for the number of jobs

	}
	m.mapdata = initialMap // Initialize with the provided map
	m.Start()              // Start the goroutine to process jobs

	return m
}
