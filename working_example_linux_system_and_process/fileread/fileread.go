package main
import "fmt"
import "os"

func main(){
   if len(os.Args)<2{
      fmt.Println("Please provide the filename")
      return
   }
   b,err:=os.ReadFile(os.Args[1])
   if err!=nil{
      fmt.Println("Error occured",err)
      return
   }
   fmt.Println(string(b))
}