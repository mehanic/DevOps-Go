package main

import(
	"fmt"
    "time"
)

func main(){
    t:=time.Now()

    fmt.Printf("Start: %s\n",t.Format(time.RFC3339))

    result1:=make(chan int)
    result2:=make(chan int)

//     go func(){
//         for{
//             for _,r:=range `-\|/`{
//                 fmt.Printf("\r%c",r)
//                 time.Sleep(time.Millisecond*100)
//             }
//         }
//     }()

    go CalculateSomething(1000,result1)
    go CalculateSomething(2000,result2)

//     time.Sleep(8*time.Second)
    fmt.Println(<-result1)
    fmt.Println(<-result2)
    fmt.Printf("Total time: %s\n",time.Since(t))

}

func CalculateSomething(n int,res chan int){
    t:=time.Now()

    result:=0
    for i:=0;i<=n;i++{
        result+=i*2
        time.Sleep(time.Millisecond*3)
    }

    fmt.Printf("Result: %d; Time: %s\n",result,time.Since(t))
    res <-result
}