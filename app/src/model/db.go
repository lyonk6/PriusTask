package model

import "database/sql"

var db *sql.DB

//SetDatabase for this application
func SetDatabase(database *sql.DB) {
    db = database
}
