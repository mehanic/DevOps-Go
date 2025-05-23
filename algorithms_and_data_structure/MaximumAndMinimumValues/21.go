package main

import "fmt"

func main() {
	a := 100
	b := 500
	c := 1000
	maxi := 0
	mini := 0
	if a > b && a > c {
		maxi = a
	} else if b > a && b > c {
		maxi = b
	} else {
		maxi = c
	}
	if a < b && a < c {
		mini = a
	} else if b < a && b < c {
		mini = b
	} else {
		mini = c
	}
	fmt.Println(maxi - mini)
}
