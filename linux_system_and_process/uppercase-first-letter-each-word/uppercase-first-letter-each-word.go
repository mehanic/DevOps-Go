package main

import (
	"fmt"

	"golang.org/x/text/cases"
	"golang.org/x/text/language"
)

func main() {
	fmt.Println(cases.Title(language.Und).String("goSAMples.dev is the best Go bLog in the world!"))
}

// go run uppercase-first-letter-each-word.go                                                                                                                
// Gosamples.dev Is The Best Go Blog In The World!
