package main

import (
	"flag"
	"fmt"
)

func main() {
	dbName := flag.String("db", "postgres", "Select a database to get QPS from.")
	dbUser := flag.String("user", "postgres", "Select a user")
	dbPass := flag.String("pass", "postgres", "Provide password for the user.")

	flag.Parse()

	fmt.Println(*dbName)
	fmt.Println(*dbUser)
	fmt.Println(*dbPass)

}
