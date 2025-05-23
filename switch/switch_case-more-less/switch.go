package main

import "fmt"

func main() {
	day := 2
	switch day {
	case 1:
		fmt.Println("Monday")
	case 2:
		fmt.Println("Tuesday", day)
	case 3:
		fmt.Println("Wednesday")
	case 4:
		fmt.Println("Thursday")
	case 5:
		fmt.Println("Friday")
	case 6:
		fmt.Println("Saturday")
	case 7:
		fmt.Println("Sunday")
	default:
		fmt.Println("Airway")
	}

	var marks int = 76
	switch {
	case marks >= 50:
		fmt.Println("You got more than 50 so your grade is C")
	case marks >= 70:
		fmt.Println("You got more than 70 so your grade is B")
	case marks >= 90:
		fmt.Println("You got more than 90 so your grade is A")
	default:
		fmt.Println("The marks which enter is not valid marks")
	}
}
