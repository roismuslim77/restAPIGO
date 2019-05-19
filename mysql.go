package main

import (
	"database/sql"
	"log"
)

func connect() *sql.DB {
	db, err := sql.Open("mysql",
	"root:12345678@tcp(127.0.0.1:3306)/godb")

	if err != nil{
		log.Fatal(err)
	}

	return db
}