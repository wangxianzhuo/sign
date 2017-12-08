package model

import (
	"database/sql"
	"testing"
	"time"

	_ "github.com/lib/pq"
	uuid "github.com/satori/go.uuid"
)

func Test_Insert(t *testing.T) {
	db, err := sql.Open("postgres", "host=127.0.0.1 port=5432 user=dev password=dev dbname=sign_in sslmode=disable")
	if err != nil {
		t.Fatalf("error: %v", err)
	}

	sign := Sign{
		ID:          "9cfa2bf8-b730-4564-bd9b-dc529c70b848",
		UserID:      uuid.NewV4().String(),
		ReferenceID: uuid.NewV4().String(),
		Tag:         "signed",
		CreatedTime: time.Now(),
	}

	err = Insert(sign, db)
	if err != nil {
		t.Fatalf("error: %v", err)
	}
}

func Test_Get(t *testing.T) {
	db, err := sql.Open("postgres", "host=127.0.0.1 port=5432 user=dev password=dev dbname=sign_in sslmode=disable")
	if err != nil {
		t.Fatalf("error: %v", err)
	}

	sign, err := Get("9cfa2bf8-b730-4564-bd9b-dc529c70b848", db)
	if err != nil {
		t.Fatalf("error: %v", err)
	}
	t.Log(sign)
}

func Test_Update(t *testing.T) {
	db, err := sql.Open("postgres", "host=127.0.0.1 port=5432 user=dev password=dev dbname=sign_in sslmode=disable")
	if err != nil {
		t.Fatalf("error: %v", err)
	}
	sign, err := Get("9cfa2bf8-b730-4564-bd9b-dc529c70b848", db)
	if err != nil {
		t.Fatalf("error: %v", err)
	}
	t.Log(sign)
	sign.Tag = "unsigned"

	err = Update(sign, db)
	if err != nil {
		t.Fatalf("error: %v", err)
	}
	sign, err = Get("9cfa2bf8-b730-4564-bd9b-dc529c70b848", db)
	if err != nil {
		t.Fatalf("error: %v", err)
	}
	t.Log(sign)
}

func Test_Delete(t *testing.T) {
	db, err := sql.Open("postgres", "host=127.0.0.1 port=5432 user=dev password=dev dbname=sign_in sslmode=disable")
	if err != nil {
		t.Fatalf("error: %v", err)
	}

	err = Delete("9cfa2bf8-b730-4564-bd9b-dc529c70b848", db)
	if err != nil {
		t.Fatalf("error: %v", err)
	}
}
