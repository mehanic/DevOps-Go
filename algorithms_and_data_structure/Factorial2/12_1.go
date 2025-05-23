package main

import (
	"fmt"
)

// func main(){
// //     numbers:=make(chan int)
// //     go generateNumbers(1000,numbers)
// //
// //     for{
// //         number,ok:=<-numbers
// //
// //         fmt.Println(number,ok)
// //
// //         if !ok{
// //             break
// //         }
// //     }
// //     numbers<-12
// //     for n:=range numbers{
// //         fmt.Println(n)
// //     }
// intCh := make(chan int, 3)
//     intCh <- 10
//     intCh <- 3
//     intCh <- 24
//     intCh <- 15  // блокировка - функция main ждет, когда освободится место в канале
//
//     fmt.Println(<-intCh)
//     fmt.Println("The End")
// }

// func generateNumbers(n int, res chan int){
//     for i:=0;i<=n;i++{
//         res <- i*2
//     }
//
//     close(res)
// }

func main() {

    intCh := make(chan int, 2)
    go factorial(5, intCh)
    fmt.Println(<-intCh)
    fmt.Println("The End")
}

func factorial(n int, ch chan<- int){

    result := 1
    for i := 1; i <= n; i++{
        result *= i
    }
    ch <- result
}