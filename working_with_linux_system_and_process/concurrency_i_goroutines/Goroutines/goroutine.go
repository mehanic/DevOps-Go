package main

import "fmt"
import "runtime"

func print (till int, message string) {
	for i :=0; i < till; i++{
		fmt.Println((i+1), message)
	}
}

func main() {
	runtime.GOMAXPROCS(2)
	go print(5, "hai")
	print (5, "hallo")

	var input string
	fmt.Scanln(&input)
}

// ╰─λ go run goroutine.go                                                                   0 (0.001s) < 15:41:11
// 1 hallo
// 2 hallo
// 3 hallo
// 4 hallo
// 5 hallo
// 1 hai
// 2 hai
// 3 hai
// 4 hai
// 5 hai
