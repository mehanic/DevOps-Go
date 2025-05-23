package persons

import "fmt"

func GetSum(a, b int) int {
	return a + b
}

func getMultiple(a, b int) int {
	return a * b
}

func GetMultiplication(a, b int) int {
	return getMultiple(a, b)
}

func main1() {
	c := GetSum(3, 4)
	fmt.Println(c)
	d := GetMultiplication(3, 4)
	fmt.Println(d)
}
