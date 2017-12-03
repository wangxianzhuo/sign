package sign

import "time"
import "database/sql"
import "fmt"

type Sign struct {
	ID          string    `json:"id"`
	UserID      string    `json:"user_id"`
	ReferenceID string    `json:"reference_id"`
	Tag         string    `json:"tag"`
	CreatedTime time.Time `json:"created_time"`
}

func Insert(sign Sign, dbConnection *sql.DB) error {
	stmt, err := dbConnection.Prepare("insert into sign(id, user_id, reference_id, tag, created_time) values($1, $2, $3, $4, $5)")
	if err != nil {
		return fmt.Errorf("insert sign error: %v", err)
	}
	defer stmt.Close()

	_, err = stmt.Exec(sign.ID, sign.UserID, sign.ReferenceID, sign.Tag, sign.CreatedTime)
	if err != nil {
		return fmt.Errorf("insert sign error: %v", err)
	}

	return nil
}

func Delete(id string, dbConnection *sql.DB) error {
	stmt, err := dbConnection.Prepare("delete from sign where id = $1")
	if err != nil {
		return fmt.Errorf("delete sign error: %v", err)
	}
	defer stmt.Close()
	_, err = stmt.Exec(id)
	if err != nil {
		return fmt.Errorf("delete sign error: %v", err)
	}

	return nil
}

func Get(id string, dbConnection *sql.DB) (sign Sign, err error) {
	row := dbConnection.QueryRow("")
	return
}
