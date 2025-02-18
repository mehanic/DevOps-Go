package conditionals

import "fmt"

func CheckCondition() {

	// Let's say we manage a pub and we want to check ages of incoming customers

	allowed_age := 18

	customer1 := 15
	customer2 := 29
	customer3 := 18

	// Normally we can put them to a list but for now we keep it simple.

	if customer1 > allowed_age {
		fmt.Println(fmt.Sprintf("%v", customer1) + " You can enter, have fun!")
	} else if customer1 == allowed_age {
		fmt.Println("Oh, yo're just 18. But you can enter, have fun!")
	} else {
		fmt.Println(fmt.Sprintf("%v", customer1) + " You can't enter, sorry for that:(")
	}

	if customer2 > allowed_age {
		fmt.Println(fmt.Sprintf("%v", customer2) + " You can enter, have fun!")
	} else if customer2 == allowed_age {
		fmt.Println("Oh, yo're just 18. But you can enter, have fun!")
	} else {
		fmt.Println(fmt.Sprintf("%v", customer2) + " You can't enter, sorry for that:(")
	}

	if customer3 > allowed_age {
		fmt.Println(fmt.Sprintf("%v", customer3) + " You can enter, have fun!")
	} else if customer3 == allowed_age {
		fmt.Println("Oh, yo're just 18. But you can enter, have fun!")
	} else {
		fmt.Println(fmt.Sprintf("%v", customer3) + " You can't enter, sorry for that:(")
	}

}
