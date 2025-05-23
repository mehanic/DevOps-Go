package main

import "fmt"

type Product struct {
	Name string
	Cost int
}

func main() {
	p0 := Product{
		Name: "product0",
		Cost: 123,
	}
	p3 := Product{
		Name: "product3",
		Cost: 400,
	}
	products := []Product{
		{Name: "product1", Cost: 100},
		{Name: "product2", Cost: 300},
		p0,
	}
	products = append(products, p3)
	fmt.Println(products)
}
