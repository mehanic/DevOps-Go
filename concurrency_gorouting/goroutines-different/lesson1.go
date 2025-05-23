package main
import "fmt"
import "flag"


func main(){
   enviroment:=flag.String("env","qe","a string")
   iteration:=flag.Int("itr",2,"an int")
   blvar:=flag.Bool("bl",false,"a boolean")
   flag.Parse()
   var str1 string
   flag.StringVar(&str1,"str1","hello bello","a string var")
   fmt.Println(*enviroment)
   fmt.Println(*iteration)
   fmt.Println(*blvar)
   fmt.Println(str1)
}