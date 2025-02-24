package main

import "fmt"

func main() {
	var counters = map[int]int{}
	for i := 0; i < 5; i++ {
		go func(counters map[int]int, th int) {
			for j := 0; j < 5; j++ {
				counters[th*10+j]++
			}
		}(counters, i)
	}
	fmt.Scanln()
	fmt.Println("counters result", counters)
}

//?
// fatal error: concurrent map writes
// fatal error: concurrent map writes

// goroutine 19 [running]:
// main.main.func1(...)
// 	/home/mehanic/structure/concurrency_i_goroutines/goroutine-race_1/race_1.go:10
// created by main.main in goroutine 1
// 	/home/mehanic/structure/concurrency_i_goroutines/goroutine-race_1/race_1.go:8 +0x34

// goroutine 1 [syscall]:
// syscall.Syscall(0x0, 0x0, 0xc00012a140, 0x1)
