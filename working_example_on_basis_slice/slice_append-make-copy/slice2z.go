// make(),append(),copy()
package main

import "fmt"

func main() {

	// s := make([]int, 3) // [0,0,0]
	s := []int{1, 3, 5}
	// append 
	s = append(s, 8, 2, 10)
	// copy
	t := make([]int, len(s))
	n := copy(t, s)
	fmt.Println(s) //[1,3,5,8,2,10]
	fmt.Println(t) //[1,3,5,8,2,10]
	fmt.Println(n) //6

}
