
package main

import (
	"database/sql"
	"encoding/json"
	"net/http"
	"strconv"

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

	http.HandleFunc("/toggle", func(w http.ResponseWriter, r *http.Request) {
		id, _ := strconv.Atoi(r.URL.Query().Get("id"))
		db.Exec(`UPDATE todos SET done = 1 - done WHERE id = ?`, id)
		w.WriteHeader(http.StatusNoContent)
	})
	http.HandleFunc("/update", func(w http.ResponseWriter, r *http.Request) {
		var t Todo
		json.NewDecoder(r.Body).Decode(&t)
		db.Exec(`UPDATE todos SET text=? , done=? WHERE id=?`, t.Text, boolToInt(t.Done), t.ID)
		w.WriteHeader(http.StatusNoContent)
	})
	http.ListenAndServe(":8080", nil)
}

func boolToInt(b bool) int { if b { return 1 }; return 0 }
