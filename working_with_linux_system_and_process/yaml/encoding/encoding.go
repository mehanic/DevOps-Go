package main

import (
	"log"
	"os"

	"gopkg.in/yaml.v2"
)

type Character struct {
	Name        string `yaml:"name"`
	Surname     string `yaml:"surname"`
	Job         string `yaml:"job,omitempty"`
	YearOfBirth int    `yaml:"year_of_birth,omitempty"`
}

func main() {
	var chars = []Character{{
		Name:        "William",
		Surname:     "Dyer",
		Job:         "professor",
		YearOfBirth: 1875,
	}, {
		Surname: "Danforth",
		Job:     "student",
	}}
	e := yaml.NewEncoder(os.Stdout)
	if err := e.Encode(chars); err != nil {
		log.Fatalln(err)
	}
}

// └> $ go run encoding.go
// - name: William
//  surname: Dyer
//  job: professor
//  year_of_birth: 1875
// - name: ""
//  surname: Danforth
//  job: student
