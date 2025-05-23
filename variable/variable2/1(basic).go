package main
import "fmt"
var global string ="I am global!"
func main() {
    //var str string = "My first program"
    //fmt.Println(str)

    //несколько
    // var i,j,k int
    // var il,jl,kl = true,2.3,"apple"
    local:="I am local"

    fmt.Println(local)
    fmt.Println(global)
    fmt.Printf("binary: %b, %b",64,102)
}