package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Enter your city: ")
	city, _ := reader.ReadString('\n')
	if city != "\n" {
		fmt.Print("You live in " + city) 
	} else {
		fmt.Println("Bye bye!!! :D") 
	}
}
