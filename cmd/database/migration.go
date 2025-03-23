package main

import (
	"errors"
	"os"

	"github.com/ilyasa1211/go-jwt-authentication/internal/infrastructure/database/sqlite/migrations"
)

func main() {

	path := "data/data.db"

	if _, err := os.Stat(path); errors.Is(err, os.ErrNotExist) {
		if err := os.MkdirAll("data", 0755); err != nil {
			panic(err)
		}

		_, err := os.Create(path)
		if err != nil {
			panic(err)
		}
	}

	migrations.CreateUserTable(path)
}
