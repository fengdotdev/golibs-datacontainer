package gomap

import (
	"bytes"
	"context"
	"encoding/gob"
)

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

func (gm *GoMap[K, V]) SetSaver(saver func(data []byte) error) {
	gm.onsaver = saver // Set the saver function
}

func (gm *GoMap[K, V]) Save() error {
	if gm.onsaver == nil {
		return nil // No saver function defined
	}

	data, err := gm.ToGOB() // Convert the map to GOB format
	if err != nil {
		return err // Return error if conversion fails
	}

	if data == nil {
		return nil // No data to save
	}

	return gm.onsaver(data) // Call the saver function with the GOB data
}

func (gm *GoMap[K, V]) ChangeLimit(newLimit int) {
	gm.limiter = make(chan struct{}, newLimit) // Change the limit for concurrent jobs
}

func (gm *GoMap[K, V]) ToGOB() ([]byte, error) {
	if gm.onsaver == nil {
		return nil, nil // No saver function defined
	}

	var buffer bytes.Buffer

	encoder := gob.NewEncoder(&buffer) // Create a new GOB encoder

	datachan, errchan := gm.Clone(context.Background()) // Get the map data to encode

	err := <-errchan // Wait for the cloning to complete
	if err != nil {
		return nil, err // Return error if cloning fails
	}

	data := <-datachan // Receive the cloned map data
	if data == nil {
		return nil, nil
	} // Return nil if no data is available
	err = encoder.Encode(data) // Encode the map data into GOB format

	if err != nil {
		return nil, err // Return error if conversion fails
	}

	binaryData := buffer.Bytes()

	return binaryData, nil // Return the GOB data
}
