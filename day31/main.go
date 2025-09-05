package main

import "fmt"

type User struct {
	Name string
	Age  int
}

func renameByValue(u User) { u.Name = "ByValue" }
func renameByPtr(u *User)  { u.Name = "ByPtr" }

func main() {
	u := User{"Budi", 27}
	renameByValue(u)
	fmt.Println("1)", u.Name) // Budi

	renameByPtr(&u)
	fmt.Println("2)", u.Name) // ByPtr

	// Jebakan: alamat var loop
	names := []string{"a", "b", "c"}
	ptrs := []*string{}
	for _, n := range names {
		nn := n // solusi: buat copy per iterasi
		ptrs = append(ptrs, &nn)
	}
	for _, p := range ptrs {
		fmt.Println(*p) // a b c (benar)
	}
}
