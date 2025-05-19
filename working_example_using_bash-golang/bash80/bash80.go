package main

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	// Define MySQL connection details
	dsn := "root:new_password@tcp(127.0.0.1:3306)/"

	// Open connection to the MySQL server
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	// Test the connection
	if err := db.Ping(); err != nil {
		log.Fatal("Failed to connect to the database:", err)
	}

	// Execute the query
	var version string
	err = db.QueryRow("SELECT VERSION()").Scan(&version)
	if err != nil {
		log.Fatal("Failed to execute query:", err)
	}

	// Print the MySQL version
	fmt.Println("MySQL version:", version)
}

