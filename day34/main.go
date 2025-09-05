
package main

import (
	"context"
	"fmt"
	"time"
)

func slow(ctx context.Context, out chan<- int) {
	defer close(out)
	for i := 1; i <= 5; i++ {
		select {
		case <-ctx.Done():
			return
		case <-time.After(300 * time.Millisecond):
			out <- i
		}
	}
}

func main() {
	ctx, cancel := context.WithTimeout(context.Background(), 800*time.Millisecond)
	defer cancel()

	ch := make(chan int)
	go slow(ctx, ch)

	for {
		select {
		case v, ok := <-ch:
			if !ok {
				fmt.Println("done")
				return
			}
			fmt.Println("got:", v)
		case <-ctx.Done():
			fmt.Println("timeout/cancel:", ctx.Err())
			return
		}
	}
}
