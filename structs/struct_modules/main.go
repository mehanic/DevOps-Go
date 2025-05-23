package main

import (
	"fmt"
    "figures/shape"
)

func main(){
    circle:=shape.NewCircle(4)
    rectangle:=shape.NewRectangle(4,7)

    fmt.Println(circle.CountSquare())
    fmt.Println(rectangle.CountSquare())
}