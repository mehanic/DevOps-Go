package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	score, valid := 5, true

	if score > 3 && valid {
		fmt.Println("good")
	}
}

func main1() {
	score, valid := 3, true

	if score > 3 && valid {
		fmt.Println("good")
	} else {
		fmt.Println("low")
	}
}

func main3() {
	score := 3

	if score > 3 {
		fmt.Println("good")
	} else if score == 3 {
		fmt.Println("on the edge")
	} else {
		fmt.Println("low")
	}
}

func main4() {
	score := 2

	if score > 3 {
		fmt.Println("good")
	} else if score == 3 {
		fmt.Println("on the edge")
	} else if score == 2 {
		fmt.Println("meh...")
	} else {
		fmt.Println("low")
	}
}

const usage = `
Feet to Meters
--------------
This program converts feet into meters.

Usage:
feet [feetsToConvert]`

func main5() {
	if len(os.Args) < 2 {
		fmt.Println(strings.TrimSpace(usage))

		// ALTERNATIVE:
		// fmt.Println("Please tell me a value in feet")

		return
	}

	arg := os.Args[1]

	feet, _ := strconv.ParseFloat(arg, 64)

	meters := feet * 0.3048

	fmt.Printf("%g feet is %g meters.\n", feet, meters)
}

// ------------------------
func main6() {
	args := os.Args

	if len(args) != 3 {
		fmt.Println("Usage: [username] [password]")
		return
	}

	u, p := args[1], args[2]

	if u != "jack" {
		fmt.Printf("Access denied for %q.\n", u)
	} else if p != "1888" {
		fmt.Printf("Invalid password for %q.\n", u)
	} else {
		fmt.Printf("Access granted to %q.\n", u)
	}
}

// -----------------------------------
const (
	usage1   = "Usage: [username] [password]"
	errUser  = "Access denied for %q.\n"
	errPwd   = "Invalid password for %q.\n"
	accessOK = "Access granted to %q.\n"
	user     = "jack"
	pass     = "1888"
)

func main7() {
	args := os.Args

	if len(args) != 3 {
		fmt.Println(usage1)
		return
	}

	u, p := args[1], args[2]

	if u != user {
		fmt.Printf(errUser, u)
	} else if p != pass {
		fmt.Printf(errPwd, u)
	} else {
		fmt.Printf(accessOK, u)
	}
}

//---------------------------

const (
	usage0       = "Usage: [username] [password]"
	errUser0     = "Access denied for %q.\n"
	errPwd0      = "Invalid password for %q.\n"
	accessOK0    = "Access granted to %q.\n"
	user0, user2 = "jack", "inanc"
	pass0, pass2 = "1888", "1879"
)

func main8() {
	args := os.Args

	if len(args) != 3 {
		fmt.Println(usage)
		return
	}

	u, p := args[1], args[2]

	if u != user && u != user2 {
		fmt.Printf(errUser, u)
	} else if u == user && p == pass {
		fmt.Printf(accessOK, u)
	} else if u == user2 && p == pass2 {
		fmt.Printf(accessOK, u)
	} else {
		fmt.Printf(errPwd, u)
	}
}
