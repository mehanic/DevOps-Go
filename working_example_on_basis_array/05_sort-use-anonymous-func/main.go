package main

import (
	"fmt"
	"sort"
)

type person struct {
	name string
	age  int
}

func main() {
	xs := []person{{"ABC", 28}, {"AAB", 29}, {"AAC", 27}}

	fmt.Println("sort by age using anonymous function")
	sort.Slice(xs, func(i, j int) bool { return xs[i].age < xs[j].age })
	fmt.Println(xs)
}
