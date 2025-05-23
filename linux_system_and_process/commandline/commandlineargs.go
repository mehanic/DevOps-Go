package main
import "fmt"
import "os"

func main(){
   for index,val:=range os.Args{
      fmt.Println(index,val)
   }
   args_list:=os.Args[1:]
   fmt.Println(args_list)
}