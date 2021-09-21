package repository

import (
	"database/sql"
	"fmt"
)

var Db *sql.DB

func init() {
	var err error
	dataSourceName := fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8",
		"todo-app", "todo-password", "sample-api-db:3306", "todo",
	)
	Db, err = sql.Open("mysql", dataSourceName)
	if err != nil {
		panic(err)
	}
}
