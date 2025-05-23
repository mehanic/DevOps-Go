package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	//bubble sort
	rand.Seed(time.Now().UTC().UnixNano())
	n := 5
	a := []int{}
	b := []int{3, 4, 5}
	for i := 0; i < n; i++ {
		k := rand.Intn(20)
		a = append(a, k)
	}
	fmt.Println(a)
	for i := 0; i < len(b); i++ {
		find := b[i]
		isExist := false
		for _, v := range a {
			if find == v {
				isExist = true
				break
			}
		}
		if isExist {
			fmt.Printf("find number %d in slice\n", find)
		} else {
			fmt.Printf("not find number %d in slice\n", find)
		}
	}

}
