package main
import "fmt"


func main(){
	var (
		a int = 3
		b int = 2
		S int
		P int
	)
	S=a*b
	P=2*(a+b)
	fmt.Printf("Triangle area S=:%d\nTriangle perimeter P=:%d\n", S, P)

}