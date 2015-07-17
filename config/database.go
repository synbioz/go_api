package config

import (
	"database/sql"
	_ "github.com/lib/pq"
	"log"
)

var db *sql.DB

func DatabaseInit() {
	var err error

	db, err = sql.Open("postgres", "user=theodelaune dbname=goapi")

	if err != nil {
		log.Fatal(err)
	}

	// Create Table cars if not exists
	createCarsTable()
}

func createCarsTable() {
	_, err := db.Exec("CREATE TABLE IF NOT EXISTS cars(id serial,manufacturer varchar(20), design varchar(20), style varchar(20), doors int, created_at timestamp default NULL, updated_at timestamp default NULL, constraint pk primary key(id))")

	if err != nil {
		log.Fatal(err)
	}
}

// Getter for db var
func Db() *sql.DB {
	return db
}
