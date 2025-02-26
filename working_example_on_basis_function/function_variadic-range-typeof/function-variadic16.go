package main

import (
	"fmt"
	"reflect"
)

func shoppingList1(items ...string) {

	fmt.Println(reflect.TypeOf(items))
	fmt.Println(items)

	for _, item := range items {
		fmt.Println(item)
	}

}

func main() {

	shoppingList1("milk", "bananas", "bread")

}

// []string
// [milk bananas bread]
// milk
// bananas
// bread
