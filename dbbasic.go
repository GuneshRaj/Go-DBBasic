package main

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
)

//TODO
func addRecord(db *sql.DB, data map[string]string) error {

	return nil
}

//TODO
func addMultipleRecord(db *sql.DB, dataList []map[string]string) error {
	return nil
}

//TODO
func updateRecord(db *sql.DB, data map[string]string) error {
	return nil
}

//TODO
func updateMultipleRecord(db *sql.DB, dataList []map[string]string) error {
	return nil
}

func updateStatement(db *sql.DB, sql string) error {
	_, errx := db.Exec(sql)
	if errx != nil {
		return errx
	}
	return nil
}

func selectStatement(db *sql.DB, sql string) ([]map[string]string, error) {
	rows, err := db.Query(sql)
	if err != nil {
		return nil, err
	}
	columns, err := rows.Columns()
	if err != nil {
		return nil, err
	}
	count := len(columns)
	values := make([]interface{}, count)
	valuePtrs := make([]interface{}, count)
	clist := []map[string]string{}

	for rows.Next() {
		for i, _ := range columns {
			valuePtrs[i] = &values[i]
		}
		rows.Scan(valuePtrs...)
		var mdata = make(map[string]string)
		for i, col := range columns {
			val := values[i]
			b, _ := val.([]byte)
			mdata[col] = string(b)
		}
		clist = append(clist, mdata)
	}
	return clist, nil
}

func selectSingleStatement(db *sql.DB, sql string) (map[string]string, error) {
	rows, err := db.Query(sql)
	if err != nil {
		return nil, err
	}
	columns, err := rows.Columns()
	if err != nil {
		return nil, err
	}
	count := len(columns)
	values := make([]interface{}, count)
	valuePtrs := make([]interface{}, count)

	var mdata = make(map[string]string)
	for rows.Next() {
		for i, _ := range columns {
			valuePtrs[i] = &values[i]
		}
		rows.Scan(valuePtrs...)

		for i, col := range columns {
			val := values[i]
			b, _ := val.([]byte)
			mdata[col] = string(b)
		}
		break
	}
	return mdata, nil
}

func selectData(db *sql.DB, sql string) (string, error) {
	data := ""
	rows, err := db.Query(sql)
	if err != nil {
		return data, err
	}
	for rows.Next() {
		rows.Scan(&data)
		break
	}
	return data, nil
}
