package main

import "fmt"

func main(){
    printType(3)
    printType("babababbaba")
    printType([]string{"one","two"})
}

func printType(value interface{}){
//     if _,ok:=value.(int);ok{
//         fmt.Println("тип аргумента int")
//     } else if _,ok:=value.(string); ok{
//         fmt.Println("тип аргумента string")
//     }else{
//         fmt.Println("тип аргумента не int и не string")
//     }
    switch value.(type){
        case int:
            fmt.Println("тип аргумента int")
        case string:
            fmt.Println("тип аргумента string")
        default:
            fmt.Println("тип аргумента не int и не string")
    }
}