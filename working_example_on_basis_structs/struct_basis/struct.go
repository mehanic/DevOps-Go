package main

import "fmt"

// в golang нет классов, есть структуры
type Point struct {
	X int //с большой отвечает за область видимости переменных
	Y int
	S string
}

func (p Point) method() {
	fmt.Println(p.X)
	fmt.Println(p.Y)
}

func main() {
	structs()
	p1 := Point{
		X: 20,
		Y: 20,
		S: "dauren",
	}
	p2 := &p1
	p2.method()
}

func structs() {
	p1 := Point{
		X: 0,
		Y: 0,
		S: "dauren",
	}
	fmt.Println(p1)
	fmt.Println(p1.X)
	fmt.Println(p1.Y)
	p1.X = 123
	p1.Y = 123

	p2 := Point{X: 3}
	fmt.Println(p1)
	fmt.Println(p2)
	g := &p1
	fmt.Println(g.X)
	fmt.Println(&g)
	fmt.Println(g)
	fmt.Println(*g)
}

// {0 0 dauren}
// 0
// 0
// {123 123 dauren}
// {3 0 }
// 123
// 0xc00005a028
// &{123 123 dauren}
// {123 123 dauren}
// 20
// 20
