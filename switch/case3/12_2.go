package main

import (
	"fmt"
    "time"
)

func main(){
    number:=make(chan int)

    go func(){
        number <-42
    }()

    time.Sleep(time.Millisecond*100)
    //fmt.Println(<-number)
    select{
        case n:=<-number:
            fmt.Println(n)
        default:
            fmt.Println("nothing")
    }
}