package main

import (
	"context"
	"fmt"
	"sync"

	"github.com/fengdotdev/golibs-datacontainer/sandbox/draft1/concurrent/gomap"
)

func main() {

	mmyMap := gomap.NewGoMap[string, int]()

	var wg sync.WaitGroup
	for i := 0; i < 100000000; i++ {

		wg.Add(1)

		go func(i int) {
			defer wg.Done()
			err := <-mmyMap.Set(context.Background(), fmt.Sprintf("key%d", i), i)
			if err != nil {
				fmt.Println("Error overwriting value:", err)
			} else {
				fmt.Println("Goroutine", i, "overwrote value for key:", fmt.Sprintf("key%d", i))
			}
		}(i)
	}

	err := <-mmyMap.Set(context.Background(), "key1", 42)

	if err != nil {
		fmt.Println("Error setting value:", err)
	}

	for i := 0; i < 100000000; i++ {
		wg.Add(1)
		go func(i int) {
			defer wg.Done()

			valueChan, errChan := mmyMap.Get(context.Background(), "key1")
			if err := <-errChan; err != nil {
				fmt.Println("Error getting value:", err)
			} else {
				value := <-valueChan
				fmt.Println("Goroutine", i, "retrieved value:", value)
			}

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
	fmt.Println("Length of map:", len)
}
