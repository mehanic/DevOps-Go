package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	age := 10

	if age > 60 {
		fmt.Println("Getting older")
	} else if age > 30 {
		fmt.Println("Getting wiser")
	} else if age > 20 {
		fmt.Println("Adulthood")
	} else if age > 10 {
		fmt.Println("Young blood")
	} else {
		fmt.Println("Booting up")
	}

	main1()
	main2()
	main3()
	main4()
	main5()
	main6()
}

func main1() {
	isSphere, radius := true, 200

	if isSphere && radius >= 200 {
		fmt.Println("It's a big sphere.")
	} else {
		fmt.Println("I don't know.")
	}
}

func main2() {
	var (
		args = os.Args
		l    = len(args) - 1
	)

	if l == 0 {
		fmt.Println("Give me args")
	} else if l == 1 {
		fmt.Printf("There is one: %q\n", args[1])
	} else if l == 2 {
		fmt.Printf(
			`There are two: "%s %s"`+"\n",
			args[1], args[2],
		)
	} else {
		fmt.Printf("There are %d arguments\n", l)
	}
}

func main3() {
	args := os.Args

	if len(args) != 2 || len(args[1]) != 1 {
		fmt.Println("Give me a letter")
		return
	}

	// I didn't use a short-if here because, it's already
	// hard to read. Do not make it harder.

	s := args[1]
	if s == "a" || s == "e" || s == "i" || s == "o" || s == "u" {
		fmt.Printf("%q is a vowel.\n", s)
	} else if s == "w" || s == "y" {
		fmt.Printf("%q is sometimes a vowel, sometimes not.\n", s)
	} else {
		fmt.Printf("%q is a consonant.\n", s)
	}
}

func main4() {
	args := os.Args

	if len(args) != 2 || len(args[1]) != 1 {
		fmt.Println("Give me a letter")
		return
	}

	s := args[1]
	if strings.IndexAny(s, "aeiou") != -1 {
		fmt.Printf("%q is a vowel.\n", s)
	} else if s == "w" || s == "y" {
		fmt.Printf("%q is sometimes a vowel, sometimes not.\n", s)
	} else {
		fmt.Printf("%q is a consonant.\n", s)
	}

	// Notice that:
	//
	// I didn't use IndexAny for the else if above.
	//
	// It's because, calling a function is a costly operation.
	// And, this way, the code is simpler.
}

func main5() {
	if len(os.Args) != 2 {
		fmt.Println("Requires age")
		return
	}

	age, err := strconv.Atoi(os.Args[1])

	if err != nil || age < 0 {
		fmt.Printf(`Wrong age: %q`+"\n", os.Args[1])
		return
	} else if age > 17 {
		fmt.Println("R-Rated")
	} else if age >= 13 && age <= 17 {
		fmt.Println("PG-13")
	} else if age < 13 {
		fmt.Println("PG-Rated")
	}
}

// 🛑 DON'T DO THIS:
//     It's hard to read.
//     It's just an exercise.

func main6() {
	if len(os.Args) != 2 {
		fmt.Println("Requires age")
		return
	} else if age, err := strconv.Atoi(os.Args[1]); err != nil || age < 0 {
		fmt.Printf(`Wrong age: %q`+"\n", os.Args[1])
		return
	} else if age > 17 {
		fmt.Println("R-Rated")
	} else if age >= 13 && age <= 17 {
		fmt.Println("PG-13")
	} else if age < 13 {
		fmt.Println("PG-Rated")
	}
}
