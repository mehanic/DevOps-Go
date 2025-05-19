package main

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	var userName, mysqlPwd, mysqlCmd, mysqlDb string

	// Prompt for MySQL user
	fmt.Print("MySQL User: ")
	fmt.Scanln(&userName)

	// Prompt for MySQL password without echoing
	fmt.Print("MySQL Password: ")
	fmt.Scanln(&mysqlPwd)

	// Prompt for MySQL command
	fmt.Print("MySQL Command: ")
	fmt.Scanln(&mysqlCmd)

	// Prompt for MySQL database
	fmt.Print("MySQL DB: ")
	fmt.Scanln(&mysqlDb)

	// Create a DSN (Data Source Name)
	dsn := fmt.Sprintf("%s:%s@/%s", userName, mysqlPwd, mysqlDb)

	// Open the database connection
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		log.Fatalf("Error connecting to the database: %v", err)
	}
	defer db.Close()

	// Execute the MySQL command
	rows, err := db.Query(mysqlCmd)
	if err != nil {
		log.Fatalf("Error executing MySQL command: %v", err)
	}
	defer rows.Close()

	// Process the result set
	cols, err := rows.Columns()
	if err != nil {
		log.Fatalf("Error fetching columns: %v", err)
	}

	// Create a slice to hold the values
	values := make([]sql.RawBytes, len(cols))
	scanArgs := make([]interface{}, len(cols))
	for i := range values {
		scanArgs[i] = &values[i]
	}

	// Print the results
	for rows.Next() {
		err := rows.Scan(scanArgs...)
		if err != nil {
			log.Fatalf("Error scanning row: %v", err)
		}

		for _, value := range values {
			// Convert to string and print
			fmt.Print(string(value), "\t")
		}
		fmt.Println()
	}

	// Check for errors after iterating through rows
	if err = rows.Err(); err != nil {
		log.Fatalf("Error iterating through rows: %v", err)
	}
}

