package main

import "fmt"

func main() {
	books := [...]string{"red", "orange", "yellow", "blue", "violet"}
	//0->red        1->orange   2->yellow        3->blue        4->violet
	fmt.Println(books)
	fmt.Println(books[3])
}
