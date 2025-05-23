package main

import "fmt"

func main(){
	var arr = []int{1,2,4,4,5,6,7,12,20}
	var n = 4
	var is int
	var x_map=make(map[int]int)
	// fmt.Println(x_map[n])
	for _,el:=range arr{
		x_map[el]++
	}
	for key,val:=range x_map{
		if val==1 && n==key{
			is=-1
			// fmt.Println("Vaalaykum assalom")
		}else if val>1 && key==n{
			// fmt.Println("salom")
			is=x_map[n]
		}
	} 
	fmt.Println(is)
}