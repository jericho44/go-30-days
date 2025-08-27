
package main

import (
	"encoding/json"
	"fmt"
)

type Todo struct {
	ID   int    `json:"id"`
	Text string `json:"text"`
	Done bool   `json:"done"`
}

func main() {
	// encode
	t := Todo{1, "Belajar Go", false}
	b, _ := json.Marshal(t)
	fmt.Println(string(b))

	// decode
	var t2 Todo
	_ = json.Unmarshal([]byte(`{"id":2,"text":"Latihan JSON","done":true}`), &t2)
	fmt.Printf("%+v\n", t2)
}
