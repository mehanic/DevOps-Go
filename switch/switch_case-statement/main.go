package main

import "fmt"

func main() {

	fmt.Println()

	x := "C"
	switch x {
	case "A":
		fmt.Println("A")
	case "B", "C": // multiple
		fmt.Println("multiple")
	case "D":
		fmt.Println("D")
		fallthrough // no default fallthrough in go, so no need break statement in switch
	default:
		fmt.Println("default")
	}

	fmt.Println()

	switch { // will choose the first true condition
	case 1 == 2:
		fmt.Println("A")
	case 3 == 3:
		fmt.Println("B")
	}

	fmt.Println()

	switchOnType(5)

	fmt.Println()
}

func switchOnType(x interface{}) {

	switch x.(type) { // this is an assert, asserting, "x is of this type"
	case int:
		fmt.Println("x is int")
	case string:
		fmt.Println("x is string")
	default:
		fmt.Println("What x is ?")
	}
}
