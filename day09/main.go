
package main

import "fmt"

func safeRun() {
	defer func() {
		if err := recover(); err != nil {
			fmt.Println("Recovered:", err)
		}
	}()
	panic("ada yang salah")
}

func main() {
	fmt.Println("Sebelum")
	safeRun()
	fmt.Println("Sesudah")
}
