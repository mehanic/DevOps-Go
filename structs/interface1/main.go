package main

import (
	"fmt"
	"math"
)

type square struct {
	side float64
}

type circle struct {
	radius float64
}

func (x square) area() float64 {
	return x.side * x.side
}

func (y circle) area() float64 {
	return math.Pi * y.radius * y.radius
}

type shape interface {
	area() float64
}

func info(sh shape) {
	fmt.Println(sh)
	fmt.Println(sh.area())
}

func main() {
	p := square{4}
	q := circle{5}
	info(p)
	info(q)
}

// {4}
// 16
// {5}
// 78.53981633974483
