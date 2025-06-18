package main

import (
	"context"
	"fmt"
	"sync"

	"github.com/fengdotdev/golibs-datacontainer/sandbox/draft1/concurrent/gomap"
)

func main() {
	defer recoverPanic()
	GoMapAndLimiter()
}

func GoMapAndLimiter() {
	defer recoverPanic()

	mmyMap := gomap.NewGoMapWithBuffer[string, int](20)

	var wg sync.WaitGroup

	for i := 0; i < 1000000; i++ {

		wg.Add(1)
		go func(i int) {
			defer recoverPanic()
			defer wg.Done()
			err := <-mmyMap.Set(context.Background(), fmt.Sprintf("key%d", i), i)
			if err != nil {
				fmt.Println("Error overwriting value:", err)
			} else {
				//fmt.Println("Goroutine", i, "overwrote value for key:", fmt.Sprintf("key%d", i))
			}
			fmt.Println("Goroutine", i, "set value for key:", fmt.Sprintf("key%d", i))

		}(i)
	}

	err := <-mmyMap.Set(context.Background(), "key1", 42)

	if err != nil {
		fmt.Println("Error setting value:", err)
	}

	for i := 0; i < 10000000; i++ {

		wg.Add(1)
		go func(i int) {
			defer recoverPanic()
			defer wg.Done()

			valueChan, errChan := mmyMap.Get(context.Background(), "key1")
			if err := <-errChan; err != nil {
				fmt.Println("Error getting value:", err)
			} else {
				<-valueChan
			}
			fmt.Println("Goroutine", i, "retrieved")

		}(i)
	}

	wg.Wait()
	fmt.Println("All goroutines completed.")

	lenchan, errchan := mmyMap.Size(context.Background())
	if err := <-errchan; err != nil {
		fmt.Println("Error getting size:", err)
		return
	}

	len := <-lenchan
	fmt.Println("Size of the map:", len)
}

func recoverPanic() {
	if r := recover(); r != nil {
		fmt.Println("Recovered from panic:", r)
	}
}
