package main

import (
	"fmt"
	"strconv"
	"time"
)

var pizzaNum = 0
var pizzaName = ""

func makeDough(stringChan chan string) {
	pizzaNum++
	pizzaName = "Pizza #" + strconv.Itoa(pizzaNum)

	fmt.Println("Make Dough and send", pizzaName, "for sauce")

	stringChan <- pizzaName

	time.Sleep(time.Millisecond * 10)
}

func addSauce(stringChan chan string) {
	pizza := <-stringChan

	fmt.Println("Add sauce and send", pizza, "for toppings")

	stringChan <- pizzaName

	time.Sleep(time.Millisecond * 10)
}

func addToppings(stringChan chan string) {

	pizza := <-stringChan

	fmt.Println("Add toppings to ", pizza, "and ship")

	time.Sleep(time.Millisecond * 10)
}

func main() {

	stringChan := make(chan string)

	for i := 0; i < 3; i++ {
		go makeDough(stringChan)
		go addSauce(stringChan)
		go addToppings(stringChan)
		time.Sleep(time.Millisecond * 5000)
	}
}

// Make Dough and send Pizza #1 for sauce
// Add toppings to  Pizza #1 and ship
// Make Dough and send Pizza #2 for sauce
// Add sauce and send Pizza #2 for toppings
// Add toppings to  Pizza #2 and ship
// Make Dough and send Pizza #3 for sauce
// Add sauce and send Pizza #3 for toppings
// Add toppings to  Pizza #3 and ship
