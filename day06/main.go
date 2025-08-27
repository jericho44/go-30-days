
package main

import "fmt"

func main() {
	sum := 0
	for i := 1; i <= 5; i++ {
		sum += i
	}
	fmt.Println("Jumlah 1..5 =", sum)

	i := 0
	for i < 3 {
		fmt.Println("i:", i)
		i++
	}
}
