package main

import (
	"database/sql"
	"fmt"
	"os"
	"time"

	_ "github.com/go-sql-driver/mysql" //we import supported libraries for database/sql
)

func main() {
	db, err := Setup()
	if err != nil {
		panic(err)
	}

	if err := Exec(db); err != nil {
		panic(err)
	}
}

// Example hold the results of our queries
type Example struct {
	Name    string
	Created *time.Time
}

// Setup configures and returns our database
// connection poold
func Setup() (*sql.DB, error) {
	db, err := sql.Open("mysql", fmt.Sprintf("%s:%s@/gocookbook?parseTime=true", os.Getenv("MYSQLUSERNAME"), os.Getenv("MYSQLPASSWORD")))
	if err != nil {
		return nil, err
	}
	return db, nil
}

// Create makes a table called example
// and populates it
func Create(db *sql.DB) error {
	// create the database
	if _, err := db.Exec("CREATE TABLE example (name VARCHAR(20), created DATETIME)"); err != nil {
		return err
	}

	if _, err := db.Exec(`INSERT INTO example (name, created) values ("Aaron", NOW())`); err != nil {
		return err
	}

	return nil
}

// Exec takes a new connection
// creates tables, and later drops them
// and issues some queries
func Exec(db *sql.DB) error {
	// uncaught error on cleanup, but we always
	// want to cleanup
	defer db.Exec("DROP TABLE example")

	if err := Create(db); err != nil {
		return err
	}

	if err := Query(db, "Aaron"); err != nil {
		return err
	}
	return nil
}
