package main

import "fmt"

func main() {
	fmt.Println(fullName("Yoso", "Prasetiyo"))
}

func fullName(firstName string, lastName string) string {
	fullName := firstName + " " + lastName
	return fullName
}

func main1() {
	fmt.Println(fullName1("Ivan", "Mazepa"))
}

func fullName1(firstName string, lastName string) string {
	if firstName == "" || lastName == "" {
		return "First Name or Last Name is Empty"
	} else {
		return firstName + " " + lastName
	}
}

func main2() {
	totalScore := askQuestion2(1, "mmwk7a")
	totalScore += askQuestion2(2, "kmzway87aa")
	totalScore += askQuestion2(3, "bkbmanta8")
	fmt.Println("Your Score: ", totalScore)
}

func askQuestion2(number int, question string) int {
	var input string
	fmt.Printf("[Question %d] : Type the following code: %s \n", number, question)
	fmt.Scan(&input)
	fmt.Printf("%s has been entered \n", input)
	fmt.Println("---------------------------------------")

	return 100
}

func main3() {
	totalScore := askQuestion3(1, "mmwk7a")
	totalScore += askQuestion3(2, "kmzway87aa")
	totalScore += askQuestion3(3, "bkbmanta8")
	fmt.Println("Your Score: ", totalScore)
}

func askQuestion3(number int, question string) int {
	var input string
	fmt.Printf("[Question %d] : Type the following code: %s \n", number, question)
	fmt.Scan(&input)

	if question == input {
		fmt.Println("Correct!")
		return 10
	} else {
		fmt.Println("Wrong!")
		return 0
	}
}
