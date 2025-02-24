package main
import "fmt"

func hello(done chan bool){
  fmt.Println("Hello World")
  done<-true
  close(done)
}

func main(){
  done:=make(chan bool)
  go hello(done)
  <-done
  fmt.Println("Main function")
}