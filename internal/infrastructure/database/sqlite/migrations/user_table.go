package migrations

import "github.com/ilyasa1211/go-jwt-authentication/internal/infrastructure/database/sqlite"

func CreateUserTable(path string) {
	db := sqlite.NewSqliteConn(path)

	if _, err := db.Exec("CREATE TABLE users (id INTEGER PRIMARY KEY AUTOINCREMENT, name VARCHAR(255) NOT NULL, email VARCHAR(255) NOT NULL UNIQUE, password VARCHAR(255) NOT NULL)"); err != nil {
		panic(err)
	}
}
