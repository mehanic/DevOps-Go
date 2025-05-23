package main

import (
	"fmt"
)

func main() {
	age := 66
	var price int

	if age < 4 {
		price = 0
	} else if age < 18 {
		price = 5
	} else if age < 65 {
		price = 10
	} else if age > 65 {
		price = 5
	}

	fmt.Printf("Your admission cost is $%d.\n", price)
}

//Your admission cost is $5.
