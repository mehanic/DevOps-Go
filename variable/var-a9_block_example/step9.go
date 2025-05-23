package main

var name string = "Alixon" // Global

var (
	box1 string  = "Hello"
	box2 int     = 11
	box3 float32 = 22.3
	box4 bool    = true
)

func main() {

	println(name)
	println(box1, box2, box3, box4)
}

// Alixon
// Hello 11 +2.230000e+001 true
