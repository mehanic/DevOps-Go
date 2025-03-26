
package main

import "fmt"

func main(){
	var arr = []int{2, 3, 5, 4, 7}
	var slice = []map[int]int{}
	for _,el_1:=range arr{
		for _,el_2:=range arr{
			if el_1%el_2==3{
				// fmt.Println(el_1,el_2)
				 x_map:=make(map[int]int)
				x_map[el_1]=el_2
				slice=append(slice,x_map)
			}
		}
	}
	fmt.Println(slice)
}