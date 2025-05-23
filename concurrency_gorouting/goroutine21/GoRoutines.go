package main

import (
	"fmt"
	"time"
)

func count(id int) {
	for i := 0; i < 10; i++ {
		fmt.Println(id, ":", i)

		time.Sleep(time.Millisecond * 1000)
	}
}

func main() {
	for i := 0; i < 10; i++ {
		go count(i)
	}

	time.Sleep(time.Millisecond * 11000)
}

// ╰─λ go run GoRoutines.go                                                                   0 (0.000s) < 15:50:02
// 9 : 0
// 4 : 0
// 5 : 0
// 6 : 0
// 0 : 0
// 1 : 0
// 2 : 0
// 7 : 0
// 8 : 0
// 3 : 0
// 3 : 1
// 2 : 1
// 8 : 1
// 6 : 1
// 9 : 1
// 1 : 1
// 5 : 1
// 0 : 1
// 4 : 1
// 7 : 1
// 4 : 2
// 6 : 2
// 8 : 2
// 9 : 2
// 0 : 2
// 1 : 2
// 5 : 2
// 3 : 2
// 7 : 2
// 2 : 2
// 8 : 3
// 5 : 3
// 3 : 3
// 9 : 3
// 0 : 3
// 1 : 3
// 4 : 3
// 6 : 3
