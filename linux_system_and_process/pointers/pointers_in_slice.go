package pointers

import "fmt"

func PointersInSlice() {
	a := [5]int{1, 2, 3, 4, 5}
	b := a

	fmt.Println(a, b)

	a[1] = 42

	fmt.Println(a, b)

	fmt.Println("As you can see, array b didn't change because it directly copied file from a.\nBut if we don't give length of the array, it will change because it follows the a dynamically.")

	c := []int{1, 2, 3, 4, 5}
	d := c

	fmt.Println(c, d)

	c[1] = 42

	fmt.Println(c, d)

}
