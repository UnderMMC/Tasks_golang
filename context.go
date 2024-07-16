package main

import (
	"context"
	"fmt"
	"time"
)

var comp int = 1

func main() {
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()

	go task(ctx)

	time.Sleep(10 * time.Second)
}

func task(ctx context.Context) {
	select {
	case <-time.After(3 * time.Second):
		fmt.Println("Task done!")
	case <-ctx.Done():
		fmt.Println("Task failed...")
	}
}
