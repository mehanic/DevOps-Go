package main
import "os"
import "fmt"
import "bytes"


func main(){
   f,err:=os.OpenFile("test1.txt",os.O_CREATE|os.O_WRONLY,0777)
   if err!=nil{
      fmt.Println("Error Occured",err)
   }
   defer f.Close()
   b:=bytes.NewBuffer(make([]byte,16))
   str1:=[]string{"hello","bello","mello","mello1","bello200"}
   for _,v:=range str1{
      fmt.Fprintf(b,"%s",v)
      b.WriteRune('\n')
      if _,err:=b.WriteTo(f);err!=nil{
         fmt.Println("Error occured",err)
      }
   }
}