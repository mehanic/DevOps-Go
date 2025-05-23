package main

import "fmt"

const pi = 3.1415
func main(){
    printCircleArea(2)
    printCircleArea(4)

}

func printCircleArea(radius int){

    fmt.Printf("Radius: %d \n",radius)
    circleArea:=calculateCircleArea(radius)
    fmt.Printf("Area: %f32 \n",circleArea)
}

func calculateCircleArea(radius int) float32{
    return float32(radius)*float32(radius)*pi
}