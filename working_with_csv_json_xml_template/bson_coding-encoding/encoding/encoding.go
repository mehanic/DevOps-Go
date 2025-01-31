package main

import (
	"log"

	"gopkg.in/mgo.v2/bson"
)

type Character struct {
	Name        string `bson:"name"`
	Surname     string `bson:"surname"`
	Job         string `bson:"job,omitempty"`
	YearOfBirth int    `bson:"year_of_birth,omitempty"`
}

func main() {
	var char = Character{
		Name:    "Robert",
		Surname: "Olmstead",
	}
	b, err := bson.Marshal(char)
	if err != nil {
		log.Fatalln(err)
	}
	log.Printf("%q", b)
}

// 2024/02/29 00:04:38 ",\x00\x00\x00\x02name\x00\a\x00\x00\x00Robert\x00\x02surname\x00\t\x00\x00\x00Olmstead\x00\x00"
