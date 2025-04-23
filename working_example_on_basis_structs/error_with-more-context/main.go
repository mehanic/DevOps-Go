package main

import (
	"fmt"
	"log"
	"math"
)

type NorgateMathError struct {
	lat, long string
	err       error
}

// implement Error()
func (n *NorgateMathError) Error() string {
	return fmt.Sprintf("a norgate math error occured: %v %v %v", n.lat, n.long, n.err)
}

func main() {
	v, err := sqrt(-10)
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Println(v)
}

func sqrt(f float64) (float64, error) {
	if f < 0 {
		err := fmt.Errorf("norgate math redux: square root of negative number: %v", f)
		return 0, &NorgateMathError{"40.2323 N", "49.12313 W", err}
	}
	return math.Sqrt(f), nil
}

// 2025/02/14 15:17:44 a norgate math error occured: 40.2323 N 49.12313 W norgate math redux: square root of negative number: -10
// exit status
