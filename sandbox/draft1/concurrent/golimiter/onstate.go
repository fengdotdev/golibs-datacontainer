package golimiter



func (l *Golimiter) ChangeLimit(newLimit int) {
	l.limit = make(chan struct{}, newLimit) // Change the limit for concurrent jobs
}
func (l *Golimiter) CurrentLimit() int {
	return len(l.limit) // Return the current limit
}

func (l *Golimiter) IsFull() bool {
	return len(l.limit) == cap(l.limit) // Check if the limiter is full
}
func (l *Golimiter) IsEmpty() bool {
	return len(l.limit) == 0 // Check if the limiter is empty
}
func (l *Golimiter) Capacity() int {
	return cap(l.limit) // Return the capacity of the limiter
}
func (l *Golimiter) CurrentUsage() int {
	return len(l.limit) // Return the current usage of the limiter
}
