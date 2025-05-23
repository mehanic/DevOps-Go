package main

import "fmt"

func main(){
	var (
		a =3
		b = 4
	)

	if a!=0 && b!=0{
		fmt.Printf("Sum a + b =: %d\nProduct a * b =: %d\nSquares a =:%d b =: %d\n",a+b,a*b,a*a,b*b)
	}
}