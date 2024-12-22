package db

import (
	"database/sql"
	_ "github.com/mattn/go-sqlite3"
)

var DB *sql.DB

func InitDB() {
	var err error
	DB, err = sql.Open("sqlite3", "api.db")

	if err != nil {
		panic("Could not connect to database.")
	}

	DB.SetMaxOpenConns(10)
	DB.SetMaxIdleConns(5)

	createTables()
}

func createTables() {
	seedData := map[string]string{
		"createUserTable": `CREATE TABLE IF NOT EXISTS users(
    		"id" INTEGER NOT NULL PRIMARY KEY AUTOINCREMENT,
    		"email" TEXT NOT NULL UNIQUE,
    		"password" TEXT NOT NULL
    		    	)`,

		"createEventsTable": `CREATE TABLE IF NOT EXISTS events(
    		"id" INTEGER NOT NULL PRIMARY KEY AUTOINCREMENT,
    		"name" TEXT NOT NULL,
    		"description" TEXT,
    		"location" TEXT NOT NULL ,
    		"date" DATETIME,
    		"user_id" INTEGER NOT NULL 
    		    	);`}

	for key, query := range seedData {
		_, err := DB.Exec(query)
		if err != nil {
			panic("Could not create table " + key + "-" + query)
		}
	}
}

func CloseDB() {
	err := DB.Close()
	if err != nil {
		return
	}
}
