package main

import "fmt"

// 簡短的宣告變數只能在方法中，在方法外要用一般宣告
var e string

func main() {

	a := 10
	b := "golang"
	c := 4.17
	d := true

	fmt.Printf("%v \n", a)
	fmt.Printf("%v \n", b)
	fmt.Printf("%v \n", c)
	fmt.Printf("%v \n", d)

	fmt.Printf("%T \n", a)
	fmt.Printf("%T \n", b)
	fmt.Printf("%T \n", c)
	fmt.Printf("%T \n", d)

	// 方法內也可以用一般宣告
	var f bool = false
	fmt.Println(f)
}

// 10
// golang
// 4.17
// true
// int
// string
// float64
// bool
// false
