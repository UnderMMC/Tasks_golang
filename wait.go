package main

import (
	"fmt"
	"sync"
)

func main() {
	var wg sync.WaitGroup
	wg.Add(2)
	done := func(id int) {
		defer wg.Done()
		fmt.Printf("Goroutine %d started:\n", id)
		for i := 0; i < 10; i++ {
			fmt.Printf("%d ", i)
		}
		fmt.Printf("\n")
		fmt.Printf("Goroutine %d finished;\n", id)
	}
	go done(1)
	go done(2)

	wg.Wait()
	fmt.Println("All Goroutines are finished!")
}
