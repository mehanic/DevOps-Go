package main

import "fmt"

func main() {
	var x, a, y, b float64

	fmt.Print("Enter the weight of the chocolate x=:")
	fmt.Scanln(&x)
	fmt.Printf("Enter the price of %f kg of chocolate a =:", x)
	fmt.Scanln(&a)
	fmt.Print("Enter the weight of the candy y=:")
	fmt.Scanln(&y)
	fmt.Printf("Enter the price of %f kg of candy a =:", y)
	fmt.Scanln(&b)
	fmt.Printf("1kg of chocolate is %f som more expensive than 1kg of candy\n", a/x-b/y)

}
