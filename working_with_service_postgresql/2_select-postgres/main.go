package main

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

type Book struct {
	isbn   string
	title  string
	author string
	price  float32
}

func main() {
	db, err := sql.Open("postgres", "postgres://postgres:new_password@localhost/postgres?sslmode=disable")
	if err != nil {
		panic(err)
	}
	defer db.Close()

	if err = db.Ping(); err != nil {
		panic(err)
	}
	fmt.Println("You connected to your database.")

	rows, err := db.Query("SELECT * FROM books_variant;")
	if err != nil {
		panic(err)
	}
	defer rows.Close()

	bks := make([]Book, 0)
	for rows.Next() {
		bk := Book{}
		err := rows.Scan(&bk.isbn, &bk.title, &bk.author, &bk.price) // order matters
		if err != nil {
			panic(err)
		}
		bks = append(bks, bk)
	}
	if err = rows.Err(); err != nil {
		panic(err)
	}

	for _, bk := range bks {
		// fmt.Println(bk.isbn, bk.title, bk.author, bk.price)
		fmt.Printf("%s, %s, %s, $%.2f\n", bk.isbn, bk.title, bk.author, bk.price)
	}
}



// #alter bond's role

// alter user bond with superuser;


// # switch to your bookstore database
// You should already have a ```bookstore``` database:

// list databases
// ```
// \l
// ```

// switch into that database

// \c bookstore


// directory of tables, if any

// \d


// # create table

// CREATE TABLE books (
//   isbn    char(14)     PRIMARY KEY NOT NULL,
//   title   varchar(255) NOT NULL,
//   author  varchar(255) NOT NULL,
//   price   decimal(5,2) NOT NULL
// );


// directory of tables

// \d


// details of table ```books```
// \d books

// # insert records

// INSERT INTO books (isbn, title, author, price) VALUES
// ('978-1503261969', 'Emma', 'Jayne Austen', 9.44),
// ('978-1505255607', 'The Time Machine', 'H. G. Wells', 5.99),
// ('978-1503379640', 'The Prince', 'Niccolò Machiavelli', 6.99);

// SELECT * FROM books;

