package main

import "fmt"


func main() {
	greetingWithFilter("Prasetiyo", badWordFilter) // Hello Prasetiyo
	greetingWithFilter1("Anjing", badWordFilter)    // Hello ***
	greetingWithFilter("Babi", badWordFilter)      // Hello ***
}

func greetingWithFilter(name string, filter func(string) string) {
	nameFiltered := filter(name)
	fmt.Println("Hello", nameFiltered)
}


func badWordFilter(name string) string {
	badWords := []string{"Babi", "Anjing", "Monyet"}
	if name == badWords[0] || name == badWords[1] || name == badWords[2] {
		return "***"
	} else {
		return name
	}
}

type Filter func(string) string

func greetingWithFilter1(name string, filter Filter) {
	nameFiltered := filter(name)
	fmt.Println("Hello", nameFiltered)
}
