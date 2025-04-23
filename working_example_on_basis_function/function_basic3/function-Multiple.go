package main



import "fmt"

func multiple(a, b string) (string, string) {
	return a, b
}

func main1() {
	a, b := "hello", "my name Naufal"

	fmt.Println(multiple(a, b))
}
