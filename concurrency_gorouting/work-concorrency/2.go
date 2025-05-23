package chapter9

import (
	"fmt"
	"runtime"
	"sync"
)

var wg sync.WaitGroup
var mu sync.Mutex

// Exercise 1: Goroutines and WaitGroup
func Exercise1() {

	wg.Add(2)

	go message1()
	go message2()

	wg.Wait()
}

func message1() {
	fmt.Println("Hello")
	wg.Done()
}

func message2() {
	fmt.Println("World")
	wg.Done()
}

// -----------------------------------------------
// Exercise 2: Structs, Methods, and Interfaces
type person struct {
	firstName string
	lastName  string
	age       int
}

// Method for the person struct
func (p *person) speak() {
	fmt.Println("Nice to meet you, my name is", p.firstName, p.lastName, "and I am", p.age, "years old")
}

// Defining an interface
type human interface {
	speak()
}

// Function that accepts any type implementing the human interface
func saySomething(h human) {
	h.speak()
}

func Exercise2() {

	p := person{"Micael", "Ramos", 18}

	saySomething(&p) // Passing the pointer to the struct
}

// -----------------------------------------------
// Exercise 3: Using Mutex to Prevent Race Conditions
func Exercise3() {

	goroutines := 500
	counter := 0

	wg.Add(goroutines)

	for i := 0; i < goroutines; i++ {
		go func() {
			mu.Lock()
			runtime.Gosched()
			counter++
			mu.Unlock()
			wg.Done()
		}()
	}

	wg.Wait()
	fmt.Println(counter)
}
