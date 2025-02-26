package main

import (
	"fmt"
)

func main() {
	//Created Variables
	// var a = 10
	// investmentAmount, years := 1000, 5.5
	// const inflationRate = 6.5 //Can't changed thereafter
	var investmentAmount float64
	var years float64
	expectedReturnRate := 5.5 //Shortcut of Autosetting of variable type - Only works when it is inside a function

	outputText("Investment Amount: ")
	//Scans the terminal for input
	fmt.Scan(&investmentAmount) //Using pointers

	fmt.Print("Expected Return Rate: ")
	fmt.Scan(&expectedReturnRate)

	fmt.Print("Years: ")
	fmt.Scan(&years)

}

func outputText(text string) {
	fmt.Print(text)
}
