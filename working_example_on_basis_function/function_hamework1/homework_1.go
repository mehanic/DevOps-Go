package main
import "fmt"

func main(){
	var n int
	fmt.Print("Enter the length of the array n=")
	fmt.Scanln(&n)
	var arr = make([]int, n)
	fmt.Printf("%T\n", arr)

	// Asking for input elements
	for i := 0; i < n; i++ {
		fmt.Printf("Enter element %d arr[%d]=", i, i)
		fmt.Scanln(&arr[i])
	}
	fmt.Println(arr)
}
