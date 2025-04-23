package main
import (
	"fmt"
	"math"
)

func main(){
	var (
		a =3
		b = 4
	)

	if a!=0 && b!=0{
		fmt.Printf("Sum a+b=: %d\nProduct a * b =: %d\nAbsolute values a =:%v b =: %v\n", a+b, a*b, math.Abs(float64(a)), math.Abs(float64(b)))


		}	
}