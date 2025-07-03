package safecontainer

var _ SafeContainerGetters[any] = (*GoContainer[any])(nil)

// Get implements SafeContainerGetters.
func (g *GoContainer[T]) Get() T {
	panic("unimplemented")
}

// IsEmpty implements SafeContainerGetters.
func (g *GoContainer[T]) IsEmpty() bool {
	panic("unimplemented")
}

// IsInitialized implements SafeContainerGetters.
func (g *GoContainer[T]) IsInitialized() bool {
	panic("unimplemented")
}

// String implements SafeContainerGetters.
func (g *GoContainer[T]) String() string {
	panic("unimplemented")
}

// Value implements SafeContainerGetters.
func (g *GoContainer[T]) Value() T {
	panic("unimplemented")
}
