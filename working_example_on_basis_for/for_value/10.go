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
	for i := 0; i < n; i++ {
		k := rand.Intn(100)
		a = append(a, k)
	}
	fmt.Println(a)
	find := 5
	isExist := false
	for _, v := range a {
		if find == v {
			isExist = true
			break
		}
	}
	if isExist {
		fmt.Println("find")
	} else {
		fmt.Println("not find")
	}
}
