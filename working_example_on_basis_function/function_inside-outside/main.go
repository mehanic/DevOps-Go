package __3

import "fmt"

func doAppend(sl []int) {
	sl = append(sl, 100)
	fmt.Println("inside: ", sl) // inside:
}
func first() {
	x := []int{1, 2, 3}
	doAppend(x)
	fmt.Println("outside: ", x) // outside:
}

func main() {
	first()
}
