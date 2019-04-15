package main

import (
	"database/sql"
	"log"
)

func connect() *sql.DB {
	// db, err := sql.Open("mysql", "root:root@tcp(127.0.0.1:3306)/golang")
	db, err := sql.Open("mysql", "dummyUser:dummyUser01@tcp(db-intern.ciupl0p5utwk.us-east-1.rds.amazonaws.com:3306)/db_intern")

	if err != nil {
		log.Fatal(err)

	}

	return db
}
 
 
 
 