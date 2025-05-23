package main
import "fmt"
import "os"

func main(){
  if len(os.Args)<3{
    fmt.Println("Please provide valid source and destination path")
    return
  }
  byts,err:=os.ReadFile(os.Args[1])
  if err!=nil{
    fmt.Println(err)
    return
  }
  err=os.WriteFile(os.Args[2],byts,0777)
  if err!=nil{
      fmt.Println(err)
      return
  }
  return
}