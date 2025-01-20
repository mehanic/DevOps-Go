package main

import "fmt"

type Planets []string

// метод типа Planets
func (planets Planets) terraform() {
	for i := range planets {
		planets[i] = "Новый " + planets[i]
	}
}

func main() {
	planets := []string{
		"Меркурий", "Венера", "Земля", "Марс",
		"Юпитер", "Сатурн", "Уран", "Нептун",
	}
	Planets(planets[3:4]).terraform()
	Planets(planets[6:]).terraform()
	fmt.Println(planets)
}
