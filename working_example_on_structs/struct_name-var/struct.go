package main

import "fmt"

type person struct {
	first string
	last  string
	age   int
}

func main() {
	p1 := person{
		first: "kira",
		last:  "Shinigami",
		age:   220,
	}
	p2 := person{
		first: "adam",
		last:  "eve",
		age:   23000,
	}
	fmt.Println(p1, p2)
	fmt.Println(p1.first, p1.last, p1.age)
	fmt.Println(p2.first, p2.last, p2.age)
}

// {kira Shinigami 220} {adam eve 23000}
// kira Shinigami 220
// adam eve 23000
