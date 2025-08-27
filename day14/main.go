
package main

import "fmt"

type Contact struct {
	Name  string
	Phone string
}

var contacts []Contact

func add(name, phone string) { contacts = append(contacts, Contact{name, phone}) }
func list() {
	for i, c := range contacts {
		fmt.Printf("%d. %s - %s\n", i+1, c.Name, c.Phone)
	}
}
func removeByName(name string) {
	out := contacts[:0]
	for _, c := range contacts {
		if c.Name != name {
			out = append(out, c)
		}
	}
	contacts = out
}

func main() {
	add("Ali", "0812")
	add("Budi", "0813")
	list()
	removeByName("Ali")
	fmt.Println("--- setelah hapus ---")
	list()
}
