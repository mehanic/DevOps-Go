package shape

import "math"

type FigureInterface interface{
    CountSquare()float64
}

type Rectangle struct{
    Width int
    Height int
}

func NewRectangle(w,h int) *Rectangle{
    return &Rectangle{
        Width:w,
        Height:h,
    }
}

func (f *Rectangle) CountSquare()float64{
    return float64(f.Width*f.Height)
}

type Circle struct{
    Radius int
}

func NewCircle(r int) *Circle{
    return &Circle{
        Radius:r,
    }
}

func (f *Circle) CountSquare()float64{
    return math.Pow(float64(f.Radius),2)*math.Pi
}