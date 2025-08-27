
package main

import "fmt"

func main() {
	m := map[string]int{"alice": 90}
	m["bob"] = 87
	if v, ok := m["carol"]; ok {
		fmt.Println(v)
	} else {
		fmt.Println("carol tidak ada")
	}
	for k, v := range m {
		fmt.Println(k, v)
	}
}
