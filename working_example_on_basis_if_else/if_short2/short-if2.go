package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"
)

func main() {
	if len(os.Args) != 2 {
		fmt.Println("Pick a number")
		return
	}

	n, err := strconv.Atoi(os.Args[1])
	if err != nil {
		fmt.Printf("%q is not a number\n", os.Args[1])
		return
	}

	if n%8 == 0 {
		fmt.Printf("%d is an even number and it's divisible by 8\n", n)
	} else if n%2 == 0 {
		fmt.Printf("%d is an even number\n", n)
	} else {
		fmt.Printf("%d is an odd number\n", n)
	}
}

func main1() {
	if len(os.Args) != 2 {
		fmt.Println("Give me a year number")
		return
	}

	year, err := strconv.Atoi(os.Args[1])
	if err != nil {
		fmt.Printf("%q is not a valid year.\n", os.Args[1])
		return
	}

	// Notice that:
	// I've intentionally created this solution as verbose
	// as I can.
	//
	// See the next exercise.

	var leap bool
	if year%400 == 0 {
		leap = true
	} else if year%100 == 0 {
		leap = false
	} else if year%4 == 0 {
		leap = true
	}

	if leap {
		fmt.Printf("%d is a leap year.\n", year)
	} else {
		fmt.Printf("%d is not a leap year.\n", year)
	}
}

func main3() {
	if len(os.Args) != 2 {
		fmt.Println("Give me a year number")
		return
	}

	year, err := strconv.Atoi(os.Args[1])
	if err != nil {
		fmt.Printf("%q is not a valid year.\n", os.Args[1])
		return
	}

	if year%4 == 0 && (year%100 != 0 || year%400 == 0) {
		fmt.Printf("%d is a leap year.\n", year)
	} else {
		fmt.Printf("%d is not a leap year.\n", year)
	}
}

func main4() {
	if len(os.Args) != 2 {
		fmt.Println("Give me a month name")
		return
	}

	// get the current year and find out whether it's a leap year
	year := time.Now().Year()
	leap := year%4 == 0 && (year%100 != 0 || year%400 == 0)

	// setting it to 28, saves me typing it below again
	days := 28

	month := os.Args[1]

	// case insensitive
	if m := strings.ToLower(month); m == "april" ||
		m == "june" ||
		m == "september" ||
		m == "november" {
		days = 30
	} else if m == "january" ||
		m == "march" ||
		m == "may" ||
		m == "july" ||
		m == "august" ||
		m == "october" ||
		m == "december" {
		days = 31
	} else if m == "february" {
		// try a "logical and operator" above.
		// like: `else if m == "february" && leap`
		if leap {
			days = 29
		}
	} else {
		fmt.Printf("%q is not a month.\n", month)
		return
	}

	fmt.Printf("%q has %d days.\n", month, days)
}
