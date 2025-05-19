package main

import (
	"bufio"
	"database/sql"
	"fmt"
	"os"
	"os/exec"
	"strings"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	// Connect to MySQL
	dsn := "user:password@tcp(127.0.0.1:3306)/" // Adjust DSN as needed
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		fmt.Printf("Error connecting to database: %v\n", err)
		return
	}
	defer db.Close()

	// Check the connection
	if err = db.Ping(); err != nil {
		fmt.Printf("Error pinging database: %v\n", err)
		return
	}

	// Get list of databases
	rows, err := db.Query("SHOW DATABASES")
	if err != nil {
		fmt.Printf("Error fetching databases: %v\n", err)
		return
	}
	defer rows.Close()

	var dbList []string
	for rows.Next() {
		var dbName string
		if err := rows.Scan(&dbName); err != nil {
			fmt.Printf("Error scanning database name: %v\n", err)
			return
		}
		dbList = append(dbList, dbName)
	}

	// Include "new..." option
	dbList = append(dbList, "new...")

	// Display database options
	fmt.Println("Select a database to initialize:")
	for i, db := range dbList {
		fmt.Printf("%d: %s\n", i+1, db)
	}

	reader := bufio.NewReader(os.Stdin)
	var selectedDB string
	for {
		fmt.Print("Enter your choice (number): ")
		input, _ := reader.ReadString('\n')
		input = strings.TrimSpace(input)

		// Check if the input is a number
		index := 0
		fmt.Sscanf(input, "%d", &index)

		if index < 1 || index > len(dbList) {
			fmt.Println("Invalid choice, please try again.")
			continue
		}

		selectedDB = dbList[index-1]
		break
	}

	if selectedDB == "new..." {
		// Prompt for new database name
		fmt.Print("Name for new database: ")
		newDBName, _ := reader.ReadString('\n')
		newDBName = strings.TrimSpace(newDBName)

		// Create new database
		_, err = db.Exec("CREATE DATABASE IF NOT EXISTS " + newDBName)
		if err != nil {
			fmt.Printf("Error creating database: %v\n", err)
			return
		}
		selectedDB = newDBName
		fmt.Printf("Created new database: %s\n", selectedDB)
	}

	// Initialize the selected database with the SQL file
	err = initializeDatabase(db, selectedDB)
	if err != nil {
		fmt.Printf("Error initializing database %s: %v\n", selectedDB, err)
		return
	}

	fmt.Printf("Successfully initialized database: %s\n", selectedDB)
}

func initializeDatabase(db *sql.DB, dbName string) error {
	// Execute SQL file to initialize the database
	file, err := os.Open("ourInit.sql")
	if err != nil {
		return fmt.Errorf("failed to open SQL file: %w", err)
	}
	defer file.Close()

	// Read the SQL file and execute it
	sqlScript, err := bufio.NewReader(file).ReadString('\n')
	if err != nil {
		return fmt.Errorf("failed to read SQL file: %w", err)
	}

	// Change the database for the execution
	_, err = db.Exec("USE " + dbName)
	if err != nil {
		return fmt.Errorf("failed to switch to database %s: %w", dbName, err)
	}

	_, err = db.Exec(sqlScript)
	if err != nil {
		return fmt.Errorf("failed to execute SQL script: %w", err)
	}

	return nil
}

