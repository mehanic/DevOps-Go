package main

import (
	"fmt"
)

func main() {
	firstName, lastName := fullName()
	fmt.Println(firstName) 
	fmt.Println(lastName)  
	fmt.Println(firstName, lastName)
}

func fullName() (string, string) {
	return "Yoso", "Prasetiyo"
}

func main1() {
	luas, keliling := calculate(10, 4)
	fmt.Println("Area : ")
	fmt.Println(luas)
	fmt.Println("Perimeter : ")
	fmt.Println(keliling)
}

func calculate(length int, width int) (int, int) {
	area := length * width
	perimeter := 2 * (length + width)
	return area, perimeter
}
