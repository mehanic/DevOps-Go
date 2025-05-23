package main

import (
	//"github.com/PacktPublishing/Go-Programming-Cookbook-Second-Edition/chapter6/database"
	//"github.com/PacktPublishing/Go-Programming-Cookbook-Second-Edition/chapter6/dbinterface"
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql" //we import supported libraries for database/sql
)

func main() {
	db, err := database.Setup()
	if err != nil {
		panic(err)
	}

	tx, err := db.Begin()
	if err != nil {
		panic(err)
	}
	// this wont do anything if commit is successful
	defer tx.Rollback()

	if err := Exec(tx); err != nil {
		panic(err)
	}
	if err := tx.Commit(); err != nil {
		panic(err)
	}
}

//---

// Create makes a table called example
// and populates it
func Create(db DB) error {
	// create the database
	if _, err := db.Exec("CREATE TABLE example (name VARCHAR(20), created DATETIME)"); err != nil {
		return err
	}

	if _, err := db.Exec(`INSERT INTO example (name, created) values ("Aaron", NOW())`); err != nil {
		return err
	}

	return nil
}

//----

// Exec replaces the Exec from the previous
// recipe
func Exec(db DB) error {

	// uncaught error on cleanup, but we always
	// want to cleanup
	defer db.Exec("DROP TABLE example")

	if err := Create(db); err != nil {
		return err
	}

	if err := Query(db); err != nil {
		return err
	}
	return nil
}

// Query grabs a new connection
// and issues some queries printing the results
func Query(db DB) error {
	name := "Aaron"
	rows, err := db.Query("SELECT name, created FROM example where name=?", name)
	if err != nil {
		return err
	}
	defer rows.Close()
	for rows.Next() {
		var e database.Example
		if err := rows.Scan(&e.Name, &e.Created); err != nil {
			return err
		}
		fmt.Printf("Results:\n\tName: %s\n\tCreated: %v\n", e.Name, e.Created)
	}
	return rows.Err()
}

// DB is an interface that is satisfied
// by an sql.DB or an sql.Transaction
type DB interface {
	Exec(query string, args ...interface{}) (sql.Result, error)
	Prepare(query string) (*sql.Stmt, error)
	Query(query string, args ...interface{}) (*sql.Rows, error)
	QueryRow(query string, args ...interface{}) *sql.Row
}

// Transaction can do anything a Query can do
// plus Commit, Rollback, or Stmt
type Transaction interface {
	DB
	Commit() error
	Rollback() error
}
