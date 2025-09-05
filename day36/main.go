
package main

import (
	"fmt"
	"sync"
)

func race() {
	x := 0
	var wg sync.WaitGroup
	for i := 0; i < 1000; i++ {
		wg.Add(1)
		go func() {
			x++ // race!
			wg.Done()
		}()
	}
	wg.Wait()
	fmt.Println("race x:", x)
}

func safe() {
	x := 0
	var wg sync.WaitGroup
	var mu sync.Mutex
	for i := 0; i < 1000; i++ {
		wg.Add(1)
		go func() {
			mu.Lock()
			x++
			mu.Unlock()
			wg.Done()
		}()
	}
	wg.Wait()
	fmt.Println("safe x:", x)
}

func main() {
	race()
	safe()
}
