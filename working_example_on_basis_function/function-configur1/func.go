package main

import "fmt"

func test(someFunction func(int) int) {
	println(someFunction(25))
}

func test2(x string) func() {
	return func() {
		println(x)
	}
}

func main() {
	first()
	first()
	sum(5, 6)
	value := sumWithResult(3, 4)
	println(value)
	sum, sub, multip, div := mathFunc(8, 2)
	println(sum, sub, multip, div)

	a := func() {
		println("a")
	}
	a()
	b := func(a int, b int) {
		println(a + b*a)
	}
	b(2, 3)

	c := func(a int, b int) int {
		return a + b*a
	}
	println(c(5, 10))

	sqrt := func(x int) int {
		return x * x
	}

	test(sqrt)
	test2("test2")()
}

func first() {
	println("Hello from first")
}

func sum(a int, b int) {
	println(a + b)
}

func sumWithResult(a int, b int) int {
	return a + b
}

func mathFunc(a int, b int) (int, int, int, int) {
	sum := a + b
	sub := a - b
	multip := a * b
	div := a / b
	return sum, sub, multip, div
}

func mathFunc2(a int, b int) (sum int, sub int, multip int, div int) {
	sum = a + b
	sub = a - b
	multip = a * b
	div = a / b
	return
}



// Function declarations
func saySomething(s string) (string, string) {
	return s, "World"
}

func itIsNotEnd(p string) (string, string) {
	return p, "wait me"
}

func main2() {
	// Using saySomething
	greeting, context := saySomething("Hello")
	fmt.Println("saySomething output:", greeting, context)

	// Using itIsNotEnd
	part1, part2 := itIsNotEnd("This is not the end")
	fmt.Println("itIsNotEnd output:", part1, part2)

	// Combining both functions
	firstMessage, secondMessage := saySomething("Start")
	finalMessage, additionalMessage := itIsNotEnd(secondMessage)

	fmt.Println("Combined function output:")
	fmt.Println("First message:", firstMessage)
	fmt.Println("Second message:", secondMessage)
	fmt.Println("Final message:", finalMessage)
	fmt.Println("Additional message:", additionalMessage)
}

