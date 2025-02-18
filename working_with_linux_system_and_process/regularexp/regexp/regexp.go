package main

import (
	"fmt"
	"regexp"
)

func main() {
	text := "This . is a python and it is easy to learn and it is popular one for dev and automation"

	// Pattern 1: i[ston] -> Matches 'is', 'it', 'io', 'in'
	myPat1 := `i[ston]`
	re1 := regexp.MustCompile(myPat1)
	fmt.Println("Matches for pattern 1:", re1.FindAllString(text, -1))

	// Pattern 2: '\.' -> Matches literal dot
	myPat2 := `\.`
	re2 := regexp.MustCompile(myPat2)
	fmt.Println("Matches for pattern 2:", re2.FindAllString(text, -1))

	// Pattern 3: Matching an IP address pattern
	text = "This is an ip address of my db1 server: 255.255.255.255 2456234512341234"
	myPat3 := `\d\d\d\.\d\d\d\.\d\d\d\.\d\d\d`
	re3 := regexp.MustCompile(myPat3)
	fmt.Println("Matches for pattern 3:", re3.FindAllString(text, -1))

	// Pattern 4: '\w' -> Word characters and '\W' -> Non-word characters
	text = "This is python @ 345 _ - ("
	fmt.Println("Matches for word characters:", regexp.MustCompile(`\w`).FindAllString(text, -1))
	fmt.Println("Matches for non-word characters:", regexp.MustCompile(`\W`).FindAllString(text, -1))

	// Pattern 5: '\d\d' -> Matches two consecutive digits
	text = "456 90 this is about decimal re98gex"
	myPat5 := `\d\d`
	re5 := regexp.MustCompile(myPat5)
	fmt.Println("Matches for pattern 5 (two consecutive digits):", re5.FindAllString(text, -1))
}

