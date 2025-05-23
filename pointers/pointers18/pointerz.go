package main

import "fmt"

func main() {
	a := 5
	var pa *int
	pa = &a //&a 
	fmt.Println(pa)
	fmt.Println(*pa) //5

}
