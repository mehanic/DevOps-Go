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
	for i := 0; i < len(a); i++ {
		a[i] = 5
	}
	fmt.Println(a)
	for i, _ := range a {
		a[i] = 10
	}
	fmt.Println(a)
}
