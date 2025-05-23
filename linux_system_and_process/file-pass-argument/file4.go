package main
import "fmt"
import "io/ioutil"
import "os"

func main(){
   if len(os.Args)<3{
     fmt.Println("Please provide file path")
     return
   }
   by,err:=ioutil.ReadFile(os.Args[1])
   if err!=nil{
      fmt.Println("Error occured",err)
   }
   err=ioutil.WriteFile(os.Args[2],by,0777)
   if err!=nil{
      fmt.Println("Error occured",err)
   }
}