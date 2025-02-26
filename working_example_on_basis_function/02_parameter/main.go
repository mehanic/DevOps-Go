package main

import "fmt"

// func main is the entry point to your program
func main() {
	fmt.Println("vim-go")

	greet("Clement")
}

// greet is a function with a parameter
// pass in an argument when calling greet
func greet(s string) {
	fmt.Println("Congratulations", s, "!")
}

// no overload
//func greet(s1, s2 string) {
//	fmt.Println(s1, s2)
//}

// func can pass multiple argument
