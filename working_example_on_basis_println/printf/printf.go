package main

import "fmt"

func main() {
	ops, ok, fail := 2350, 543, 433

	fmt.Println(
		"total:", ops, "- success:", ok, "/", fail,
	)

	fmt.Printf(
		"total: %d - success: %d / %d\n",
		ops, ok, fail,
	)

	main1()
	main2()
	main4()
	main5()
	main6()
	main7()
	main8()
	main9()
	main10()
	main11()
	main12()
	main13()
}

func main1() {
	var brand string

	// prints the string in quoted-form like this ""
	fmt.Printf("%q\n", brand)

	brand = "Google"
	fmt.Printf("%q\n", brand)
}

func main2() {
	// without newline
	fmt.Println("hihi")

	// with newline:
	//   \n = escape sequence
	//   \  = escape character
	fmt.Println("hi\nhi")

	// escape characters:
	//   \\ = \
	//   \" = "
	fmt.Println("hi\\n\"hi\"")

	// I'm using multiple declarations instead of singular
	var (
		speed int
		heat  float64
		off   bool
		brand string
	)

	fmt.Printf("%T\n", speed)
	fmt.Printf("%T\n", heat)
	fmt.Printf("%T\n", off)
	fmt.Printf("%T\n", brand)
}

func main4() {
	var (
		planet   = "venus"
		distance = 261
		orbital  = 224.701
		hasLife  = false
	)

	// swiss army knife %v verb
	fmt.Printf("Planet: %v\n", planet)
	fmt.Printf("Distance: %v millions kms\n", distance)
	fmt.Printf("Orbital Period: %v days\n", orbital)
	fmt.Printf("Does %v have life? %v\n", planet, hasLife)

	// argument indexing - indexes start from 1
	fmt.Printf(
		"%v is %v away. Think! %[2]v kms! %[1]v OMG.\n",
		planet, distance,
	)

}

func main5() {
	fmt.Printf("I'm %d years old.\n", 30)
}

func main6() {
	fmt.Printf("My name is %s and my lastname is %s.\n", "Inanc", "Gumus")

	// BONUS
	msg := "My name is %s and my lastname is %s.\n"
	fmt.Printf(msg, "Inanc", "Gumus")
}

func main7() {
	tf := false
	fmt.Printf("These are %t claims.\n", tf)
}

func main8() {
	fmt.Printf("Temperature is %.1f degrees.\n", 29.5)
}

func main9() {
	fmt.Printf("%q\n", "hello world")
}

func main10() {
	fmt.Printf("Type of %d is %[1]T\n", 3)
}

func main11() {
	fmt.Printf("Type of %.2f is %[1]T\n", 3.14)
}

func main12() {
	fmt.Printf("Type of %s is %[1]T\n", "hello")
}

func main13() {
	fmt.Printf("Type of %t is %[1]T\n", true)
}


// total: 2350 - success: 543 / 433
// total: 2350 - success: 543 / 433
// ""
// "Google"
// hihi
// hi
// hi
// hi\n"hi"
// int
// float64
// bool
// string
// Planet: venus
// Distance: 261 millions kms
// Orbital Period: 224.701 days
// Does venus have life? false
// venus is 261 away. Think! 261 kms! venus OMG.
// I'm 30 years old.
// My name is Inanc and my lastname is Gumus.
// My name is Inanc and my lastname is Gumus.
// These are false claims.
// Temperature is 29.5 degrees.
// "hello world"
// Type of 3 is int
// Type of 3.14 is float64
// Type of hello is string
// Type of true is bool
