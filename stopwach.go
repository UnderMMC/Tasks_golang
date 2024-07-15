package main

import (
	"fmt"
	"time"
)

func main() {
	ticker := time.NewTicker(time.Second)
	done := make(chan bool)

	go func() {
		for {
			select {
			case <-done:
				return
			case t := <-ticker.C:
				fmt.Println("Time at: ", t)
			}
		}
	}()
	time.Sleep(16000 * time.Second)
	ticker.Stop()
	done <- true
	fmt.Println("Ticker stopped")
}
