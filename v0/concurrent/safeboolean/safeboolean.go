package safebolean

// SafeBoolean is an interface that defines methods for a thread-safe boolean value.
type SafeBoolean interface {
	IsTrue() bool
	IsFalse() bool
	String() string
	Get() bool
	Value() bool
	Set(value bool)
	Toggle()
	Operate(op func(bool) bool) // Allows to operate on the boolean value with a function in a thread-safe manner 
	StreamBoolean
}

type StreamBoolean interface {
	// Subscribe allows a channel to receive updates when the boolean value changes.
	Subscribe(boolChan chan bool)
	// Stream returns a channel that streams the boolean value changes.
	// underhood, make a channel and subscribe to it
	Stream() chan bool
	// Await waits for the next boolean value change and returns it.
	Await() bool
	// Clean cleans up the subscribers and closes the channels.
	Clean()
	// GetSubscribers returns a list of channels that are subscribed to the boolean value changes.
	// Warning: channels are many to one.
	GetSubscribers() []chan bool
}
