package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main1() {
	//задает максимальное число которое может быть рандомным
	rand.Seed(time.Now().UTC().UnixNano())

	n := 5
	a := []int{}
	for i := 0; i < n; i++ {
		k := rand.Intn(100)
		a = append(a, k)
	}
	fmt.Println(a)
	//sumi
	sumi := 0
	for i := 0; i < len(a); i++ {
		sumi += a[i]
	}
	fmt.Println(sumi)
}
