package main

import (
	"encoding/gob"
	"log"
	"strings"
)

type Character struct {
	Name        string `gob:"name"`
	Surname     string `gob:"surname"`
	Job         string `gob:"job,omitempty"`
	YearOfBirth int    `gob:"year_of_birth,omitempty"`
}

func main() {
	var char = Character{
		Name:    "Albert",
		Surname: "Wilmarth",
		Job:     "assistant professor",
	}
	s := strings.Builder{}
	e := gob.NewEncoder(&s)
	if err := e.Encode(char); err != nil {
		log.Fatalln(err)
	}
	log.Printf("%q", s.String())
}


// ┌─────(free)─────(~/GOLANG/Hands-On-Systems-Programming-with-Go/Chapter10/gob/encoding) 
//  └> $ go run encoding.go 
// 2024/02/29 00:06:04 "D\xff\x81\x03\x01\x01\tCharacter\x01\xff\x82\x00\x01\x04\x01\x04Name\x01\f\x00\x01\aSurname\x01\f\x00\x01\x03Job\x01\f\x00\x01\vYearOfBirth\x01\x04\x00\x00\x00*\xff\x82\x01\x06Albert\x01\bWilmarth\x01\x13assistant professor\x00"
