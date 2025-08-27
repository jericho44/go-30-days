
package main

import (
	"database/sql"
	"net/http"

	"github.com/gin-gonic/gin"
	_ "modernc.org/sqlite"
)

type Todo struct {
	ID   int    `json:"id"`
	Text string `json:"text" binding:"required"`
	Done bool   `json:"done"`
}

func main() {
	db, _ := sql.Open("sqlite", "file:todos.db")
	defer db.Close()
	db.Exec(`CREATE TABLE IF NOT EXISTS todos(id INTEGER PRIMARY KEY, text TEXT NOT NULL, done INTEGER NOT NULL DEFAULT 0)`)

	r := gin.Default()

	r.GET("/health", func(c *gin.Context) { c.String(http.StatusOK, "ok") })

	r.GET("/todos", func(c *gin.Context) {
		rows, _ := db.Query(`SELECT id,text,done FROM todos ORDER BY id DESC`)
		defer rows.Close()
		out := []Todo{}
		for rows.Next() {
			var t Todo
			var d int
			rows.Scan(&t.ID, &t.Text, &d)
			t.Done = d == 1
			out = append(out, t)
		}
		c.JSON(200, out)
	})

	r.POST("/todos", func(c *gin.Context) {
		var t Todo
		if err := c.ShouldBindJSON(&t); err != nil {
			c.JSON(400, gin.H{"error": err.Error()})
			return
		}
		db.Exec(`INSERT INTO todos(text,done) VALUES(?,0)`, t.Text)
		c.Status(201)
	})

	r.PATCH("/todos/:id/toggle", func(c *gin.Context) {
		id := c.Param("id")
		db.Exec(`UPDATE todos SET done = 1 - done WHERE id = ?`, id)
		c.Status(204)
	})

	r.PUT("/todos/:id", func(c *gin.Context) {
		id := c.Param("id")
		var t Todo
		if err := c.ShouldBindJSON(&t); err != nil {
			c.JSON(400, gin.H{"error": err.Error()})
			return
		}
		db.Exec(`UPDATE todos SET text=?, done=? WHERE id=?`, t.Text, boolToInt(t.Done), id)
		c.Status(204)
	})

	r.DELETE("/todos/:id", func(c *gin.Context) {
		id := c.Param("id")
		db.Exec(`DELETE FROM todos WHERE id=?`, id)
		c.Status(204)
	})

	r.Run(":8080")
}

func boolToInt(b bool) int { if b { return 1 }; return 0 }
