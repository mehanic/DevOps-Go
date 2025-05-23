package main

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	_ "github.com/lib/pq"
)

// Global DB connection
var db *sql.DB

// Connect to the PostgreSQL database
func connect() (*sql.DB, error) {
	// Get the database connection details from environment variables
	connStr := fmt.Sprintf(
		"host=%s port=%s user=%s password=%s dbname=crossplane_examples sslmode=disable",
		os.Getenv("DB_HOST"),
		os.Getenv("DB_PORT"),
		os.Getenv("DB_USERNAME"),
		os.Getenv("DB_PASSWORD"),
	)

	var err error
	db, err = sql.Open("postgres", connStr)
	if err != nil {
		return nil, err
	}

	err = db.Ping()
	if err != nil {
		return nil, err
	}

	fmt.Println("Connected to the PostgreSQL database!")
	return db, nil
}

// Get the PostgreSQL version
func version() {
	fmt.Println("Connecting to the PostgreSQL database...")
	var dbVersion string
	err := db.QueryRow("SELECT version()").Scan(&dbVersion)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("PostgreSQL version:", dbVersion)
}

// Create the users table
func createTables() {
	command := `CREATE TABLE IF NOT EXISTS users (
		id SERIAL PRIMARY KEY,
		first_name TEXT NOT NULL,
		last_name TEXT NOT NULL
	);`

	_, err := db.Exec(command)
	if err != nil {
		log.Fatal("Error creating table:", err)
	}
}

// Insert a user into the users table
func insertUser(firstName, lastName string) int {
	sqlStatement := `INSERT INTO users(first_name, last_name) VALUES($1, $2) RETURNING id;`
	var id int
	err := db.QueryRow(sqlStatement, firstName, lastName).Scan(&id)
	if err != nil {
		log.Fatal("Error inserting user:", err)
	}
	return id
}

// Insert multiple users into the users table
func insertUsers(users [][2]string) {
	sqlStatement := `INSERT INTO users(first_name, last_name) VALUES($1, $2)`
	for _, user := range users {
		_, err := db.Exec(sqlStatement, user[0], user[1])
		if err != nil {
			log.Fatal("Error inserting users:", err)
		}
	}
}

// Select all users from the users table
func selectUsers() {
	rows, err := db.Query("SELECT id, first_name, last_name FROM users;")
	if err != nil {
		log.Fatal("Error selecting users:", err)
	}
	defer rows.Close()

	for rows.Next() {
		var id int
		var firstName, lastName string
		err := rows.Scan(&id, &firstName, &lastName)
		if err != nil {
			log.Fatal("Error scanning row:", err)
		}
		fmt.Printf("ID: %d, First Name: %s, Last Name: %s\n", id, firstName, lastName)
	}
	if err := rows.Err(); err != nil {
		log.Fatal("Error with rows:", err)
	}
}

func main() {
	// Connect to the database
	_, err := connect()
	if err != nil {
		log.Fatal("Failed to connect to the database:", err)
	}
	defer db.Close()

	// Get PostgreSQL version
	version()

	// Create the users table
	createTables()

	// Insert a single user
	id := insertUser("Eric", "Charles")
	fmt.Printf("Inserted user with ID: %d\n", id)

	// Insert multiple users
	users := [][2]string{
		{"John", "Doe"},
		{"Foo", "Bar"},
	}
	insertUsers(users)

	// Select and print all users
	selectUsers()
}
