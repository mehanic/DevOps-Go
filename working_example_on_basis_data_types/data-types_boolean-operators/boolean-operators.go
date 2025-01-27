package main

import (
	"fmt"
	"strings"
)

func main() {
	speed := 100 // #1
	// speed := 10 // #2

	fast := speed >= 80
	slow := speed < 20

	fmt.Printf("fast's type is %T\n", fast)

	fmt.Printf("going fast? %t\n", fast)
	fmt.Printf("going slow? %t\n", slow)

	fmt.Printf("is it 100 mph? %t\n", speed == 100)
	fmt.Printf("is it not 100 mph? %t\n", speed != 100)

	main1()
	fmt.Println("--------------------")
	main2()
	fmt.Println("--------------------")
	main3()
	fmt.Println("--------------------")
	main4()
	fmt.Println("--------------------")
	main5()
	fmt.Println("--------------------")
	main6()
	fmt.Println("--------------------")
	main7()
	fmt.Println("--------------------")
	main8()
	fmt.Println("--------------------")
	main9()

}

func main1() {
	var speed int
	// speed = "100" // NOT OK

	var running bool
	// running = 50 // NOT OK

	var force float64
	// speed = force // NOT OK

	_, _, _ = speed, running, force
}

func main2() {
	var speed int
	speed = 50 // OK

	var running bool
	running = true // OK

	var force float64
	speed = int(force)

	_, _, _ = speed, running, force
}

func main3() {
	speed := 100 // #1
	// speed := 10 // #2

	fast := speed >= 80
	slow := speed < 20

	fmt.Printf("going fast? %t\n", fast)
	fmt.Printf("going slow? %t\n", slow)

	fmt.Printf("is it 100 mph? %t\n", speed == 100)
	fmt.Printf("is it not 100 mph? %t\n", speed != 100)

	fmt.Println(strings.Repeat("-", 25))

	// -------------------------
	// #3
	speedB := 150.5
	faster := speedB > 100 // ok: untyped

	fmt.Println("is the other car going faster?", faster)

	// ERROR: Type Mismatch
	// faster = speedB > speed
	faster = speedB > float64(speed)
	fmt.Println("is the other car going faster?", faster)

	// #4
	// ERROR:
	// only numeric values are comparable with
	// ordering operators: <, <=, >, >=

	// fmt.Println(true > faster)
}

func main4() {
	// remove the comments and run
	// I've commented the lines it's because of the warnings

	// fmt.Println("true  && true  =", true && true)
	fmt.Println("true  && false =", true && false)
	fmt.Println("false && true  =", false && true)
	// fmt.Println("false && false =", false && false)
}

func main5() {
	speed := 100
	fmt.Println("within limits?",
		speed >= 40 && speed <= 55,
	)

	speed = 20
	fmt.Println("within limits?",
		speed >= 40 && speed <= 55,
		// ^- short-circuits in the first expression here
		//    because, it becomes false
	)

	speed = 50
	fmt.Println("within limits?",
		speed >= 40 && speed <= 55,
	)

	// ERROR: invalid
	// both operands should be booleans
	// 1 && 2
	fmt.Println(1 == 1 && 2 == 2)
}

func main6() {
	// remove the comments and run
	// I've commented the lines it's because of the warnings

	// fmt.Println("true  || true  =", true || true)
	fmt.Println("true  || false =", true || false)
	fmt.Println("false || true  =", false || true)
	// fmt.Println("false || false =", false || false)
}

func main7() {
	color := "red"

	fmt.Println("reddish colors?",
		// true || false => true (short-circuits)
		color == "red" || color == "dark red",
	)

	color = "dark red"

	fmt.Println("reddish colors?",
		// false || true => true
		color == "red" || color == "dark red",
	)

	fmt.Println("greenish colors?",
		// false || false => false
		color == "green" || color == "light green",
	)
}

func main8() {
	fmt.Println(
		"hi" == "hi" && 3 > 2,    //   true  && true  => true
		"hi" != "hi" || 3 > 2,    //   false || true  => true
		!("hi" != "hi" || 3 > 2), // !(false || true) => false
	)
}

func main9() {
	var on bool

	on = !on
	fmt.Println(on)

	on = !!on
	fmt.Println(on)
}

// fast's type is bool
// going fast? true
// going slow? false
// is it 100 mph? true
// is it not 100 mph? false
// --------------------
// --------------------
// going fast? true
// going slow? false
// is it 100 mph? true
// is it not 100 mph? false
// -------------------------
// is the other car going faster? true
// is the other car going faster? true
// --------------------
// true  && false = false
// false && true  = false
// --------------------
// within limits? false
// within limits? false
// within limits? true
// true
// --------------------
// true  || false = true
// false || true  = true
// --------------------
// reddish colors? true
// reddish colors? true
// greenish colors? false
// --------------------
// true true false
// --------------------
// true
// true
