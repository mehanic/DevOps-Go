package pointers

import "fmt"

func PointerExtractor(num int) {

	/*

		Pointer is a variable that is used to store the memory address of another variable.
		Pointers in Golang is also termed as the special variables.
		The variables are used to store some data at a particular memory address in the system.
		The memory address is always found in hexadecimal format(starting with 0x like 0xFFAAF etc.).

	*/

	fmt.Printf("The Number is %v \n", num)

	// declaration of pointer with *
	var p *int // this is passing by referance.

	// initialization of pointer with &
	p = &num

	fmt.Println("Address of number is ", p)

	newnumber := num + 10

	fmt.Printf("\nThe New number is %v ", newnumber)

	var newp *int
	newp = &newnumber

	fmt.Println("\nThe Address of new number is ", newp)
	fmt.Println("\nThe Address of main number is ", p)

}
