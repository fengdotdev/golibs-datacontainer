package safecontainer

var _ SafeContainerSetters[any] = (*GoContainer[any])(nil)

// Operate implements SafeContainerSetters.
func (g *GoContainer[T]) Operate(func(T) (T, error)) error {
	panic("unimplemented")
}

// OperatePtr implements SafeContainerSetters.
func (g *GoContainer[T]) OperatePtr(func(*T)) error {
	panic("unimplemented")
}

// Set implements SafeContainerSetters.
func (g *GoContainer[T]) Set(T) {
	panic("unimplemented")
}
