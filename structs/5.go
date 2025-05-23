package main

import "fmt"

type Book struct {
	Name  string
	Price int
}

func main() {
	b1 := Book{
		Name:  "book1",
		Price: 150,
	}
	b1.Price = 300
	b2 := Book{
		Name:  "book2",
		Price: 132,
	}
	totalSum := b1.Price + b2.Price
	fmt.Println(totalSum)
}
