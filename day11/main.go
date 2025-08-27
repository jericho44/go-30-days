
package main

import "fmt"

func main() {
	a := []int{1,2,3,4}
	b := make([]int, 0, 8)
	b = append(b, a...)
	a[0] = 99
	fmt.Println("a:", a, "b:", b) // b tetap
}
