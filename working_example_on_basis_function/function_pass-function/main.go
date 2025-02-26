package __1

import "fmt"

// public and private
// import "github.com/repo/examples/say"
// package say
func PrintHello() {
	fmt.Println("Hello")
}
func printWorld() {
	fmt.Println("World")
}
func PrintHelloWorld() {
	PrintHello()
	printWorld()
}

func main() {
	//say.PrintHello()
	PrintHello()
	//say.PrintHelloWorld()
	PrintHelloWorld()
}
