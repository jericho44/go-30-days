package main

import "fmt"

func main() {
	var score int
	fmt.Print("Nilai: ")
	fmt.Scan(&score)

	if score >= 90 {
		fmt.Println("A")
	} else if score >= 75 {
		fmt.Println("B")
	} else {
		fmt.Println("C")
	}

	switch {
	case score%2 == 0:
		fmt.Println("Genap")
	default:
		fmt.Println("Ganjil")
	}
}
