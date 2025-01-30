package main

import (
	"fmt"
	"strings"
)

func commaCode(spam []string) string {
	spam[len(spam)-1] = "and " + spam[len(spam)-1]
	commaList := strings.Join(spam, ", ")
	return commaList
}

func main() {
	spam := []string{"apples", "bananas", "tofu", "cats"}
	fmt.Println(commaCode(spam))
}

//apples, bananas, tofu, and cats