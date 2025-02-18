package main
import "fmt"
import "encoding/json"

type Person struct{
  Name string `json:"name"`
  Id int64 `json:"id"`
  Email string `json:"email"`
}

func main(){
  p1:=Person{"bapan",1,"bapan@gmail.com"}
  p2:=Person{"bapan21",2,"bapan21@gmail.com"}
  byte1,err:=json.Marshal(p1)
  if err!=nil{
     fmt.Println(err)
  }
  fmt.Println(string(byte1))
  byte2,err:=json.Marshal(p2)
  if err!=nil{
    fmt.Println(err)
  }
  fmt.Println(string(byte2))
  p3:=Person{}
  data:=json.RawMessage(string(byte1))
  byte3,err:=data.MarshalJSON()
  if err!=nil{
    fmt.Println(err)
  }
  err=json.Unmarshal(byte3,&p3)
  if err!=nil{
    fmt.Println(err)
  }
  fmt.Println(p3)
  byte4,err:=json.Marshal(p3)
  fmt.Println(string(byte4))
}