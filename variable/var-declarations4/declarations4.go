package main

import "fmt"

func main() {
	var speed int
	// speed = "100"

	var running bool
	// running = 1

	var force float64
	// speed = force

	// Go automatically converts the typeless
	//   integer literal to float64 automatically
	force = 1

	// keep the compiler happy
	_, _, _ = speed, running, force

	println("--------------")
	var (
		name   string
		age    int
		famous bool
	)

	// Example #1
	name = "Newton"
	age = 84
	famous = true

	fmt.Println(name, age, famous)

	// Example #2
	name = "Somebody"
	age = 20
	famous = false

	fmt.Println(name, age, famous)

	// save the previous value of the variable
	// to a new variable
	var prevName string
	prevName = name

	// overwrite the value of the original variable
	// by assigning to it
	name = "Einstein"

	fmt.Println("previous name:", prevName)
	fmt.Println("current name :", name)

	println("--------------")

	var counter int

	fmt.Println("counter's name : counter")
	fmt.Println("counter's value:", counter)
	fmt.Printf("counter's type : %T\n", counter)

	counter = 10 // OK
	// counter = "ten" // NOT OK

	fmt.Println("counter's value:", counter)

	counter++
	fmt.Println("counter's value:", counter)
	println("--------------")

}

// --------------
// Newton 84 true
// Somebody 20 false
// previous name: Somebody
// current name : Einstein
// --------------
// counter's name : counter
// counter's value: 0
// counter's type : int
// counter's value: 10
// counter's value: 11
// --------------
