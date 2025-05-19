package main

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	// Define the Data Source Name (DSN) for MySQL
	dsn := "root:Training2@^@tcp(127.0.0.1:3306)/"

	// Open a new database connection
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		log.Fatalf("Failed to connect to MySQL: %v", err)
	}
	defer db.Close()

	// Check if the connection is alive
	if err := db.Ping(); err != nil {
		log.Fatalf("Failed to ping database: %v", err)
	}

	// Create a new database called "testdb"
	_, err = db.Exec("CREATE DATABASE testdb")
	if err != nil {
		log.Fatalf("Failed to create database: %v", err)
	}

	fmt.Println("Database 'testdb' created successfully")
}

