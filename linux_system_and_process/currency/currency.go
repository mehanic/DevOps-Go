package main

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
)

func main() {
	// start with our user input
	// of fifteen dollars and 93 cents
	userInput := "15.93"

	pennies, err := ConvertStringDollarsToPennies(userInput)
	if err != nil {
		panic(err)
	}

	fmt.Printf("User input converted to %d pennies\n", pennies)

	// adding 15 cents
	pennies += 15

	dollars := ConvertPenniesToDollarString(pennies)

	fmt.Printf("Added 15 cents, new values is %s dollars\n", dollars)
}

// ConvertStringDollarsToPennies takes a dollar amount
// as a string, i.e. 1.00, 55.12 etc and converts it
// into an int64
func ConvertStringDollarsToPennies(amount string) (int64, error) {
	// check if amount can convert to a valid float
	_, err := strconv.ParseFloat(amount, 64)
	if err != nil {
		return 0, err
	}

	// split the value on "."
	groups := strings.Split(amount, ".")

	// if there is no . result will still be
	// captured here
	result := groups[0]

	// base string
	r := ""

	// handle the data after the "."
	if len(groups) == 2 {
		if len(groups[1]) != 2 {
			return 0, errors.New("invalid cents")
		}
		r = groups[1]
	}

	// pad with 0, this will be
	// 2 0's if there was no .
	for len(r) < 2 {
		r += "0"
	}

	result += r

	// convert it to an int
	return strconv.ParseInt(result, 10, 64)
}

// ---
// ConvertPenniesToDollarString takes a penny amount as
// an int64 and returns a dollar string representation
func ConvertPenniesToDollarString(amount int64) string {
	// parse the pennies as a base 10 int
	result := strconv.FormatInt(amount, 10)

	// check if negative, will set it back later
	negative := false
	if result[0] == '-' {
		result = result[1:]
		negative = true
	}

	// left pad with 0 if we're passed in value < 100
	for len(result) < 3 {
		result = "0" + result
	}
	length := len(result)

	// add in the decimal
	result = result[0:length-2] + "." + result[length-2:]

	// from the negative we stored earlier!
	if negative {
		result = "-" + result
	}

	return result
}
