package main

/*
void half(double* f) {
    *f = *f/2;
}
*/
import "C"

import (
	"fmt"
	"math"
	"unsafe"
)

func main() {
	a := float64(math.Pi)
	fmt.Println(a)
	C.half((*C.double)(unsafe.Pointer(&a)))
	fmt.Println(a)
}

// 3.141592653589793
// 1.5707963267948966
