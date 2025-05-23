package main

import (
	"flag"
	"fmt"
)

type hobbies []string

func main() {
	fptr := flag.String("Name", "Bapan", "Persons name")
	iptr := flag.Int("Age", 20, "Age of the person")
	lying := flag.Bool("Lying", false, "Is he lying")
	var country string
	flag.StringVar(&country, "country", "", "Country name")
	flag.Parse()
	fmt.Println("The name of the person is:", *fptr)
	fmt.Println("The age of the person is:", *iptr)
	fmt.Println("Is the person lying:", *lying)
	fmt.Println("The country of the person is:", country)
	for _, i := range flag.Args() {
		fmt.Println(i)
	}
}
