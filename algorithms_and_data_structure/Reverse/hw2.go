package main

import (
	"fmt"
    "math"
)

func main(){
    n:=5
    degreeOfTwo(n)
    mas:=[]int{1,2,3,4,5}
    Reverse(mas)
    mas2:=make([]int,3)
    copy(mas2,mas)
    fmt.Println(mas2)
}

func degreeOfTwo(n int){
    for i:=1;i<=n;i++{
        fmt.Println(math.Pow(2,float64(i)))
    }
}

func Reverse(mas []int){
    for i,j:=0,len(mas)-1;i<j;i,j = i+1,j-1{
        mas[i],mas[j]=mas[j],mas[i]
    }
    fmt.Println(mas)
}