package main

import "fmt"

func main() {
	var c interface{}
	d := 3
	c = d
	var k int
	k = c.(int)
	fmt.Println(k + 3)
}
