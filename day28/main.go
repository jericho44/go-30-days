
package main

import (
	"log"
	"net/http"
	"time"
)

func logger(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()
		next.ServeHTTP(w, r)
		log.Printf("%s %s %v", r.Method, r.URL.Path, time.Since(start))
	})
}

func hello(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello with middleware"))
}

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/", hello)
	http.ListenAndServe(":8080", logger(mux))
}
