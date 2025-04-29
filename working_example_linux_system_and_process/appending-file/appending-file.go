package main

import (
	"fmt"
	"os"
)

func main() {
	file, err := os.OpenFile("example.txt",
		os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer file.Close()
	data := "\nThis is some additional text"
	_, err = fmt.Fprintln(file, data)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("Data appended successfully")
}


//└> $ go run appending_file.go
//Data appended successfully

