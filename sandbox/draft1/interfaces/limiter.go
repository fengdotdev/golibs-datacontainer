package interfaces

type Limiter interface {
	LimiterAcquirer
	LimiterState
}

type LimiterState interface {
}

type LimiterAcquirer interface {
	Acquire()
	Release()
}
