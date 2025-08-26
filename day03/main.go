package main

import "fmt"

func main() {
	var a int = 10
	var b float64 = 3.5
	c := float64(a) + b
	fmt.Printf("c=%.2f, tipe=%T\n", c, c)
}
