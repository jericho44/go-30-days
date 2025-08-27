
package main

import "fmt"

func main() {
	var a, b float64
	var op string
	fmt.Print("Masukkan: <angka> <op> <angka> (cth: 3 + 2): ")
	fmt.Scan(&a, &op, &b)

	switch op {
	case "+":
		fmt.Println(a + b)
	case "-":
		fmt.Println(a - b)
	case "*":
		fmt.Println(a * b)
	case "/":
		if b == 0 {
			fmt.Println("Error: bagi 0")
			return
		}
		fmt.Println(a / b)
	default:
		fmt.Println("Operator tidak dikenal")
	}
}
