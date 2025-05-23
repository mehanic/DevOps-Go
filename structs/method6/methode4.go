package main

import "fmt"

type person struct {
	name string
	age  int
}

func (p person) print() {
	fmt.Println("Назва:", p.name)
	fmt.Println("Вік:", p.age)
}

func (p person) eat(meal string) {
	fmt.Println(p.name, "ест", meal)
}

func main() {

	var tom = person{name: "Tom", age: 24}
	tom.print()
	tom.eat("борщ с м'ясом, і обов'язково дві порції")
}

// Назва: Tom
// Вік: 24
// Tom ест борщ с м'ясом, і обов'язково дві порції
