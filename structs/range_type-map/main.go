
package main

import (
	"fmt"

	"github.com/mitchellh/mapstructure"
)

type Point struct { // Тип Структура Point
	X int // Поля структуры
	Y int // Поля структуры
}

func (p Point) method() { // Метод Структуры Point
	fmt.Printf("call Point method p.X = %d\n", p.X)
}

func main() {
	// Итератор for range
	arr := []string{"a", "b", "c"}
	for i, l := range arr {
		fmt.Printf("arr[%d] = %s\n", i, l)
	}

	for _, l2 := range arr {
		fmt.Printf("l2 = %s\n", l2)
	}

	for i2 := range arr {
		fmt.Printf("i2 = %d\n", i2)
	}

	// Структура map
	//pointsMap := map[string]Point{} // Объявили map и проинициализировали её как nil
	pointsMap := map[string]Point{ // Объявили map и проинициализировали
		"b": {
			X: 13,
			Y: 15,
		},
	}
	//otherMap := map[string]int{} // Объявили map и проинициализировали её как nil
	otherMap := map[string]int{ // Объявили map и проинициализировали
		"x": 23,
		"y": 25,
	}
	otherPointsMap := make(map[int]Point) // Объявили map и проинициализировали
	var oneMorePointsMap map[string]Point // Объявили map, без инициализаций
	pointsMap["a"] = Point{
		X: 1,
		Y: 2,
	}
	fmt.Println(pointsMap)
	fmt.Println(pointsMap["a"])
	fmt.Println(pointsMap["a"].X)
	fmt.Println(pointsMap["a"].Y)
	pointsMap["a"].method()

	fmt.Println(otherMap)
	fmt.Println(otherMap["x"], otherMap["y"])
	p1 := Point{}
	mapstructure.Decode(otherMap, &p1) // Декодируем в структуру Point
	fmt.Println("Декодируем в структуру Point:", p1)
	fmt.Println(p1.X, p1.Y)
	for k, v := range otherMap { // Итерация Map
		fmt.Printf("key=%s value=%d\n", k, v)
	}

	otherPointsMap[1] = Point{X: 1, Y: 2}
	fmt.Println(otherPointsMap[1])

	if oneMorePointsMap == nil {
		//oneMorePointsMap = map[string]Point{} // Проинициализировали её как nil
		oneMorePointsMap = map[string]Point{
			"a": {Y: 1, X: 2},
		}
	}
	fmt.Println(oneMorePointsMap)

	// Проверяем наличия ключа в Map
	key := "a"
	value, ok := oneMorePointsMap[key]
	if ok {
		fmt.Printf("key=%s exist in map\n", key)
		fmt.Println(value)
	} else {
		fmt.Printf("key=%s does exist in map\n", key)
		fmt.Println(value)
	}
}


// arr[0] = a
// arr[1] = b
// arr[2] = c
// l2 = a
// l2 = b
// l2 = c
// i2 = 0
// i2 = 1
// i2 = 2
// map[a:{1 2} b:{13 15}]
// {1 2}
// 1
// 2
// call Point method p.X = 1
// map[x:23 y:25]
// 23 25
// Декодируем в структуру Point: {23 25}
// 23 25
// key=x value=23
// key=y value=25
// {1 2}
// map[a:{2 1}]
// key=a exist in map
// {2 1}
