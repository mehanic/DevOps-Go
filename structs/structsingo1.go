package main
import "fmt"
import "math"

func main() {
    
    rect1 := Rectangle{0,50,10,10}
    fmt.Println("Rectangle is ", rect1.width, "wide")

    fmt.Println("Area of rectangle =", rect1.area())
    rect := Rectangle2 {20,50}
    circle := Circle{5}
    
    fmt.Println("Rectangle Area = ", getArea(rect))
    fmt.Println("Circle area = ", getArea(circle))
}

type Rectangle struct {
    
    leftX float64
    topY float64
    height float64
    width float64
}

func (rect *Rectangle) area() float64 {
    return rect.width * rect.height
}

type Shape interface {
    area() float64
}

type Rectangle2 struct {
    height float64
    width float64
}

type Circle struct {
    radius float64
}

func (r Rectangle2) area() float64 {
    return r.height * r.width
}

func (c Circle) area() float64 {
    return math.Pi * math.Pow(c.radius, 2)
}

func getArea(shape Shape) float64 {
    return shape.area()
}