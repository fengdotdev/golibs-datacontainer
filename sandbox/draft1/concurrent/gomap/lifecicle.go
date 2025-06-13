package gomap

func (gm *GoMap[K, V]) Start() {

	go func() {
		for job := range gm.jobs {
			job() // Execute the job
		}
	}()

}

func (gm *GoMap[K, V]) Stop() {
	close(gm.jobs) // Close the jobs channel to stop the goroutine
}

