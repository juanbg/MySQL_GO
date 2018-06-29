package database

import (
	"database/sql"
	"log"
)

/*
DBConn Function which opens a new connection with db
*/
func DBConn() (db *sql.DB) {
	dbDriver := "mysql"
	dbUser := "root"
	dbPass := "vrajupaj"
	dbName := "example"
	db, err := sql.Open(dbDriver, dbUser+":"+dbPass+"@/"+dbName)

	if err != nil {
		panic(err.Error())
	} else {
		log.Println("Connection was opened correctly")
	}

	return db
}
