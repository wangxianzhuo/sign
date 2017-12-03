package db

import (
	"database/sql"
	"fmt"
	"testing"
	"time"

	_ "github.com/lib/pq"
)

func Test_CreateTable(t *testing.T) {
	db, err := sql.Open("postgres", "host=localhost port=5432 user=dev password=dev dbname=dc sslmode=disable")
	if err != nil {
		t.Fatalf("database connect error: %v", err)
	}
	defer db.Close()

	cols := make(map[string]string)
	cols["id"] = "VARCHAR(36) NOT NULL"
	cols["user_id"] = "VARCHAR(36) NOT NULL"
	cols["reference_id"] = "VARCHAR(36) NOT NULL"
	cols["tag"] = "VARCHAR(50) DEFAULT 'signed'"
	cols["created_time"] = "TIMESTAMP WITH TIME ZONE"

	conditions := "PRIMARY KEY (id)"

	err = CreateTable("sign", cols, conditions, db)
	if err != nil {
		t.Fatalf("%v", err)
	}
}

func Test_DropTable(t *testing.T) {
	db, err := sql.Open("postgres", "host=localhost port=5432 user=dev password=dev dbname=dc sslmode=disable")
	if err != nil {
		t.Fatalf("database connect error: %v", err)
	}
	defer db.Close()

	err = DropTable("sign", db)
	if err != nil {
		t.Fatalf("%v", err)
	}
}

func Test_FormatInterface(t *testing.T) {
	var a int = -1
	var b int8 = -2
	var c int16 = -3
	var d int32 = -4
	var e int64 = 5

	var f uint = 1
	var g uint8 = 2
	var h uint16 = 3
	var i uint32 = 4
	var j uint64 = 6

	var aa float32 = 1.034
	var bb float64 = -102.343

	var ff = true

	var ss = "hello"

	sss := dd{
		a: 1,
		b: false,
		c: "go",
	}

	if formatInterface(a) != "-1" {
		t.Errorf("format int error")
	}
	if formatInterface(b) != "-2" {
		t.Errorf("format int8 error")
	}
	if formatInterface(c) != "-3" {
		t.Errorf("format int16 error")
	}
	if formatInterface(d) != "-4" {
		t.Errorf("format int32 error")
	}
	if formatInterface(e) != "5" {
		t.Errorf("format int64 error")
	}

	if formatInterface(f) != "1" {
		t.Errorf("format uint error")
	}
	if formatInterface(g) != "2" {
		t.Errorf("format uint8 error")
	}
	if formatInterface(h) != "3" {
		t.Errorf("format uint16 error")
	}
	if formatInterface(i) != "4" {
		t.Errorf("format uint32 error")
	}
	if formatInterface(j) != "6" {
		t.Errorf("format uint64 error")
	}

	if formatInterface(aa) != "1.034" {
		t.Logf("format float32 error, get value %v", formatInterface(aa))
	}
	if formatInterface(bb) != "-102.343" {
		t.Logf("format float64 error, get value %v", formatInterface(bb))
	}

	if formatInterface(ff) != "true" {
		t.Errorf("format bool error")
	}

	if formatInterface(ss) != "hello" {
		t.Errorf("format string error")
	}

	if formatInterface(sss) != "a = 1, b = false, c = go" {
		t.Errorf("format Stringer error, get value %v", formatInterface(sss))
	}
}

type dd struct {
	a int
	b bool
	c string
}

func (d dd) String() string {
	return fmt.Sprintf("a = %v, b = %v, c = %v", d.a, d.b, d.c)
}

func Test_Insert(t *testing.T) {
	db, err := sql.Open("postgres", "host=localhost port=5432 user=dev password=dev dbname=dc sslmode=disable")
	if err != nil {
		t.Fatalf("database connect error: %v", err)
	}
	defer db.Close()

	cols := make(map[string]interface{})
	cols["id"] = "1"
	cols["user_id"] = "22"
	cols["reference_id"] = "33"
	cols["tag"] = "signed"
	cols["created_time"] = time.Now()

	err = Insert("sign", cols, "", db)
	if err != nil {
		t.Fatalf("error: %v", err)
	}
}

func Test_Get(t *testing.T) {
	db, err := sql.Open("postgres", "host=localhost port=5432 user=dev password=dev dbname=dc sslmode=disable")
	if err != nil {
		t.Fatalf("database connect error: %v", err)
	}
	defer db.Close()

	cols := []string{"id", "user_id", "reference_id", "tag", "created_time"}
	condition := "id = '1'"

	result, err := Get("sign", cols, condition, db)
	if err != nil {
		t.Fatalf("error: %v", err)
	}

	fmt.Printf("result %v", result)
}
