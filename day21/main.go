
package main

import (
	"fmt"
	"time"
)

func download(part int, out chan<- string) {
	time.Sleep(time.Duration(200+part*100) * time.Millisecond)
	out <- fmt.Sprintf("part %d selesai", part)
}

func main() {
	out := make(chan string)
	parts := 10
	for i := 1; i <= parts; i++ {
		go download(i, out)
	}
	for i := 0; i < parts; i++ {
		fmt.Println(<-out)
	}
	fmt.Println("Semua selesai")
}
