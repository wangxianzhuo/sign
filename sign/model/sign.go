package model

import (
	"database/sql"
	"fmt"
	"time"

	uuid "github.com/satori/go.uuid"
)

type Sign struct {
	ID          string    `json:"id"`
	UserID      string    `json:"user_id"`
	ReferenceID string    `json:"reference_id"`
	Tag         string    `json:"tag"`
	CreatedTime time.Time `json:"created_time"`
}

func NewSign(userID, referenceID, tag string) Sign {
	return Sign{
		ID:          uuid.NewV4().String(),
		UserID:      userID,
		ReferenceID: referenceID,
		Tag:         tag,
		CreatedTime: time.Now(),
	}
}

func (s Sign) String() string {
	return fmt.Sprintf("Sign{ID = %s, UserID = %s, ReferenceID = %s, Tag = %s, CreatedTime = %s}", s.ID, s.UserID, s.ReferenceID, s.Tag, s.CreatedTime.Format("2006-1-2 15:04:05"))
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
	var userID string
	var referenceID string
	var tag string
	var createdTime time.Time

	row := dbConnection.QueryRow("select user_id, reference_id, tag, created_time from sign where id = $1", id)
	err = row.Scan(&userID, &referenceID, &tag, &createdTime)
	if err != nil {
		return Sign{}, fmt.Errorf("get sign by id [%v] error: %v", id, err)
	}

	sign.ID = id
	sign.UserID = userID
	sign.ReferenceID = referenceID
	sign.Tag = tag
	sign.CreatedTime = createdTime
	return
}

func Update(sign Sign, dbConnection *sql.DB) error {
	stmt, err := dbConnection.Prepare("update sign set user_id = $1, reference_id = $2, tag = $3 where id = $4")
	if err != nil {
		return fmt.Errorf("update sign [%v] error: %v", sign.ID, err)
	}
	defer stmt.Close()

	_, err = stmt.Exec(sign.UserID, sign.ReferenceID, sign.Tag, sign.ID)
	if err != nil {
		return fmt.Errorf("update sign [%v] error: %v", sign.ID, err)
	}

	return nil
}

func GetByUserIDAndReferenceID(userID, referenceID string, dbConnection *sql.DB) ([]Sign, error) {
	stmt, err := dbConnection.Prepare("select id, tag, created_time from sign where user_id = $1 and reference_id = $2")
	if err != nil {
		return nil, fmt.Errorf("get signs [user id = %s, reference id %s] error: %v", userID, referenceID, err)
	}
	defer stmt.Close()

	rows, err := stmt.Query(userID, referenceID)
	if err != nil {
		return nil, fmt.Errorf("get signs [user id = %s, reference id %s] error: %v", userID, referenceID, err)
	}

	var signs []Sign
	for rows.Next() {
		var id string
		var tag string
		var createdTime time.Time

		err := rows.Scan(&id, &tag, &createdTime)
		if err != nil {
			return nil, fmt.Errorf("get signs [user id = %s, reference id %s] error: %v", userID, referenceID, err)
		}

		sign := NewSign(userID, referenceID, tag)
		sign.ID = id
		sign.CreatedTime = createdTime
		signs = append(signs, sign)
	}

	return signs, nil
}
