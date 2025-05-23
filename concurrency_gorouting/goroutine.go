package main
import "fmt"

func main(){
   go func(){
     for i:=0;i<10;i++{
        fmt.Println(i)
     }
   }()

   go func(){
     for i:=0;i<10;i++{
       fmt.Println("hello")
     }
   }()

   var input string
   fmt.Scanln(&input)
}