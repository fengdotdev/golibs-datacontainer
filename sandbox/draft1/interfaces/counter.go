package interfaces

type Counter[T CounterNumber] interface {
	Increment()
	Decrement()
	Value() T
	Reset()
	Set(value T)
}
type CounterOn[T CounterNumber] interface {
	On(condition func(value T) bool, action func())
	OnEqualTo(value T, action func())
	OnGreaterOrEqualTo(value T, action func())
	OnLessOrEqualTo(value T, action func())
	DoneOnGreaterOrEqualTo(value T) chan struct{}
}

type CounterEquals[T CounterNumber] interface {
	EqualTo(value T) bool
	GreaterOrEqualTo(value T) bool
	LessOrEqualTo(value T) bool
}

type CounterNumber interface {
	int | int8 | int16 | int32 | int64 |
		uint | uint8 | uint16 | uint32 | uint64 |
		float32 | float64
}
