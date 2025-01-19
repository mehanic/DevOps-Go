package utils

import "math"

func EpsEqual(eps float64, f1, f2 float64) bool {
	return math.Abs(f1-f2) <= eps
}
