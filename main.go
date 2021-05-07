package main

import (
	"database/sql"
	"flag"
	"fmt"
	"log"

	_ "github.com/lib/pq"
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
	var qps int

	// qps (Query Per Second) is the combination of total commits and rollbacks
	// in a database so to get qps for a specific database in postgresql first get the
	// total commits for a database from pg_stat_database table which stores all the
	// transactions, commits, rollbacks for each database, and get the total rollback
	// for specific database from the same pg_stat_database table and add them together
	// then qps is acquired
	row := db.QueryRow("SELECT xact_commit+xact_rollback FROM pg_stat_database WHERE datname = $1", dbName)

	// retreiving the qps value from the row
	err = row.Scan(&qps)
	if err != nil {
		log.Println(err.Error())
	}
	log.Println(qps)
}
