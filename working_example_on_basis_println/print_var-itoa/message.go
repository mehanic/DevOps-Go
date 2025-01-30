package main

import (
	"fmt"
	"strconv"
)

func main() {
	age := 23
	message := "Happy " + strconv.Itoa(age) + "rd Birthday!"
	fmt.Println(message)
}

//Happy 23rd Birthday!