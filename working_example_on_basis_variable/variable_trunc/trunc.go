package main

import (
	"fmt"
	"math"
)

func main() {
	var num float64
	fmt.Printf("enter a float type number\n")
	fmt.Scanf("%f", &num)
	res := math.Trunc(num)
	fmt.Println(res)
}
