package main
import "fmt"
import "sync"

func main(){
  var wg sync.WaitGroup
  func1:=func(wg *sync.WaitGroup){
    defer wg.Done()
    for i:=0;i<10;i++{
       fmt.Println("hello")
    }
  }
  func2:=func(wg *sync.WaitGroup){
    defer wg.Done()
    for i:=0;i<10;i++{
      fmt.Println("bello")
    }
  }
  wg.Add(1)
  go func1(&wg)
  wg.Add(1)
  go func2(&wg)
  wg.Wait()
}