package main

import "fmt"

// 1)Пользователь вводит два числа. Одно присваивается одной переменной,
// а второе - другой. Необходимо поменять значения переменных так,
// чтобы значение первой оказалось во второй, а второй - в первой.
//
// a := 3
// b := 4
// fmt.Println(a,b)//3 4
// //ваш код
// fmt.Println(a,b)//4 3
func main() {
	a := 3
	b := 4
	fmt.Println(a, b)
	//1
	//c := a
	//a = b
	//b = c
	//2
	a, b = b, a
	fmt.Println(a, b)
}

// 3 4
// 4 3
