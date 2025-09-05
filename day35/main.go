
package main

import (
	"fmt"
	"sync"
)

func worker(id int, jobs <-chan int, results chan<- int, wg *sync.WaitGroup) {
	defer wg.Done()
	for j := range jobs {
		results <- (j * j) // contoh: kuadrat
	}
}

func main() {
	jobs := make(chan int)
	results := make(chan int, 10)

	var wg sync.WaitGroup
	workerN := 3
	for i := 1; i <= workerN; i++ {
		wg.Add(1)
		go worker(i, jobs, results, &wg)
	}

	// producer
	go func() {
		for i := 1; i <= 10; i++ { jobs <- i }
		close(jobs)
	}()

	// closer untuk results
	go func() {
		wg.Wait()
		close(results)
	}()

	// fan-in consumption
	sum := 0
	for r := range results { sum += r }
	fmt.Println("sum of squares:", sum)
}
