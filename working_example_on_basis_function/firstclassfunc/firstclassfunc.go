package main
import "fmt"


func add(myfunc func(a,b int64)int64,a,b int64){
   fmt.Println(myfunc(a,b))
}

func main(){
   test:=func(a,b int64)int64{
      return a+b
   }
   add(test,5,6)
}