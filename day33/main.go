
package main

import "fmt"

func main() {
	// Unbuffered
	ch := make(chan int)

	go func() {
		for i := 1; i <= 3; i++ { ch <- i }
		close(ch)
	}()

	for v := range ch {
		fmt.Println("recv:", v)
	}

	// Buffered
	b := make(chan string, 2)
	b <- "A"
	b <- "B"
	// b <- "C" // ini akan blok (buffer penuh) kalau tidak ada receiver
	fmt.Println(<-b, <-b)
}
