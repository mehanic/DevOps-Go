package main

import (
	"database/sql"
	"fmt"
	"net/http"

	_ "github.com/lib/pq"
)

var db *sql.DB

func init() {
	var err error
	db, err = sql.Open("postgres", "postgres://postgres:new_password@localhost/postgres?sslmode=disable")
	if err != nil {
		panic(err)
	}

	if err = db.Ping(); err != nil {
		panic(err)
	}
	fmt.Println("You connected to your database.")
}

type Book struct {
	isbn   string
	title  string
	author string
	price  float32
}

func main() {
	http.HandleFunc("/books", booksIndex)
	http.ListenAndServe(":8080", nil)
}

func booksIndex(w http.ResponseWriter, r *http.Request) {
	if r.Method != "GET" {
		http.Error(w, http.StatusText(http.StatusMethodNotAllowed), http.StatusMethodNotAllowed)
		return
	}

	rows, err := db.Query("SELECT * FROM books_variant")
	if err != nil {
		http.Error(w, http.StatusText(500), 500)
		return
	}
	defer rows.Close()

	bks := make([]Book, 0)
	for rows.Next() {
		bk := Book{}
		err := rows.Scan(&bk.isbn, &bk.title, &bk.author, &bk.price) // order matters
		if err != nil {
			http.Error(w, http.StatusText(500), 500)
			return
		}
		bks = append(bks, bk)
	}
	if err = rows.Err(); err != nil {
		http.Error(w, http.StatusText(500), 500)
		return
	}

	for _, bk := range bks {
		fmt.Fprintf(w, "%s, %s, %s, $%.2f\n", bk.isbn, bk.title, bk.author, bk.price)
	}
}
// 978-1503261969, Emma, Jayne Austen, $9.44
// 978-1505255607, The Time Machine, H. G. Wells, $5.99
// 978-1503379640, The Prince, Niccolò Machiavelli, $6.99

// curl -i localhost:8080/books
// HTTP/1.1 200 OK
// Date: Sat, 01 Feb 2025 15:59:52 GMT
// Content-Length: 248
// Content-Type: text/plain; charset=utf-8

// 978-1505255607, The Time Machine, H. G. Wells, $5.99
// 978-1503261969, Emma, Jayne Austen, $9.44
// 978-1470184841, Metamorphosis, Franz Kafka, $5.90
// 978-1470184800, Milan, Nic Peterson, $7.90
// 978-1503379640, The Prince, Niccolò Machiavelllllli, $6.99

