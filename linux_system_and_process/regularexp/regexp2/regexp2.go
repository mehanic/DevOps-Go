package main

import (
	"fmt"
	"regexp"
)

func main() {
	// First text and pattern
	text1 := "This is a pythonnnn and python aaa haaaafd xyzaaaaaaaa"
	myPat1 := `\bxyza{8}\b`

	// Compile the regular expression
	re1 := regexp.MustCompile(myPat1)

	// Find all matches in the text
	matches1 := re1.FindAllString(text1, -1)
	fmt.Println(matches1)

	// Second text and pattern
	text2 := "xaz asdfa sdf xaaz xaaaaaaaz xaaaaz  xz"
	myPat2 := `xa?z`

	// Compile the regular expression
	re2 := regexp.MustCompile(myPat2)

	// Find all matches in the text
	matches2 := re2.FindAllString(text2, -1)
	fmt.Println(matches2)
}

