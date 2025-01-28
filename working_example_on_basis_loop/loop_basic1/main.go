package main

import "fmt"

func main() {
	fmt.Println("Welcome to loops in GoLang")

	days := []string{"Sunday", "Tuesday", "Wednesday", "Friday", "Saturday"}

	fmt.Println(days)

	fmt.Println("Loops example 1:")

	for d := 0; d < len(days); d++ {
		fmt.Println(days[d])
	}

	fmt.Println("Loops example 2:")

	for i := range days {
		fmt.Println(days[i])
	}

	fmt.Println("Loops example 3:")

	for index, day := range days {
		fmt.Printf("Index is %v and value is %v\n", index, day)
	}

	rogueValue := 1

	for rogueValue < 10 {

		if rogueValue == 2 {
			goto lco
		}

		if rogueValue == 5 {
			break
		}
		fmt.Println("Value is: ", rogueValue)
		rogueValue++
	}

lco:
	fmt.Println("Jumping at LearnCodeOnline")

	anotherValue := 1
	for anotherValue < 10 {
		if anotherValue == 5 {
			anotherValue++
			continue
		}
		fmt.Println("Value is: ", anotherValue)
		anotherValue++
	}
}


// Welcome to loops in GoLang
// [Sunday Tuesday Wednesday Friday Saturday]
// Loops example 1:
// Sunday
// Tuesday
// Wednesday
// Friday
// Saturday
// Loops example 2:
// Sunday
// Tuesday
// Wednesday
// Friday
// Saturday
// Loops example 3:
// Index is 0 and value is Sunday
// Index is 1 and value is Tuesday
// Index is 2 and value is Wednesday
// Index is 3 and value is Friday
// Index is 4 and value is Saturday
// Value is:  1
// Jumping at LearnCodeOnline
// Value is:  1
// Value is:  2
// Value is:  3
// Value is:  4
// Value is:  6
// Value is:  7
// Value is:  8
// Value is:  9
