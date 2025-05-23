package main

import (
	"day01/readdb"
	"flag"
	"fmt"
)

func main() {
	filename := flag.String("f", "", "json or xml file")
	flag.Parse()
	// fmt.Println(*filename)

	db, err, format := readdb.ParsFile(*filename)
	if err != nil {
		fmt.Println(err.Error())

		return
	}
	readdb.PrintRecipes(db, format)
}
