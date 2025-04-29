package main

import (
	"fmt"
	"io"
	"os"
)

func main() {

	fmt.Println("Go tutorial fo file handling")
	content := "This need to be in a file"
	file, err := os.Create("./rohan.txt")
	defer file.Close()
	checkNilErr(err)

	len, err := io.WriteString(file, content)
	checkNilErr(err)

	fmt.Println("Length is :", len)
	readFile("rohan.txt")
}

func readFile(filename string) {
	op, err := os.ReadFile(filename)
	checkNilErr(err)
	fmt.Printf("Type of output is %T", op)
	fmt.Println("Output is :", string(op))
}

func checkNilErr(err error) {
	if err != nil {
		panic(err)
	}
}


// Go tutorial fo file handling
// Length is : 25
// Type of output is []uint8Output is : This need to be in a file

