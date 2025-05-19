package main

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	// Set up the database connection
	dsn := "root:Training2@^@tcp(127.0.0.1:3306)/" // Adjust the connection string as necessary
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}
	defer db.Close()

	// Test the connection
	if err := db.Ping(); err != nil {
		log.Fatalf("Failed to ping database: %v", err)
	}

	// Create the database
	_, err = db.Exec("CREATE DATABASE testdb")
	if err != nil {
		log.Fatalf("Failed to create database: %v", err)
	}

	fmt.Println("Database 'testdb' created successfully.")
}

