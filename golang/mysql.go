package main

import (
	"database/sql"
	"log"
)

func connect() *sql.DB {
	db, err := sql.Open("mysql", "root:root@tcp(127.0.0.1:3306)/golang")

	if err != nil {
		log.Fatal(err)
		
	}

	return db
}


