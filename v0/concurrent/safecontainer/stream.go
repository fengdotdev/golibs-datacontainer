package safecontainer

var _ StreamContainer[any] = (*GoContainer[any])(nil)

// Await implements StreamContainer.
func (g *GoContainer[T]) Await() T {
	panic("unimplemented")
}

// Clean implements StreamContainer.
func (g *GoContainer[T]) Clean() {
	panic("unimplemented")
}

// GetSubscribers implements StreamContainer.
func (g *GoContainer[T]) GetSubscribers() []chan T {
	panic("unimplemented")
}

// Stream implements StreamContainer.
func (g *GoContainer[T]) Stream() chan T {
	panic("unimplemented")
}

// Subscribe implements StreamContainer.
func (g *GoContainer[T]) Subscribe(boolChan chan T) {
	panic("unimplemented")
}
