package main

import (
	"context"
	"fmt"
)

func main() {

	ch := make(chan bool)
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	fmt.Println("Do you want to continue? (yes or no):")
	var answer string
	fmt.Scan(&answer)

	switch answer {
	case "yes":
		go Speach(ctx, ch)
		<-ch
	case "no":
		fmt.Println("Goroutine stopped :(")
		ctx.Done()
	}
}

func Speach(ctx context.Context, ch chan bool) {
	fmt.Println("Goroutine started :)")
	ch <- true
}
