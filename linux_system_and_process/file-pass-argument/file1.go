package main
import "fmt"
import "os"

func main(){
  if len(os.Args)<2{
    fmt.Println("Please provide file path")
    return
  }
  f,err:=os.OpenFile(os.Args[1],os.O_CREATE|os.O_WRONLY,0777)
  defer f.Close()
  if err!=nil{
    fmt.Println(err)
    return
  }
  fmt.Fprintf(f,"%s","hello world")
}