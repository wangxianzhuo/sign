package db

import (
	"database/sql"
	"testing"

	_ "github.com/lib/pq"
)

func Test_DB(t *testing.T) {
	db, err := sql.Open("postgres", "host=localhost port=5432 user=dev password=dev dbname=dc sslmode=disable")
	if err != nil {
		t.Fatalf("database connect error: %v", err)
	}

	cols := make(map[string]string)
	cols["id"] = "VARCHAR(36) NOT NULL"
	cols["user_id"] = "VARCHAR(36) NOT NULL"
	cols["reference_id"] = "VARCHAR(36) NOT NULL"
	cols["tag"] = "VARCHAR(50) DEFAULT 'signed'"
	cols["created_time"] = "TIMESTAMP WITH TIME ZONE"

	conditions := "PRIMARY KEY (id)"

	err = CreateTable("sign", cols, conditions, db)
	if err != nil {
		t.Fatalf("create table %v error %v", "sign", err)
	}
}
