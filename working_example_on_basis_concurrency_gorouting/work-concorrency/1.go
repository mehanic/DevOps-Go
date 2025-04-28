package chapter9

import (
	"fmt"
	"runtime"
	"sync"
	"sync/atomic"
	"time"
)

var wg sync.WaitGroup
var mu sync.Mutex

// -----------------------------------------------
// Annotation 1: Simple Goroutine Execution with WaitGroup
func Annotation1() {

	wg.Add(2)

	go function1()
	go function2()

	wg.Wait()
}

func function1() {
	for i := 0; i < 100; i++ {
		fmt.Println("Function 1:", i)
		time.Sleep(2000)
	}
	wg.Done()
}

func function2() {
	for i := 0; i < 100; i++ {
		fmt.Println("Function 2:", i)
	}
	wg.Done()
}

// -----------------------------------------------
// Annotation 2: Using Mutex to Avoid Race Condition
// - Simple example of how to prevent race conditions
// - Run with: go run -race main.go
func Annotation2() {

	counter := 0
	goroutines := 1000

	wg.Add(goroutines)

	fmt.Println("CPUs:", runtime.NumCPU())

	for i := 0; i < goroutines; i++ {
		go func() {
			mu.Lock()
			v := counter
			runtime.Gosched()
			v++
			counter = v
			mu.Unlock()
			wg.Done()
		}()
		fmt.Println("Goroutines:", runtime.NumGoroutine())
	}
	wg.Wait()
	fmt.Println("Goroutines:", runtime.NumGoroutine())
	fmt.Println("Final Value:", counter)
}

// -----------------------------------------------
// Annotation 3: Using Atomic to Avoid Race Conditions
func Annotation3() {

	var counter int64
	counter = 0
	goroutines := 1000

	wg.Add(goroutines)

	fmt.Println("CPUs:", runtime.NumCPU())

	for i := 0; i < goroutines; i++ {
		go func() {
			atomic.AddInt64(&counter, 1)
			runtime.Gosched()
			fmt.Println("Counter:\t", atomic.LoadInt64(&counter))
			wg.Done()
			wg.Done()
		}()
	}
	wg.Wait()
	fmt.Println("Final Value:", counter)
}
