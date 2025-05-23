package main

import (
	"fmt"
    "errors"
)

const pi = 3.1415
func main(){
    printCircleArea(-2)
    printCircleArea(4)

}

func printCircleArea(radius int){
    circleArea,err:= calculateCircleArea(radius)
    if(err!=nil){
        fmt.Println(err.Error())
        return
    }
    fmt.Printf("Radius: %d \n",radius)
    fmt.Printf("Area: %f32 \n",circleArea)
}

func calculateCircleArea(radius int) (float32,error){
    if radius<=0{
        return float32(0),errors.New("Радиус должен быть строго больше нуля!")
    }
    return float32(radius)*float32(radius)*pi,nil
}