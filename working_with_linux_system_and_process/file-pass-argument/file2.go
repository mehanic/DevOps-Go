package main
import "fmt"
import "os"
import "io"


func main(){
   if len(os.Args)<3{
      fmt.Println("Please provide source and destination file names")
      return
   }
   source,err:=os.Open(os.Args[1])
   defer source.Close()
   if err!=nil{
      fmt.Println(err)
      return
   }
   destination,err:=os.OpenFile(os.Args[2],os.O_CREATE|os.O_WRONLY,0777)
   defer destination.Close()
   if err!=nil{
     fmt.Println(err)
     return
   }
   n,err:=io.Copy(destination,source)
   if err!=nil && err!=io.EOF{
      fmt.Println(err)
   }
   fmt.Printf("No of bytes copied is %d\n",n)
}