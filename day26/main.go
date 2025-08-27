
package main

import (
	"database/sql"
	"encoding/json"
	"net/http"

	_ "modernc.org/sqlite"
)

type Todo struct {
	ID   int    `json:"id"`
	Text string `json:"text"`
	Done bool   `json:"done"`
}

func main() {
	db, _ := sql.Open("sqlite", "file:todos.db?_pragma=busy_timeout(5000)")
	defer db.Close()
	db.Exec(`CREATE TABLE IF NOT EXISTS todos(id INTEGER PRIMARY KEY, text TEXT, done INTEGER)`)

	http.HandleFunc("/create", func(w http.ResponseWriter, r *http.Request) {
		var t Todo
		json.NewDecoder(r.Body).Decode(&t)
		db.Exec(`INSERT INTO todos(text,done) VALUES(?,0)`, t.Text)
		w.WriteHeader(http.StatusCreated)
	})
	http.HandleFunc("/list", func(w http.ResponseWriter, r *http.Request) {
		rows, _ := db.Query(`SELECT id,text,done FROM todos`)
		defer rows.Close()
		out := []Todo{}
		for rows.Next() {
			var t Todo
			var done int
			rows.Scan(&t.ID, &t.Text, &done)
			t.Done = done == 1
			out = append(out, t)
		}
		json.NewEncoder(w).Encode(out)
	})
	http.ListenAndServe(":8080", nil)
}
