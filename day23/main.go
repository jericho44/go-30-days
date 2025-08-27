
package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/ping", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "pong")
	})
	http.HandleFunc("/hello/", func(w http.ResponseWriter, r *http.Request) {
		name := r.URL.Path[len("/hello/"):]
		fmt.Fprintf(w, "Hello, %s\n", name)
	})
	http.ListenAndServe(":8080", nil)
}
