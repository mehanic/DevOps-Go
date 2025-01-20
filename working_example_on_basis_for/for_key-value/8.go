package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	rand.Seed(time.Now().UTC().UnixNano())

	n := 5
	a := []int{}
	for i := 0; i < n; i++ {
		k := rand.Intn(100)
		a = append(a, k)
	}
	fmt.Println(a)
	fmt.Println("for")
	for i := 0; i < len(a); i++ {
		fmt.Print(a[i]+1, " ")
	}
	fmt.Println("")
	fmt.Println("for range")
	for _, v := range a {
		fmt.Print(v+1, " ")
	}
}
