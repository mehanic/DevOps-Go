package main

import (
	"fmt"
	"log"
	"math"
	"strconv"
)

func main() {
	var r, cX, cY, x, y string
	//r - radius of the circle
//cX, cY - coordinates of the center of the circle
//x, y - coordinates of any given point
//determine whether the point (x, y) is inside the circle or on its boundary
//using Euclidean distance

	fmt.Print("r=")
	fmt.Scan(&r)
	fmt.Print("cx=")
	fmt.Scan(&cX)
	fmt.Print("cy=")
	fmt.Scan(&cY)
	fmt.Print("x=")
	fmt.Scan(&x)
	fmt.Print("y=")
	fmt.Scan(&y)
	rInt, err := strconv.Atoi(r)
	if err != nil {
		log.Fatal(err)
	}
	cXInt, err := strconv.Atoi(cX)
	if err != nil {
		log.Fatal(err)
	}
	cYInt, err := strconv.Atoi(cY)
	if err != nil {
		log.Fatal(err)
	}
	xInt, err := strconv.Atoi(x)
	if err != nil {
		log.Fatal(err)
	}
	yInt, err := strconv.Atoi(y)
	if err != nil {
		log.Fatal(err)
	}
	part1 := math.Pow(float64(xInt-cXInt), 2.0)
	part2 := math.Pow(float64(yInt-cYInt), 2.0)
	d := math.Sqrt(part1 + part2)
	if d <= float64(rInt) {
		fmt.Println("correct")
	} else {
		fmt.Println("not correct")
	}
}

// ─λ go run 6.go                                            0 (0.001s) < 17:41:23
// r=4
// cx=5
// cy=3
// x=6
// y=3
// correct
