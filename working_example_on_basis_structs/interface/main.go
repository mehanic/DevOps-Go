package main

import (
	fmt "fmt"
	"math"
	"reflect"
)

//geometry interface definition
type geometry interface {
	Area() float64
	Perimeter() float64
}

//  circle struct and function receivers
type circle struct {
	Radius float64
}

func (c circle) Area() float64 {
	return math.Pi * math.Pow(c.Radius, 2)
}
func (c circle) Perimeter() float64 {
	return 2 * math.Pi * c.Radius
}

// rectangle struct and function receivers
type rectangle struct {
	Width  float64
	Height float64
}

func (r rectangle) Area() float64 {
	return r.Width * r.Height
}
func (r rectangle) Perimeter() float64 {
	return 2*r.Width + 2*r.Height
}

func measure(g geometry) {

	shape := reflect.TypeOf(g).Name()
	fmt.Printf("%s dimensios are: %g \n", shape, g)
	fmt.Println("Area is :", g.Area())
	fmt.Println("Perimeter is :", g.Perimeter())

}

func main() {
	myCircle := circle{Radius: 3.4}
	myRectangle := rectangle{3, 5}
	fmt.Println(myCircle.Area())
	measure(myCircle)
	measure(myRectangle)
}


// 36.316811075498
// circle dimensios are: {3.4} 
// Area is : 36.316811075498
// Perimeter is : 21.362830044410593
// rectangle dimensios are: {3 5} 
// Area is : 15
// Perimeter is : 16
