package main

import (
	"fmt"
	"os"
	"time"
)

func main() {
	city := "Paris"

	switch city {
	case "Paris":
		fmt.Println("France")
	}

	main1()
	main2()
	main3()
	main4()
	main5()
	main6()
	main7()
	main8()
	main9()
}

func main1() {
	if len(os.Args) < 2 {
		fmt.Println("Please provide a city name.")
		return
	}

	city := os.Args[1]

	switch city {
	case "Paris":
		fmt.Println("France")
	case "Tokyo":
		fmt.Println("Japan")
	}
}

func main2() {
	if len(os.Args) < 2 {
		fmt.Println("Please provide a city name.")
		return
	}

	city := os.Args[1]

	switch city {
	case "Paris":
		fmt.Println("France")
	case "Tokyo":
		fmt.Println("Japan")
	default:
		fmt.Println("Where?")
	}
}

func main3() {
	if len(os.Args) < 2 {
		fmt.Println("Please provide a city name.")
		return
	}

	city := os.Args[1]

	switch city {
	case "Paris", "Lyon":
		fmt.Println("France")
	case "Tokyo":
		fmt.Println("Japan")
	}
}

func main4() {
	i := 10
	switch {
	case i > 0:
		fmt.Println("positive")
	case i < 0:
		fmt.Println("negative")
	default:
		fmt.Println("zero")
	}
}

func main5() {
	i := 142
	switch {
	case i > 100:
		fmt.Print("big positive number")
	case i > 0:
		fmt.Print("positive number")
	default:
		fmt.Print("number")
	}
	fmt.Println()
}

func main6() {
	i := 142
	switch {
	case i > 100:
		fmt.Print("big ")
		fallthrough
	case i > 0:
		fmt.Print("positive ")
		fallthrough
	default:
		fmt.Print("number")
	}
	fmt.Println()
}

func main7() {
	switch i := 10; {
	case i > 0:
		fmt.Println("positive")
	case i < 0:
		fmt.Println("negative")
	default:
		fmt.Println("zero")
	}
}

func main8() {
	switch h := time.Now().Hour(); {
	case h >= 18:
		fmt.Println("good evening")
	case h >= 12:
		fmt.Println("good afternoon")
	case h >= 6:
		fmt.Println("good morning")
	default:
		fmt.Println("good night")
	}
}

func main9() {
	if len(os.Args) < 2 {
		fmt.Println("Gimme a month name.")
		return
	}

	switch m := os.Args[1]; m {
	case "Dec", "Jan", "Feb":
		fmt.Println("Winter")
	case "Mar", "Apr", "May":
		fmt.Println("Spring")
	case "Jun", "Jul", "Aug":
		fmt.Println("Summer")
	case "Sep", "Oct", "Nov":
		fmt.Println("Fall")
	default:
		fmt.Printf("%q is not a month.\n", m)
	}
}


// France
// Please provide a city name.
// Please provide a city name.
// Please provide a city name.
// positive
// big positive number
// big positive number
// positive
// good afternoon
// Gimme a month name.
