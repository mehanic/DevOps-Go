package main

import (
	"database/sql"
	"log"

	_ "github.com/lib/pq"
)

func main() {
	connectionString := "postgres://postgres:together88@localhost/my_app_dev?sslmode=disable"
	db, err := sql.Open("postgres", connectionString)
	if err != nil {
		log.Panicln("errrrr")
		log.Fatal(err)
	}
	log.Print("connected")
	// log.Printf("%v %t", db, db)
	// age := 21
	actorRows, err := db.Query("SELECT actor_name FROM actor")
	if err != nil {
		log.Fatalln(err)
		// panic(err)
	}
	// log.Println("do something")
	// log.Print(rows)
	for actorRows.Next() {
		var actorName sql.NullString
		err = actorRows.Scan(&actorName)
		if err != nil {
			log.Fatal(err)
		}
		log.Printf("%#v", actorName)
	}
	// …
}
