
package main

import "fmt"

func divMod(a, b int) (int, int) {
	return a / b, a % b
}

func main() {
	q, r := divMod(10, 3)
	fmt.Println("q=", q, "r=", r)
}
