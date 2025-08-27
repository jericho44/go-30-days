
package main

import "fmt"

func increment(n *int) { *n++ }

func main() {
	x := 10
	increment(&x)
	fmt.Println(x) // 11
}
