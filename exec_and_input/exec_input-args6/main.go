package main


import (
	"fmt"
	"os"
)


func main() {
	fmt.Println("There are", len(os.Args)-1, "people !")
	fmt.Println("Hello great", os.Args[1], "!")
	fmt.Println("Hello great", os.Args[2], "!")
	fmt.Println("Hello great", os.Args[3], "!")
	fmt.Println("Hello great", os.Args[4], "!")
	fmt.Println("Hello great", os.Args[5], "!")
	fmt.Println("Nice to meet you all.")
}


// main.go Alice Bob Charlie Dave Eve
// There are 5 people !
// Hello great Alice !
// Hello great Bob !
// Hello great Charlie !
// Hello great Dave !
// Hello great Eve !
// Nice to meet you all.
