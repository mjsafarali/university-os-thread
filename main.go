package main

import (
	"fmt"
	"sync"
)

var wg sync.WaitGroup //used to wait the app, until all thread done

func zeroTo100_I() {
	const threadName = 1
	defer wg.Done()

	for i := 1; i <= 100; i++ {
		fmt.Printf("[Thread %d]: #%d\n", threadName, i)
	}
}

func zeroTo100_II() {
	const threadName = 2
	defer wg.Done()

	for i := 1; i <= 100; i++ {
		fmt.Printf("[Thread %d]: #%d\n", threadName, i)
	}
}

func zeroTo100_III() {
	const threadName = 3
	defer wg.Done()

	for i := 1; i <= 100; i++ {
		fmt.Printf("[Thread %d]: #%d\n", threadName, i)
	}
}

func main() {
	wg.Add(3)

	go zeroTo100_I()
	go zeroTo100_II()
	go zeroTo100_III()

	wg.Wait()
}
