package golimiter

func (l *Golimiter) Acquire() {
	l.limit <- struct{}{} // Acquire a slot in the limiter
}
func (l *Golimiter) Release() {
	<-l.limit // Release a slot in the limiter
}
