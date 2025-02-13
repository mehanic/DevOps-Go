package main

import "fmt"

type person struct {
	name string
	age  int
}

func main() {

	tom := person{name: "Tom", age: 22}
	var tomPointer *person = &tom
	tomPointer.age = 29
	fmt.Println(tom.age) // 29
	(*tomPointer).age = 32
	fmt.Println(tom.age) // 32
}

// 29
// 32
