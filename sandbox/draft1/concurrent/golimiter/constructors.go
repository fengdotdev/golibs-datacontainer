package golimiter


func New(limit int) *Golimiter {
	return &Golimiter{
		limit: make(chan struct{}, limit), // Buffered channel to limit concurrent jobs
	}
}
