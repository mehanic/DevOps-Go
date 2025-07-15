package main

import "fmt"

func main() {
	var a interface{} // в го для слаботипизрованных данных выделен интерфейс
	a = "jelly"

	fmt.Printf("Name is %s", a)
	a = 42

	fmt.Printf("Name is %d", a)
}

//Name is jellyName is 42