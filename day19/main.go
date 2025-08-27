
package main

import (
	"fmt"
	"time"
)

func worker(name string) {
	for i := 1; i <= 3; i++ {
		fmt.Println(name, i)
		time.Sleep(200 * time.Millisecond)
	}
}

func main() {
	go worker("A")
	go worker("B")
	time.Sleep(1 * time.Second)
	fmt.Println("Done")
}
