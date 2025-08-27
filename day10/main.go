
package main

import "fmt"

func main() {
	arr := [3]int{1, 2, 3}
	sl := []int{10, 20}
	sl = append(sl, 30)
	fmt.Println(arr, sl, len(sl), cap(sl))
}
