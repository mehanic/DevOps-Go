package main

import "fmt"

func main() {
	xi := make([]int, 0, 10)

	for i := 0; i < 100; i++ {
		xi = append(xi, i)
		fmt.Printf("value:%v len:%v cap:%v\n", i, len(xi), cap(xi))
	}

	fmt.Println("xi", xi)
}
