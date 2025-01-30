package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"
)

func main() {
	
	args := os.Args
	if len(args) != 2 {
		fmt.Println("Give me the magnitude of the earthquake.")
		return
	}

	richter, err := strconv.ParseFloat(args[1], 64)
	if err != nil {
		fmt.Println("I couldn't get that, sorry.")
		return
	}

	var desc string

	switch r := richter; {
	case r < 2:
		desc = "micro"
	case r >= 2 && r < 3:
		desc = "very minor"
	case r >= 3 && r < 4:
		desc = "minor"
	case r >= 4 && r < 5:
		desc = "light"
	case r >= 5 && r < 6:
		desc = "moderate"
	case r >= 6 && r < 7:
		desc = "strong"
	case r >= 7 && r < 8:
		desc = "major"
	case r >= 8 && r < 10:
		desc = "great"
	default:
		desc = "massive"
	}

	fmt.Printf("%.2f is %s\n", richter, desc)
}

func main1() {
	args := os.Args
	if len(args) != 2 {
		fmt.Println("Tell me the magnitude of the earthquake in human terms.")
		return
	}

	var richter string

	desc := args[1]

	switch desc {
	case "micro":
		richter = "less than 2.0"
	case "very minor":
		richter = "2 - 2.9"
	case "minor":
		richter = "3 - 3.9"
	case "light":
		richter = "4 - 4.9"
	case "moderate":
		richter = "5 - 5.9"
	case "strong":
		richter = "6 - 6.9"
	case "major":
		richter = "7 - 7.9"
	case "great":
		richter = "8 - 9.9"
	case "massive":
		richter = "10+"
	default:
		richter = "unknown"
	}

	fmt.Printf(
		"%s's richter scale is %s\n",
		desc, richter,
	)
}

//--

const (
	usage       = "Usage: [username] [password]"
	errUser     = "Access denied for %q.\n"
	errPwd      = "Invalid password for %q.\n"
	accessOK    = "Access granted to %q.\n"
	user, user2 = "jack", "inanc"
	pass, pass2 = "1888", "1879"
)

func main2() {
	args := os.Args

	if len(args) != 3 {
		fmt.Println(usage)
		return
	}

	u, p := args[1], args[2]

	// More readable, right? 👍
	switch {
	case u != user && u != user2:
		fmt.Printf(errUser, u)
	case u == user && p == pass:
		// notice this one (no more duplication)
		fallthrough
	case u == user2 && p == pass2:
		fmt.Printf(accessOK, u)
	default:
		fmt.Printf(errPwd, u)
	}
}

//--

const usage0 = `[command] [string]

Available commands: lower, upper and title`

func main3() {
	args := os.Args
	if len(args) != 3 {
		fmt.Println(usage0)
		return
	}

	cmd, str := args[1], args[2]
	switch cmd {
	case "lower":
		fmt.Println(strings.ToLower(str))
	case "upper":
		fmt.Println(strings.ToUpper(str))
	case "title":
		fmt.Println(strings.Title(str))
	default:
		fmt.Printf("Unknown command: %q\n", cmd)
	}
}

//--

func main4() {
	if len(os.Args) != 2 {
		fmt.Println("Give me a month name")
		return
	}

	year := time.Now().Year()
	leap := year%4 == 0 && (year%100 != 0 || year%400 == 0)

	days, month := 28, os.Args[1]

	switch strings.ToLower(month) {
	case "april", "june", "september", "november":
		days = 30
	case "january", "march", "may", "july",
		"august", "october", "december":
		days = 31
	case "february":
		if leap {
			days = 29
		}
	default:
		fmt.Printf("%q is not a month.\n", month)
		return
	}

	fmt.Printf("%q has %d days.\n", month, days)
}


// $ go run main.go 5.7
// 5.70 is moderate

// $ go run main1.go light
// light's richter scale is 4 - 4.9

// $ go run main2.go jack 1888
// Access granted to jack.

// $ go run main3.go lower "Hello World"
// hello world

// $ go run main4.go February
// "February" has 28 days. (In a non-leap year)
