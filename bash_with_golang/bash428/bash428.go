package main

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	// Set up the database connection
	// Replace "root:Training2@^@tcp(127.0.0.1:3306)/" with your actual connection string
	dsn := "root:Training2@^@tcp(127.0.0.1:3306)/"
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}
	defer db.Close()

	// Test the connection
	if err := db.Ping(); err != nil {
		log.Fatalf("Failed to ping database: %v", err)
	}

	// Prepare and execute the query
	var version string
	query := "SELECT VERSION()"
	err = db.QueryRow(query).Scan(&version)
	if err != nil {
		log.Fatalf("Failed to execute query: %v", err)
	}

	// Print the MySQL version
	fmt.Println("MySQL Version:", version)
}

