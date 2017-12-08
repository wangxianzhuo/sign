package env

import (
	"database/sql"

	_ "github.com/lib/pq"
)

var DB *sql.DB

func InitDB(connection string) error {
	var err error

	DB, err = sql.Open("postgres", connection)
	if err != nil {
		return err
	}
	return nil
}
