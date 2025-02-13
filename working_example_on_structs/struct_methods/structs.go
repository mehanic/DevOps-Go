package main

import "fmt"

func Structs() {
	fmt.Println(product{name: "Jack", brand: "Apple", price: 40})
	c := product{"a", "b", 3}
	println(c.brand)
	c.save()
}

func (p product) save() {
	fmt.Println("super")
}

type product struct {
	name  string
	brand string
	price int
}


func main() {
	Structs()
}

// {Jack Apple 40}
// b
// super