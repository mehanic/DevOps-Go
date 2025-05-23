package main

import "fmt"

func main() {
	name := "Nikola"
	fmt.Println(name)
	name, age := "Marie", 66
	fmt.Println(name, age)
	name, birth := "Albert", 1879
	fmt.Println(name, birth)

	println("--------------")

	i := 314
	f := 3.14
	s := "Hello"
	b := true

	fmt.Println(
		"i:", i,
		"f:", f,
		"s:", s,
		"b:", b,
	)
	println("--------------")

	a, b := 14, true

	fmt.Println(a, b)

	println("--------------")

	a, c := 42, "good"

	fmt.Println(a, c)

	println("--------------")

	sum := 27 + 3.5

	fmt.Println(sum)
	println("--------------")

	on, _ := true, true

	fmt.Println(on)
	println("--------------")

	age, yourAge := 10, 20
	age, ratio := 42, 3.14
	fmt.Println(age, yourAge, ratio)

	var speed int
	fmt.Println(speed)

	speed = 100
	fmt.Println(speed)

	speed = speed - 25
	fmt.Println(speed)

}

// Nikola
// Marie 66
// Albert 1879
// --------------
// i: 314 f: 3.14 s: Hello b: true
// --------------
// 14 true
// --------------
// 42 good
// --------------
// 30.5
// --------------
// true
// --------------
// 42 20 3.14
// 0
// 100
// 75
