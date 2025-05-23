package main

import "fmt"

type person struct {
	name string
	age  int
}

func main() {
	p := &person{"clement", 29}

	fmt.Println(p)
	fmt.Printf("%T\n", p)
	fmt.Println(p.name)
	fmt.Println(p.age)
}

// &{clement 29}
// *main.person
// clement
// 29
