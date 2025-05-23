package main

import (
	"fmt"
	"strconv"
)

func main() {

	randInt := 5
	randFloat := 10.5
	randString := "100"
	randString2 := "250.5"

	fmt.Println(float64(randInt))
	fmt.Println(int(randFloat))

	newInt, _ := strconv.ParseInt(randString, 0, 64)
	fmt.Println(newInt)

	newFloat, _ := strconv.ParseFloat(randString2, 64)
	fmt.Println(newFloat)

}

// ╰─λ go run TypeCasting1.go                                                                                                       0 (0.001s) < 14:39:25
// 5
// 10
// 100
// 250.5
