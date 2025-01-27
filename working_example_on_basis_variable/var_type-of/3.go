package main

import (
	"fmt"
	"reflect"
)

func main() {
	var c interface{}
	c = "1231231"
	fmt.Println(c, reflect.TypeOf(c))
	c = 12323
	fmt.Println(c, reflect.TypeOf(c))
	var d int64
	d = 3
	c = d
	fmt.Println(c, reflect.TypeOf(c))
	arr := []int{1, 23, 54}
	c = arr
	fmt.Println(c, reflect.TypeOf(c))
}
