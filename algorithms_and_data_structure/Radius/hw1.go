package main

import (
	"fmt"
    "errors"
    "math"
)
func main(){
    printAreas(2,4,6)

}

func printAreas(radius,length,width int){
    circleArea,err:= calculateCircleArea(radius)
    recArea,err2:=calculateRectangleArea(length,width)
    if(err!=nil) || (err2!=nil){
        fmt.Println(err.Error())
        return
    }
    fmt.Printf("Radius: %d \n",radius)
    fmt.Printf("Area circ: %f32 \n",circleArea)

    fmt.Println()
    fmt.Println("Area rec: ",recArea)
}

func calculateCircleArea(radius int) (float32,error){
    if radius<=0{
        return float32(0),errors.New("Радиус должен быть строго больше нуля!")
    }
    return float32(radius)*float32(radius)*math.Pi,nil
}

func calculateRectangleArea(length,width int)(int,error){
    if length<=0 || width<=0{
        return int(0),errors.New("Радиус должен быть строго больше нуля!")
    }
    return length*width,nil
}