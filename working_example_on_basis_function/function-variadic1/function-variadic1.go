package main

import "fmt"

func main() {
	sayHello1()
	sayHello1("Sammy")
	sayHello1("Sammy", "Jessica", "Drew", "Jamie")
}

func sayHello1(names ...string) {
	if len(names) == 0 {
		fmt.Println("nobody to greet")
		return
	}
	for _, n := range names {
		fmt.Printf("Hello %s\n", n)
	}
}

