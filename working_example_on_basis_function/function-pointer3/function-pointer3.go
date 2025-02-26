package main

import "fmt"

func newProfile(name string, age *int) {
	var defaultAge int = 18
	if age == nil {
		age = &defaultAge
	}
	fmt.Println(name, *age)
}

func main() {

	newProfile("John", nil) // will use default age 18
}

