package main

import "fmt"

func main() {
	age := 30
	println("My age is " + fmt.Sprint(age))
	fmt.Printf("My age is %d", age)

	name := "John"
	fmt.Printf("My age is %d My name is %s", age, name)

	value := 100.679
	fmt.Printf("My age is %d My name is %s My value is %f", age, name, value)

	bool := true
	fmt.Printf("My age is %d My name is %s My value is %f This is %t", age, name, value, bool)
}

//My age is 30
//My age is 30My age is 30 My name is JohnMy age is 30
// My name is John My value is 100.679000My age is 30 My name is John My value is 100.679000 This is true%
