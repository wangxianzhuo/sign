package db

import (
	"bytes"
	"database/sql"
	"fmt"
	"strings"
)

func CreateTable(tableName string, cols map[string]string, conditions string, dbConnection *sql.DB) error {
	var buffer bytes.Buffer
	buffer.WriteString("create table " + tableName + "(")
	for k, v := range cols {
		buffer.WriteString(k)
		buffer.WriteString(" ")
		buffer.WriteString(v)
		buffer.WriteString(",")
	}
	buffer.WriteString(conditions)
	buffer.WriteString(")")

	sql := buffer.String()
	_, err := dbConnection.Exec(sql)
	if err != nil {
		return err
	}

	return nil
}

func Get(tableName string, cols []string, condition string, dbConnection *sql.DB) (interface{}, error) {
	c := strings.Join(cols, " ,")
	if condition != "" {
		condition = "where " + condition
	}
	sql := fmt.Sprintf("select %s from %s %s", c, tableName, condition)

	rows, err := dbConnection.Query(sql)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	result := make([]map[string]interface{}, 1)
	for rows.Next() {
		row := make(map[string]interface{})
		values := make([]string, len(cols))

		err := rows.Scan(&values)
		if err != nil {
			return nil, err
		}
		for index := 0; index < len(cols); index++ {
			row[cols[index]] = values[index]
		}
		result = append(result, row)
	}
	err = rows.Err()
	if err != nil {
		return nil, err
	}

	return result, nil
}
