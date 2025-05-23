package main

import "fmt"

func main() {
    
    listNums := []float64{1,2,3,4,5,6}

    fmt.Println("Sum :", addThemUp(listNums))
    num1, num2 := next2Vals(5)
    fmt.Println("Next two values", num1, num2)
    fmt.Println(subtractThem(1,2,3,4,5,6))

    //clojures 
    
    num3 := 3
    doubleNum := func() int {
        num3 *= 2
        return num3
    }
    
    fmt.Println(doubleNum())
    fmt.Println(factorial(3))
    
    // defer
    defer printTwo()
    printOne()
    
    fmt.Println(safeDiv(3, 0))
    fmt.Println(safeDiv(3, 2))
    demPanic()
    
    //pointers in go
    x := 0
    changeXVal(x)
    
    fmt.Println("x =", x)
    
    changeXValNow(&x)
    
    fmt.Println("x* = ", x)
    
    fmt.Println(" Memory address of x = ", &x)
    yPtr := new(int)
    changeYValNow(yPtr)
    fmt.Println("y =", *yPtr)
}

func addThemUp(numbers []float64) float64 {
    
    sum := 0.0
    
    for _, val := range numbers {
        
        sum += val
        
    }
    return sum
}

func next2Vals(number int) (int, int) {
    return number+1, number+2
}


func subtractThem(args ...int) int {
    
    finalValue := 0
    for _, value := range args {
        
        finalValue -= value
    }
    
    return finalValue
}

func factorial(num int) int {
    if num == 0 {
        return 1
    }
    
    return num * factorial(num -1)
}

// defers

func printOne() {
    fmt.Println("Defer 1")
}

func printTwo() {
    fmt.Println("Defer 2")
}

func safeDiv(num1, num2 int) int {
   // if i comment this out it will throw an error
    defer func() {
        fmt.Println(recover())
    }()
    
    solution := num1 / num2
    return solution
}

//panic

func demPanic() {
    defer func() {
        fmt.Println(recover())
    }()
    
    panic("PANIC")
}


// pointers in go

func changeXVal(x int) {
    x = 2
}

func changeXValNow(x *int) {
    *x = 2
}

func changeYValNow(yPtr *int) {
    *yPtr = 100
}