package model

import "database/sql"

var db *sql.DB
var dbName string

//SetDatabase for this application
func SetDatabase(database *sql.DB, n string) {
	db = database
	dbName = n
}
