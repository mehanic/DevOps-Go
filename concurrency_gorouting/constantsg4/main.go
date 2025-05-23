package main

const a string = "this is someting"

func main() {
	const b = "this also something"
	println(a)
	println(b)
}

// ─λ go run main.go                                                                        0 (0.001s) < 13:56:43
// this is someting
// this also something
