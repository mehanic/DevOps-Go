package main
import "flag"
import "fmt"

func main(){
   word:=flag.String("word","foo","a string")
   num:=flag.Int("numb",42,"an int")
   boolptr:=flag.Bool("con",false,"a bool")
   var svar string
   flag.StringVar(&svar,"svar","bar","a string var")
   flag.Parse()
   fmt.Println("word:",*word)
   fmt.Println("numb:",*num)
   fmt.Println("con:",*boolptr)
   fmt.Println("svar:",svar)
}