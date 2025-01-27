package main

import (
	"fmt"
	"strconv"
)

// go doc fmt println
func main() {
	//
	s1 := "3"
	s2 := "4"
	//strconv
	//convert from string to int
	n1, _ := strconv.Atoi(s1)
	n2, _ := strconv.Atoi(s2)
	c := n1 + n2
	//convert from int to string
	s3 := strconv.Itoa(c)
	fmt.Println(s1 + s2 + s3)

	//func названиеФункции(параметры) возвращаемые_значения
	//возвращаемые_значения := библиотека.название_фукнции(значение_параметров)
}

// 347
