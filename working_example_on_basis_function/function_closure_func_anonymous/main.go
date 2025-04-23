package main

import (
	"fmt"
)

func main()  {
    x :=40

    increment := func() int{
        x++
        return x
    }
    fmt.Println(increment())
    fmt.Println(increment())
    fmt.Println(increment)// this will return the address of func expression

    //func () int is called anonymous function and
    // the variable increment is called function expression ie 
    // assigning  function to the varible
}