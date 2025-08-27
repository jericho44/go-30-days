
package main

import (
	"database/sql"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	_ "modernc.org/sqlite"
)

type Todo struct {
	ID   int    `json:"id"`
	Text string `json:"text" binding:"required,min=3"`
	Done bool   `json:"done"`
}

func main() {
	port := os.Getenv("PORT")
	if port == "" { port = "8080" }

	db, _ := sql.Open("sqlite", "file:todos.db")
	defer db.Close()
	db.Exec(`CREATE TABLE IF NOT EXISTS todos(id INTEGER PRIMARY KEY, text TEXT NOT NULL, done INTEGER NOT NULL DEFAULT 0)`)

	r := gin.New()
	r.Use(gin.Logger(), gin.Recovery(), cors())

	r.GET("/health", func(c *gin.Context) { c.String(http.StatusOK, "ok") })
	// ... (routes sama seperti Hari 29, boleh copy persis) ...
	r.Run(":" + port)
}

func cors() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Header("Access-Control-Allow-Origin", "*")
		c.Header("Access-Control-Allow-Methods", "GET,POST,PUT,PATCH,DELETE,OPTIONS")
		c.Header("Access-Control-Allow-Headers", "Content-Type, Authorization")
		if c.Request.Method == http.MethodOptions {
			c.Status(http.StatusNoContent)
			c.Abort()
			return
		}
		c.Next()
	}
}
