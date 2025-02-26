package main

import "fmt"

func Set(x ...int)[]int{
	var arr []int=[]int{}
	fmt.Printf("%T %#v\n",x,arr)
	for i:=0 ; i<len(x);i++{
		for j:=i ; j<len(x);j++{
			if x[i]!=x[j]{
				arr=append(arr,x[i])
			
				}else{
							
			}
		}
	}
	return arr
}
func main(){
	fmt.Println(Set([]int{1,2,2,4,5,4,3,6,7,8,7,6,9,6,7,8,9}...))
	// fmt.Println(Set([]int{1,2,2}...))
}

// []int []int{}
// [1 1 1 1 1 1 1 1 1 1 1 1 1 1 1 1 2 2 2 2 2 2 2 2 2 2 2 2 2 2 2 2 2 2 2 2 2 2 2 2 2 2 2 2 4 4 4 4 4 4 4 4 4 4 4 4 5 5 5 5 5 5 5 5 5 5 5 5 4 4 4 4 4 4 4 4 4 4 4 3 3 3 3 3 3 3 3 3 3 6 6 6 6 6 6 6 7 7 7 7 7 7 8 8 8 8 8 8 7 7 7 7 7 6 6 6 6 9 9 9 6 6 6 7 7 8]