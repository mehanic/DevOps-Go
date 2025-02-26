package __2

import "fmt"

func changeValueAtZeroIndex(array [2]int) {
	array[0] = 3
	fmt.Println("inside: ", array[0]) // Will print 3
}
func first() {
	x := [2]int{}
	changeValueAtZeroIndex(x)
	fmt.Println(x) // Will print 0
}

func main() {
	first()
}
