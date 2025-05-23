package main
import "fmt"
import "os"

func main(){
  file,err:=os.OpenFile("test1.txt",os.O_CREATE|os.O_WRONLY,0777)
  if err!=nil{
    fmt.Println(err)
  }
  defer file.Close()
  fmt.Fprintf(file,"hello bello")
}