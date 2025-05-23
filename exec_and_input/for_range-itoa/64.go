package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	//strconv.Itoa-> numbers to string
	//strconv.Atoi-> string to numbers
	//рекурсию
	n := 2100
	primeNumbers := []string{}
	for i := 2; i < n; i++ {
		counter := 0
		for j := 2; j <= i; j++ {
			if i%j == 0 {
				counter += 1
			}
		}
		if counter == 1 {
			primeNumbers = append(primeNumbers, strconv.Itoa(i))
		}
	}
	s := strings.Join(primeNumbers, "")
	var c int
	arr := []int{}
	fmt.Scan(&c)
	for i := 0; i < c; i++ {
		var d int
		fmt.Scan(&d)
		arr = append(arr, d)
	}
	for _, v := range arr {
		fmt.Print(string(s[v-1]))
	}
}
