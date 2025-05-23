package main

import "fmt"

const pi = 3.14

const (
	hello = "Привет"
	e     = 2.718
)

const (
	zero  = iota
	_     // Пустая переменная, пропуск iota
	three // = 3
)

const (
	// нетипизированная константа
	year = 2017
	// типизированная константа
	yearType int = 2017
)

func main() {
	fmt.Println(hello)
	fmt.Println(three)

	var month int32 = 13
	// можно
	fmt.Println(year + month)
	// низя
	// invalid operation: yearType + month (mismatched types int and int32)
	// fmt.Println(yearType + month)
}

// Привет
// 2
// 2030
