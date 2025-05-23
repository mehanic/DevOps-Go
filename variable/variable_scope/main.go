package main

import "fmt"

var packVar = "Package Scope"


func main() {

	{
		var blockVar = "Block Scope"
		fmt.Println(blockVar)
		var packVar = "Block Override Scope"
		fmt.Println(packVar)

	}

	var funcVar = "Func Scope"
	fmt.Println(funcVar)
	fmt.Println(packVar)

	packVar = "Block Override Scope"
	fmt.Println(packVar)
	{
		fmt.Println(packVar)
	}


}

//OUTPUT
/*
Block Scope
Block Override Scope
Func Scope
Package Scope
Block Override Scope
Block Override Scope
*/
