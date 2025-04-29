package pointers

import "fmt"

/*
All variables in Go have a zero value.
This is true even for a pointer. If you declare a pointer to a type,
but assign no value, the zero value will be nil. nil is a way to say that
“nothing has been initialized” for the variable.
*/

type Auto struct {
	brand string
}

func CheckPointer() {
	var auto *Auto

	fmt.Println(auto)
	ChangeBrand(auto)
	fmt.Println(auto)

}

func ChangeBrand(auto *Auto) {
	if auto == nil {
		fmt.Println("auto is nil")
		return
	}

	auto.brand = "Mercedes Benz"
	fmt.Println(auto)

}
