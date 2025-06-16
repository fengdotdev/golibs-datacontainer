package main

import (
	"context"
	"fmt"
	"os"
	"sync"
	"time"

	"github.com/fengdotdev/golibs-datacontainer/sandbox/draft1/concurrent/golimiter"
	"github.com/fengdotdev/golibs-datacontainer/sandbox/draft1/concurrent/gomap"
)

func main() {
	GoMapAndLimiter()
}

func GoMapAndLimiter() {

	mmyMap := gomap.NewGoMap[string, int]()
	defer mmyMap.Save()
	mmyMap.SetSaver(
		func(data []byte) error {
			// save on disk

			path := "data.gob"
			file, err := os.Create(path)
			if err != nil {
				return fmt.Errorf("failed to create file: %w", err)
			}
			defer file.Close()
			_, err = file.Write(data)
			if err != nil {
				return fmt.Errorf("failed to write data to file: %w", err)
			}
			fmt.Println("Data saved to", path)
			return nil
		},
	)
	var wg sync.WaitGroup
	ltr := golimiter.New(100) // Create a new limiter with a limit of 100 concurrent jobs
	running := bool(true)
	go func() {

		for running {
			time.Sleep(5 * time.Second) //
			fmt.Println("Current limit:", ltr.CurrentLimit())
			fmt.Println("Is limiter full?", ltr.IsFull())
		}

	}()

	go func() {
		for running {
			time.Sleep(3 * time.Minute) //
			err := mmyMap.Save()
			if err != nil {
				panic(fmt.Sprintf("Error saving map: %v", err))
			}
		}
	}()

	for i := 0; i < 1000; i++ {

		wg.Add(1)
		ltr.Acquire() // Acquire a slot in the limiter
		go func(i int) {
			defer ltr.Release() // Ensure the slot is released when done

			defer wg.Done()
			err := <-mmyMap.Set(context.Background(), fmt.Sprintf("key%d", i), i)
			if err != nil {
				fmt.Println("Error overwriting value:", err)
			} else {
				//fmt.Println("Goroutine", i, "overwrote value for key:", fmt.Sprintf("key%d", i))
			}

		}(i)
	}

	err := <-mmyMap.Set(context.Background(), "key1", 42)

	if err != nil {
		fmt.Println("Error setting value:", err)
	}

	for i := 0; i < 1000; i++ {

		wg.Add(1)
		ltr.Acquire() // Acquire a slot in the limiter
		go func(i int) {
			defer ltr.Release() // Ensure the slot is released when done
			defer wg.Done()

			valueChan, errChan := mmyMap.Get(context.Background(), "key1")
			if err := <-errChan; err != nil {
				fmt.Println("Error getting value:", err)
			} else {
				<-valueChan
			}

		}(i)
	}

	wg.Wait()
	running = false // Stop the goroutine that prints the current limit
	fmt.Println("All goroutines completed.")
	err = mmyMap.Save()
	if err != nil {
		panic(fmt.Sprintf("Error saving map: %v", err))
	}
	fmt.Println("Map saved successfully.")

	lenchan, errchan := mmyMap.Size(context.Background())
	if err := <-errchan; err != nil {
		fmt.Println("Error getting size:", err)
		return
	}

	<-lenchan
}
