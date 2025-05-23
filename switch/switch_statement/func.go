package switchstatement

import "fmt"

func CreateSwitch(month int) {

	/*
		We can use the switch statement to select one of many code blocks to be executed.

		- The expression is evaluated once.
		- The value of the switch expression is compared with the values of each case.
		- If there is a match, the associated block of code is executed.
		- The default keyword is optional. It specifies some code to run if there is no case match.
	*/

	switch month {
	case 1:
		fmt.Println("It's January!")
	case 2:
		fmt.Println("It's February!")
	case 3:
		fmt.Println("It's March!")
	case 4:
		fmt.Println("It's April!")
	case 5:
		fmt.Println("It's May!")
	case 6:
		fmt.Println("It's June!")
	case 7:
		fmt.Println("It's July!")
	case 8:
		fmt.Println("It's August!")
	case 9:
		fmt.Println("It's September!")
	case 10:
		fmt.Println("It's October!")
	case 11:
		fmt.Println("It's November!")
	case 12:
		fmt.Println("It's December!")
	default:
		fmt.Println("Not a month!") // The default keyword specifies some code to run if there is no case match.
	}

	// All the case values should have the same type as the switch expression. Otherwise, the compiler will raise an error.

}
