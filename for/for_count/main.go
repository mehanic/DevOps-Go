package main

import (
	"fmt"
	"github.com/romzasandysman/datafile"
	"log"
)

func main() {
	lines, err := datafile.GetStrings("votes.txt")

	if err != nil{
		log.Fatal(err)
	}

	counts := make(map[string]int)

	for _, line := range lines {
		counts[line]++
	}

	for name, count := range counts{
		fmt.Printf("%s: %d\n\n", name, count)
	}
}
