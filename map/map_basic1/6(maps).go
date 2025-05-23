package main

import "fmt"

func main(){
//     ages:= make(map[string]int)
    ages:=map[string]int{
        "Ann":20,
        "Tema":22,
        "Dad":48,
        "Kira":7,
    }
//     age,exist:=ages["Kira"]
    ages["Kira"]=9
    var ages2 map[string]int
//     delete(ages,"Kira")
//     ages["Ann"]=20
//     ages["Tema"]=22
//     ages["Dad"]=48
//     fmt.Println(ages)
//     fmt.Println("ages == nil->",ages==nil)
    for key,value:= range ages{
        fmt.Printf("%s - %d\n",key,value)
    }
    fmt.Println()
    for key,value:= range ages2{
            fmt.Printf("%s - %d\n",key,value)
    }
    fmt.Println()
    ages2=ages
    delete(ages2,"Kira")
    fmt.Println()
    for key,value:= range ages{
        fmt.Printf("%s - %d\n",key,value)
    }
    fmt.Println()
    for key,value:= range ages2{
            fmt.Printf("%s - %d\n",key,value)
    }
}