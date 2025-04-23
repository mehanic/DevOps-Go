package main
import "fmt"

func main(){
	var (
		a =4
		b=2
		c=3
		V int
		S int
	)

	V=a*b*c
	S = 2*(a*b+b*c+c*a)

	fmt.Printf("Parallelepiped volume V=:%d\nParallelepiped total surface area S=:%d\n", V, S)


}