package main

import (
	"fmt"
)

func main() {
	greeting()
}

func greeting() {
	fmt.Print("yes")
}

func main1() {
	for i := 1; i <= 3; i++ {
		greeting1()
	}
}

func greeting1() {
	fmt.Println("yes")
}
