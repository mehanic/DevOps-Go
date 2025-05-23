package main

import(
	"fmt"
    "time"
)

func main(){
    //fact:=make(chan int)
    go func(){
        for{
            for _,r:=range `-\|/`{
                fmt.Printf("\r%c",r)
                time.Sleep(time.Millisecond*100)
            }
        }
    }()
    go fmt.Println(Fact(45))
    time.Sleep(time.Millisecond*1000)
}

func Fact(n uint)uint{
    if n==0{
        return 1
    }
    return n*Fact(n-1)
}