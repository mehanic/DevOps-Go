package main

import "fmt"

var safe = true

func main() {
	fmt.Println(safe)
	println("--------------")

	tosafe1, speed := true, 50
	fmt.Println(tosafe1, speed)

	println("--------------")

	name, lastname := "Nikola", "Tesla"
	fmt.Println(name, lastname)

	birth, death := 1856, 1943
	fmt.Println(birth, death)

	on, off := true, false
	fmt.Println(on, off)

	degree, ratio, heat := 10.55, 30.5, 20.
	fmt.Println(degree, ratio, heat)

	nFiles, valid, msg := 10, true, "hello"

	fmt.Println(nFiles, valid, msg)
	println("--------------")

	var safe1 bool
	safe1, speed1 := true, 50
	fmt.Println(safe1, speed1)

}

// true
// --------------
// true 50
// --------------
// Nikola Tesla
// 1856 1943
// true false
// 10.55 30.5 20
// 10 true hello
// --------------
// true 50
