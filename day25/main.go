
package main

import (
	"encoding/json"
	"net/http"
	"strconv"
	"sync"
)

type Todo struct {
	ID   int    `json:"id"`
	Text string `json:"text"`
	Done bool   `json:"done"`
}

var (
	todos = []Todo{}
	mtx   sync.Mutex
	idSeq = 1
)

func list(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(todos)
}

func create(w http.ResponseWriter, r *http.Request) {
	var t Todo
	json.NewDecoder(r.Body).Decode(&t)
	mtx.Lock()
	t.ID = idSeq
	idSeq++
	todos = append(todos, t)
	mtx.Unlock()
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(t)
}

func toggle(w http.ResponseWriter, r *http.Request) {
	idStr := r.URL.Query().Get("id")
	id, _ := strconv.Atoi(idStr)
	mtx.Lock()
	defer mtx.Unlock()
	for i := range todos {
		if todos[i].ID == id {
			todos[i].Done = !todos[i].Done
			json.NewEncoder(w).Encode(todos[i])
			return
		}
	}
	http.NotFound(w, r)
}

func main() {
	http.HandleFunc("/todos", func(w http.ResponseWriter, r *http.Request) {
		if r.Method == http.MethodGet { list(w, r); return }
		if r.Method == http.MethodPost { create(w, r); return }
		http.NotFound(w, r)
	})
	http.HandleFunc("/toggle", toggle)
	http.ListenAndServe(":8080", nil)
}
