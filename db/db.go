package db

import (
	"bytes"
	"database/sql"
	"fmt"
	"strconv"
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
		return fmt.Errorf("create table %v error: %v", tableName, err)
	}

	return nil
}

func DropTable(tableName string, dbConnection *sql.DB) error {
	sql := fmt.Sprintf("drop table %v", tableName)
	_, err := dbConnection.Exec(sql)
	if err != nil {
		return fmt.Errorf("drop table %v error: %v", tableName, err)
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

	result := make([]map[string]interface{}, 0)
	for rows.Next() {
		row := make(map[string]interface{})
		values := make([]interface{}, len(cols))
		for index := 0; index < len(cols); index++ {
			tmp := ""
			values[index] = &tmp
		}

		err := rows.Scan(values...)
		if err != nil {
			return nil, err
		}
		for index := 0; index < len(cols); index++ {
			// tmp := values[index]

			// row[cols[index]] = tmp
		}
		result = append(result, row)
	}
	err = rows.Err()
	if err != nil {
		return nil, err
	}

	return result, nil
}

func Insert(tableName string, cols map[string]interface{}, condition string, dbConnection *sql.DB) error {
	var buffer bytes.Buffer

	buffer.WriteString("insert into ")
	buffer.WriteString(tableName)
	buffer.WriteString("(")
	cn := make([]string, len(cols))
	cv := make([]string, len(cols))
	args := make([]interface{}, len(cols))
	i := 0
	for k, v := range cols {
		cn[i] = k
		cv[i] = "$" + strconv.FormatInt(int64(i+1), 10)
		args[i] = v
		i++
	}
	buffer.WriteString(strings.Join(cn, ","))
	buffer.WriteString(") values(")
	buffer.WriteString(strings.Join(cv, ","))
	buffer.WriteString(")")

	fmt.Println(buffer.String())
	stmt, err := dbConnection.Prepare(buffer.String())
	if err != nil {
		return fmt.Errorf("Insert into %v error: %v", tableName, err)
	}
	defer stmt.Close()

	_, err = stmt.Exec(args...)
	if err != nil {
		return fmt.Errorf("Insert into %v error: %v", tableName, err)
	}
	return nil
}

func Delete(tableName string, condition string, dbConnection *sql.DB) error {
	return nil
}
func formatInterface(input interface{}) string {
	switch v := input.(type) {
	case int:
		return strconv.FormatInt(int64(v), 10)
	case int8:
		return strconv.FormatInt(int64(v), 10)
	case int16:
		return strconv.FormatInt(int64(v), 10)
	case int32:
		return strconv.FormatInt(int64(v), 10)
	case int64:
		return strconv.FormatInt(int64(v), 10)
	case float32:
		return strconv.FormatFloat(float64(v), 'f', 10, 64)
	case float64:
		return strconv.FormatFloat(float64(v), 'f', 10, 64)
	case bool:
		return strconv.FormatBool(v)
	case uint:
		return strconv.FormatUint(uint64(v), 10)
	case uint8:
		return strconv.FormatUint(uint64(v), 10)
	case uint16:
		return strconv.FormatUint(uint64(v), 10)
	case uint32:
		return strconv.FormatUint(uint64(v), 10)
	case uint64:
		return strconv.FormatUint(uint64(v), 10)
	case string:
		return "'" + v + "'"
	case fmt.Stringer:
		return "'" + v.String() + "'"
	default:
		return ""
	}
}
