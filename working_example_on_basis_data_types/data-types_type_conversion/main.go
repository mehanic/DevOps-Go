package main

import "fmt"

func main() {

	x := 10
	y := 10.0

	fmt.Printf("%v %T\n", x, x)
	fmt.Printf("%v %T\n", y, y)
	/*
		10 int
		10 float64
	*/

	/*
		fmt.Println(x + y)
		\main.go:16:16: invalid operation: x + y (mismatched types int and float64)
	*/
	//Even both value are numbers, they are diffrent type so you cant operate them. int8 ve int16 bile toplanamaz.

	fmt.Println(x + int(y))
	fmt.Println(y + float64(x))

	var x1 int16 = 128
	var y1 int8
	y1 = int8(x1)

	fmt.Println(y1) // -128

	a := "100"
	b := 106
	fmt.Println(a + string(b))     //100
	fmt.Println(a + fmt.Sprint(b)) //1001

}

// 10 int
// 10 float64
// 20
// 20
// -128
// 100j
// 100106
