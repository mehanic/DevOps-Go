package main
import "fmt"
import "os"
import "bytes"

func main(){
   f,err:=os.OpenFile(os.Args[1],os.O_CREATE|os.O_WRONLY,0777)
   if err!=nil{
      fmt.Println(err)
   }
   defer f.Close()
   b:=bytes.NewBuffer(make([]byte,20))
   str1:=[]string{"hello1","bello1","tello1","mello2","chello3"}
   for _,v:=range str1{
      fmt.Fprintf(b,"%s",v)
      b.WriteRune('\n')
      if _,err:=b.WriteTo(f);err!=nil{
         fmt.Println(err)
      }
   }
}