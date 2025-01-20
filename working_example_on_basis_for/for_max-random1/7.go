package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	//задает максимальное число которое может быть рандомным
	rand.Seed(time.Now().UTC().UnixNano())

	n := 5
	a := []int{}
	chet := []int{}
	nechet := []int{}
	for i := 0; i < n; i++ {
		k := rand.Intn(100)
		a = append(a, k)
	}
	fmt.Println(a)
	//sumi
	for i := 0; i < len(a); i++ {
		if a[i]%2 == 0 {
			chet = append(chet, a[i])
		} else {
			nechet = append(nechet, a[i])
		}
	}
	fmt.Println(chet)
	fmt.Println(nechet)
}
