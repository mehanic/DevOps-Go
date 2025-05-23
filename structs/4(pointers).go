package main

import "fmt"

func main(){
    x:=10

    fmt.Println("X before increment: ",x)
    increment(&x)

    fmt.Println("X after increment: ",x)
}

func increment(p *int){
    *p+=1
}