package main

import "fmt"

func main() {

	// Print the menu of days
	fmt.Println("1. Monday \n2. Tuesday \n3. Wednesday \n4. Thursday \n5. Friday \n6. Saturday \n7. Sunday")

	// Declare a variable to store the user's choice
	var choice int

	// Prompt the user for input
	fmt.Println("Enter a number (1-7) to choose a day:")
	fmt.Scanf("%d", &choice)

	// Use a switch statement to map the choice to a day
	switch choice {
	case 1:
		fmt.Println("Monday")
	case 2:
		fmt.Println("Tuesday")
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
		fmt.Println("Invalid choice! Please enter a number between 1 and 7.")
	}
}


// 1. Monday 
// 2. Tuesday 
// 3. Wednesday 
// 4. Thursday 
// 5. Friday 
// 6. Saturday 
// 7. Sunday
// Enter a number (1-7) to choose a day:
// 2
// Tuesday