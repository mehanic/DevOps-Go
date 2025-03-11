package main

import (
	"fmt"
	"sort"
)

type person struct {
	name string
	age  int
}

type SortByAge []person

func (s SortByAge) Len() int {
	return len(s)
}

func (s SortByAge) Less(i, j int) bool {
	return s[i].age < s[j].age
}

func (s SortByAge) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}

type SortByName []person

func (s SortByName) Len() int {
	return len(s)
}

func (s SortByName) Less(i, j int) bool {

	return s[i].name < s[j].name
}

func (s SortByName) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}

func main() {
	xs := []person{{"ABC", 28}, {"AAB", 29}, {"AAC", 27}}

	fmt.Println("sort by name")
	sort.Sort(SortByName(xs))
	fmt.Println(xs)

	fmt.Println("reverse")
	sort.Sort(sort.Reverse(SortByName(xs)))
	fmt.Println(xs)

	fmt.Println("sort by age")
	sort.Sort(SortByAge(xs))
	fmt.Println(xs)

	fmt.Println("reverse")
	sort.Sort(sort.Reverse(SortByAge(xs)))
	fmt.Println(xs)
}
