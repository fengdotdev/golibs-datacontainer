package safecontainer

type SafeContainer[T any] interface {
	SafeContainerSetters[T]
	SafeContainerGetters[T]
}

type SafeContainerWithStream[T any] interface {
	SafeContainerSetters[T]
	SafeContainerGetters[T]
	StreamContainer[T]
}

type SafeContainerSetters[T any] interface {
	Set(T)
	Operate(func(T) (T, error)) error
	OperatePtr(func(*T)) error
}

type SafeContainerGetters[T any] interface {
	IsInitialized() bool
	Get() T
	Value() T
	IsEmpty() bool
	String() string
}

type StreamContainer[T any] interface {
	// Subscribe allows a channel to receive updates when the boolean value changes.
	Subscribe(boolChan chan T)
	// Stream returns a channel that streams the boolean value changes.
	// underhood, make a channel and subscribe to it
	Stream() chan T
	// Await waits for the next boolean value change and returns it.
	Await() T
	// Clean cleans up the subscribers and closes the channels.
	Clean()
	// GetSubscribers returns a list of channels that are subscribed to the boolean value changes.
	// Warning: channels are many to one.
	GetSubscribers() []chan T
}
