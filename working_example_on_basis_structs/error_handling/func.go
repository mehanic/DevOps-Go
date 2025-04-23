package error_handling

import (
	"fmt"
	"os"
)

/*

If you have written any Go code you have probably encountered the built-in error type.
Go code uses error values to indicate an abnormal state. For example, the os.Open function
returns a non-nil error value when it fails to open a file.

*/

func OpenFile() {
	file, err := os.Open("filename1.txt") // This file is not exist, so it will give error.

	if err != nil { // if error exist
		if path_error, ok := err.(*os.PathError); ok { // type assertion
			fmt.Println("File could not be found! : ", path_error.Path)
			return
		} else {
			fmt.Println("File could not be found!")
		}

	} else {
		fmt.Println(file.Name())
	}
}

func NumberCheck(num interface{}) {
	number, ok := num.(int)
	fmt.Println(number, ok)
}

func GetNumber() {

	var number1 interface{} = 7

	NumberCheck(number1)

	var number2 interface{} = "Berkay"

	NumberCheck(number2)
}
