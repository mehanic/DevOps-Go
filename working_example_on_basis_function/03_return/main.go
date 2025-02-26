package main

import "fmt"

func main() {
	fmt.Println(greet("Clement", "Orange"))

	n, s := greet2("OH", "HI")
	fmt.Println(n, s)
}

// func greet returns a string
func greet(s1, s2 string) string {
	return fmt.Sprintf("Congratulations %s and %s !!!", s1, s2)
}

// func greet2 returns an int and a string
func greet2(s1, s2 string) (int, string) {
	return 2, fmt.Sprintf("%s %s", s1, s2)
}

// func can return multiple value
