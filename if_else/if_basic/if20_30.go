package main

import (
	"fmt"
	"math"
)

// Function If20 determines which of the two values (b or c) is closer to a.
func If20(a, b, c float64) (closer string, distance float64) {
	diffB := math.Abs(a - b)
	diffC := math.Abs(a - c)
	if diffB < diffC {
		closer = "b"
		distance = diffB
	} else {
		closer = "c"
		distance = diffC
	}
	return
}

// Function If21 determines the quadrant of a point based on x and y.
func If21(x, y float64) (result int16) {
	if x == 0 && y == 0 {
		result = 0
	} else if x == 0 {
		result = 2
	} else if y == 0 {
		result = 1
	} else {
		result = 3
	}
	return
}

// Function If22 determines which quadrant the point (x, y) is in.
func If22(x, y float64) (quadrant int16) {
	if x > 0 && y > 0 {
		quadrant = 1
	} else if x < 0 && y > 0 {
		quadrant = 2
	} else if x < 0 && y < 0 {
		quadrant = 3
	} else if x > 0 && y < 0 {
		quadrant = 4
	}
	return
}

// Function If23 determines the fourth vertex of a parallelogram.
func If23(x1, y1, x2, y2, x3, y3 int16) (x4, y4 int16) {
	if x2 == x3 {
		x4 = x1
	} else if x3 == x1 {
		x4 = x2
	} else if x1 == x2 {
		x4 = x3
	}
	if y2 == y3 {
		y4 = y1
	} else if y3 == y1 {
		y4 = y2
	} else if y1 == y2 {
		y4 = y3
	}
	return
}

// Function If24 calculates a value based on the input x.
func If24(x float64) (result float64) {
	if x > 0 {
		result = 2 * math.Sin(x)
	} else {
		result = 6 - x
	}
	return
}

// Function If25 calculates a value based on the input x.
func If25(x int16) (result int16) {
	if x < -2 || x > 2 {
		result = 2 * x
	} else {
		result = -3 * x
	}
	return
}

// Function If26 calculates a piecewise function based on x.
func If26(x float64) (result float64) {
	if x <= 0 {
		result = -x
	} else if x > 0 && x < 2 {
		result = x * x
	} else {
		result = 4
	}
	return
}

// Function If27 checks if x is negative, even, or odd.
func If27(x float64) (result float64) {
	integerPart := int(x)
	if x < 0 {
		result = 0
	} else if integerPart%2 == 0 {
		result = 1
	} else {
		result = -1
	}
	return
}

// Function If28 determines the number of days in a given year.
func If28(year int16) (days int16) {
	if (year%4 == 0 && year%100 != 0) || (year%400 == 0) {
		days = 366
	} else {
		days = 365
	}
	return
}

// Function If29 determines the characteristics of a number x.
func If29(x int16) (description string) {
	if x < 0 && -x%2 == 1 {
		description = "Negative odd"
	} else if x < 0 && -x%2 == 0 {
		description = "Negative even"
	} else if x > 0 && x%2 == 0 {
		description = "Positive even"
	} else if x > 0 && x%2 == 1 {
		description = "Positive odd"
	} else {
		description = "Zero"
	}
	return
}

// Function If30 determines the characteristics of a number based on its parity and digit count.
func If30(x int16) (description string) {
	if x%2 == 1 && x <= 9 {
		description = "Odd single-digit"
	} else if x%2 == 0 && x <= 9 {
		description = "Even single-digit"
	} else if x%2 == 1 && x <= 99 && x >= 10 {
		description = "Odd two-digit"
	} else if x%2 == 0 && x <= 99 && x >= 10 {
		description = "Even two-digit"
	} else if x%2 == 1 && x <= 999 && x >= 100 {
		description = "Odd three-digit"
	} else if x%2 == 0 && x <= 999 && x >= 100 {
		description = "Even three-digit"
	}
	return
}

func main() {
	// Example for If20:
	fmt.Println(If20(10, 7, 12)) // Output: ("b", 3)

	// Example for If21:
	fmt.Println(If21(0, 5)) // Output: 2

	// Example for If22:
	fmt.Println(If22(-3, 4)) // Output: 2

	// Example for If23:
	fmt.Println(If23(0, 0, 1, 0, 0, 1)) // Output: (1, 1)

	// Example for If24:
	fmt.Println(If24(3)) // Output: 1.4112000805986722

	// Example for If25:
	fmt.Println(If25(3)) // Output: 6

	// Example for If26:
	fmt.Println(If26(1.5)) // Output: 2.25

	// Example for If27:
	fmt.Println(If27(4)) // Output: 1

	// Example for If28:
	fmt.Println(If28(2024)) // Output: 366

	// Example for If29:
	fmt.Println(If29(-5)) // Output: "Negative odd"

	// Example for If30:
	fmt.Println(If30(23)) // Output: "Odd two-digit"
}
