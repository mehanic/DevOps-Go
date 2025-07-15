package main

import (
	"fmt"
)

func main() {
	fmt.Println("Tutorial for interfaces")

	// Creating instances of rectangle and triangle
	r := rectangle{4.5, 6}
	t := triangle{10, 5}

	// Storing pointers to these instances in a slice of shapes interface type
	sh := []shapes{r, t}

	// Printing the area of each shape
	fmt.Println("Area of rectangle with sides", r.length, ",", r.breadth, "is", sh[0].area())
	fmt.Println("Area of triangle with height and base", t.height, ",", t.base, "is", sh[1].area())
}

// Defining the shapes interface with an area method
type shapes interface {
	area() float32
}

// Defining the rectangle struct
type rectangle struct {
	length  float32
	breadth float32
}

// Implementing the area method for the rectangle struct
func (r rectangle) area() float32 {
	return r.breadth * r.length
}

// Defining the triangle struct
type triangle struct {
	height float32
	base   float32
}

// Implementing the area method for the triangle struct
func (t triangle) area() float32 {
	return (1.0 / 2.0) * t.base * t.height
}


// Tutorial for interfaces
// Area of rectangle with sides 4.5 , 6 is 27
// Area of triangle with height and base 10 , 5 is 25
