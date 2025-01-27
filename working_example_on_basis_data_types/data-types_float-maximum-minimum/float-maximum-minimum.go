package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Printf("min float64: %.50e\n", math.SmallestNonzeroFloat64)
	fmt.Printf("max float64: %.50e\n", math.MaxFloat64)

	fmt.Printf("min float32: %.50e\n", math.SmallestNonzeroFloat32)
	fmt.Printf("max float32: %.50e\n", math.MaxFloat32)
}

// min float64: 4.94065645841246544176568792868221372365059802614325e-324
// max float64: 1.79769313486231570814527423731704356798070567525845e+308
// min float32: 1.40129846432481707092372958328991613128026194187652e-45
// max float32: 3.40282346638528859811704183484516925440000000000000e+38
