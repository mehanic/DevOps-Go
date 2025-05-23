package main

import "fmt"

func main(){
	var (
		a int =3
		V int
		S int

	)

	V = a*a*a
	S= 6*a
	fmt.Printf("Cube volume V=:%d\nCube total surface area S=:%d\n", V, S)

}