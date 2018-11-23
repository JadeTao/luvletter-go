package conf

import (
	"database/sql"
)

var db *sql.DB

// GetDB ...
func GetDB() *sql.DB {
	if db != nil {
		return db
	}
	newDB, err := sql.Open("mysql", DBConfig)
	db = newDB
	if err != nil {
		panic(err)
	}
	return newDB
}
