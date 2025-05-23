package main

import "fmt"

func main() {
	var x, a, y float64

	fmt.Print("Enter the weight of the candy x=:")
	fmt.Scanln(&x)
	fmt.Printf("Enter the price of %f kg of candy a =:", x)
	fmt.Scanln(&a)
	fmt.Print("Enter how many kilograms of candy you want to buy a =:")
	fmt.Scanln(&y)
	fmt.Printf("The price of 1kg of candy is %f som, so %fkg of candy will cost %f som\n", a/x, y, a/x*y)

}
