package main

import (
	"fmt"
	"sync"
	"time"
)

func work(id int, wg *sync.WaitGroup) {
	defer wg.Done()
	time.Sleep(150 * time.Millisecond)
	fmt.Println("worker", id, "done")
}

func main() {
	var wg sync.WaitGroup
	for i := 1; i <= 5; i++ {
		wg.Add(1)
		go work(i, &wg)
	}
	wg.Wait()
	fmt.Println("all done")
}
