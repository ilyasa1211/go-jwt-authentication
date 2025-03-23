package sqlite

import (
	"database/sql"

	_ "github.com/mattn/go-sqlite3"
)

func NewSqliteConn(path string) *sql.DB {
	db, err := sql.Open("sqlite3", path)

	if err != nil {
		panic(err)
	}

	return db
}
