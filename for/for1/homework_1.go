package main

import "fmt"

func main(){
	var number = 782838
	var max = number%10
	for number>0{
		if max<number%10{
			max=number
		}
		number/=10
	}
	fmt.Println(max)
}
