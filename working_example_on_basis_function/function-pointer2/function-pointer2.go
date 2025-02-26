package main
import "fmt"

func createPointer(x int) *int{
	p := new(int)
	*p = x
	return p
}

func main() {

	p1 := createPointer(7)
	fmt.Println("p1:", *p1)     // p1: 7
	p2 := createPointer(10)
	fmt.Println("p2:", *p2)     // p2: 10
	p3 := createPointer(28)
	fmt.Println("p3:", *p3)     // p3: 28
}
