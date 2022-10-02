package infra

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/mattn/go-sqlite3"
)

type SqlHandler struct {
	Conn *sql.DB
}

func NewSqlHandler() *SqlHandler {
	conn, err := sql.Open("sqlite3", "db/article.db")
	if err != nil {
		panic(err)
	}

	tableName := "articles"
	cmd := fmt.Sprintf(`
		CREATE TABLE IF NOT EXISTS %s (
			id INTEGER PRIMARY KEY AUTOINCREMENT,
			author STRING,
			title STRING,
			content STRING
		)`, tableName)

	_, err = conn.Exec(cmd)
	if err != nil {
		log.Fatal(err)
	}

	sqlHandler := new(SqlHandler)
	sqlHandler.Conn = conn
	return sqlHandler
}
