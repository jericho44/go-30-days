
package main

import "fmt"

type User struct {
	Name string
	Age  int
}

func (u User) Introduce() {
	fmt.Printf("Hi, saya %s (%d)\n", u.Name, u.Age)
}

func main() {
	u := User{"Rico", 27}
	u.Introduce()
}
