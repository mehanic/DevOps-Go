package main

import (
	"fmt"

	"github.com/mitchellh/mapstructure"
)

type Point struct {
	X int `mapstructure:"xx"`
	Y int `mapstructure:"yy"`
}

func (p Point) method() {
	fmt.Println(p.X)
	fmt.Println(p.Y)
}

func main() {
	pointsMap := map[string]Point{
		"b": {
			Y: 1,
			X: 2,
		},
	} //фигурные скобки в конце означает что мы этот тип инициализировали
	// otherMap := map[string]int{}
	//otherMap2 := make(map[string]int) //если make то фигурные скобки не нужны

	otherMap2 := make(map[string]Point)
	pointsMap["a"] = Point{X: 1, Y: 2}
	fmt.Println(pointsMap)

	fmt.Println("------------------------")
	otherMap2["b"] = Point{X: 1, Y: 2}
	//var otherMap3 map[string]Point //объявли но не пронициализировали
	//otherMap3["b"] = Point{X: 1, Y: 2}
	otherMap2["a"].method()
	value, ok := otherMap2["b"]

	if ok {
		fmt.Println(value)
	} else {
		fmt.Println("no key")
	}
	fmt.Println("------------------------")
	otherMap5 := map[string]int{
		"xx": 5,
		"yy": 5,
	}
	p1 := Point{}
	mapstructure.Decode(otherMap5, &p1)
	fmt.Println(p1)

	for key, value := range otherMap5 {
		fmt.Println(key, value)
	}
}

// map[a:{1 2} b:{2 1}]
// 0
// 0
// {1 2}
// {5 5}
// xx 5
// yy 5
