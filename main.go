package main

import (
	"database/sql"
	"flag"
	"fmt"
	"log"
)

func main() {
	dbName := flag.String("db", "postgres", "Select a database to get QPS from.")
	dbUser := flag.String("user", "postgres", "Select a user")
	dbPass := flag.String("pass", "postgres", "Provide password for the user.")

	flag.Parse()

	// creating a string that holds the connection info for connecting to a database.
	conStr := fmt.Sprintf("user=%s dbname=%s password=%s sslmode=disable", *dbUser, *dbName, *dbPass)

	// opening a connection to the database and returning it.
	db, err := sql.Open("postgres", conStr)
	if err != nil {
		log.Fatalln(err.Error())
	}
	_ = db
}
