package main

import (
    "fmt"
    "time"
)

func main() {
    // Global variable
    stringVar := fmt.Sprintf("Global Var, %s!", time.Now().Format("20060102"))
    fmt.Println(stringVar)

    // Call local_var function
    localVar() // Call local_var function
}

func localVar() {
    // Local variable
    var localVar string = fmt.Sprintf("Local Var, %s!", time.Now().Format("20060102"))
    fmt.Println(localVar)
}

