package main

import "fmt"

type Point struct {
	X, Y int
}

// func movePoint(p Point, x, y int) Point {
// 	p.X += x
// 	p.Y += y
// 	return p

// }

//	func movePointPtr(p *Point, x, y int) {
//		p.X += x
//		p.Y += y
//	}
//
// методы структуры
func (p Point) movePoint(x, y int) Point {
	p.X += x
	p.Y += y
	return p

}
func (p *Point) movePointPtr(x, y int) {
	p.X += x
	p.Y += y
}

func main() {
	p := Point{
		1,
		2,
	}
	fmt.Println(p)
	fmt.Println(p.movePoint(1, 1))
	fmt.Println(p)
	p.movePointPtr(1, 2)
	fmt.Println(p)
	v := &p
	v.movePoint(1, 2)
	fmt.Println(p)
	v.movePointPtr(2, 2)
	fmt.Println(p)
}

// ─λ go run methods.go                                                                                                                                                0 (0.001s) < 12:57:07
// {1 2}
// {2 3}
// {1 2}
// {2 4}
// {2 4}
// {4 6}

// {1 2}
// {2 3}
// {1 2}
// {2 4}
// {2 4}
// {4 6}