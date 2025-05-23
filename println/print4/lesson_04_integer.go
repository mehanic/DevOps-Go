package main

import (
	"fmt"
	//"math"
)

func main () {
	var defaultInt int64
	intVar := 10

	fmt.Println("2 + 3 = ", 2 + 3)
	fmt.Println("2 - 3 = ", 2 - 3)
	fmt.Println("2 * 3 = ", 2 * 3)
	fmt.Println("int(2 * 3) = ", int(2 * 3))
	fmt.Println("defaultInt = ", defaultInt)

	intVar --;	fmt.Println("intVar :" , intVar)
	intVar ++;	fmt.Println("intVar :" , intVar)
	intVar += 10;	fmt.Println("intVar :" , intVar)
	intVar -= 5;	fmt.Println("intVar :" , intVar)
	intVar *= 3;	fmt.Println("intVar :" , intVar)
	intVar /= 7;	fmt.Println("intVar :" , intVar)
}
