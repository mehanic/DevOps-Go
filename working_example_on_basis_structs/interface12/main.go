package main

import "fmt"

// the type who implements all func in Shape is also type Shape
type Shape interface {
	area() float64
}

type Square struct {
	side float64
}

// Square implemets func area, so Square is also type Shape
func (r Square) area() float64 {
	return r.side * r.side
}

func printShapeArea(z Shape) {
	fmt.Println(z.area())
}

func main() {
	// z is a Square
	z := Square{24}
	// pass Square as Shape
	printShapeArea(z)
}

//576