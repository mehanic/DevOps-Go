package main

import "fmt"

type ProductType struct {
	Name string
}

type Product struct {
	Name  string
	Price int
	Type  ProductType
}

func main() {
	type1 := ProductType{Name: "pc"}
	type2 := ProductType{Name: "tv"}
	p1 := Product{Name: "lenovo", Price: 300, Type: type1}
	p2 := Product{Name: "samsung", Price: 1000, Type: type2}
	fmt.Println(p1.Name)
	fmt.Println(p1.Type.Name)
	fmt.Println(p2.Name)
	fmt.Println(p2.Type.Name)
}
