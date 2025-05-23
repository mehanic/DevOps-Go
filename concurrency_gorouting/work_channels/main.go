package main

import (
	"fmt"
	"runtime"
	"time"
)

func Worker(funcs []func() error, nG int, nErrors int) {
	chErrors := make(chan error, nErrors)
	chEnd := make(chan int)
	if nG < 1 {
		nG = 1
	}
	if nG > len(funcs) {
		nG = len(funcs)
	}
	chunk := len(funcs) / nG
	for i := 0; i < nG; i++ {
		go func(extI int) {
			// var err error
			for {
				select {
				case <-chEnd:
					fmt.Printf("Stop gorutine %d\n", extI+1)
					return
				default:
					for k := 0; k < chunk; k++ {
						j := chunk*extI + k
						if j > len(funcs) {
							return
						}
						if err := funcs[j](); err != nil {
							chErrors <- err
						}
					}
				}
			}
		}(i)
	}
	closed := false
	timer := time.NewTimer(24 * time.Hour)
	for runtime.NumGoroutine() > 1 {
		select {
		case e := <-chErrors:
			if closed == false {
				fmt.Printf("Reading an error: Value %v\n", e)
				close(chEnd)
				closed = true
				timer = time.NewTimer(10 * time.Second)
			}
		case <-timer.C:
			fmt.Printf("Time out 10s\n")
			return
		default:
		}
	}
}
func f1() error {
	for i := 0; i < 10; i++ {
		fmt.Printf("from f1 %d\n", i)
		time.Sleep(time.Second)
		if i > 8 {
			fmt.Printf("error from 1 \n")
			return fmt.Errorf("error f1 %d", i)
		}
	}
	return nil
}

func f2() error {
	for i := 0; i < 10; i++ {
		fmt.Printf("from f2 %d\n", i)
		time.Sleep(2 * time.Second)
		if i > 5 {
			fmt.Printf("error from 2 \n")
			return fmt.Errorf("error f2 %d", i)
		}
	}
	return nil
}

func f3() error {
	for i := 0; i < 10; i++ {
		fmt.Printf("from f3 %d\n", i)
		time.Sleep(3 * time.Second)
		if i > 3 {
			fmt.Printf("error from 3 \n")
			return fmt.Errorf("error f3 %d", i)
		}
	}
	return nil
}

func f4() error {
	for i := 0; i < 10; i++ {
		fmt.Printf("from f4 %d\n", i)
		time.Sleep(300 * time.Millisecond)
	}
	fmt.Printf("f4 is done \n")
	return nil
}

func main() {
	var funcs []func() error
	funcs = append(funcs, f4, f2, f3, f1)
	Worker(funcs, 2, 2)
}

// from f4 0
// from f3 0
// from f4 1
// from f4 2
// from f4 3
// from f4 4
// from f4 5
// from f4 6
// from f4 7
// from f4 8
// from f4 9
// from f3 1
// f4 is done
// from f2 0
// from f2 1
// from f3 2
// from f2 2
// from f3 3
// from f2 3
// from f2 4
// from f3 4
// from f2 5
// error from 3
// from f1 0
// Reading an error: Value error f3 4
// from f2 6
// from f1 1
// from f1 2
// error from 2
// Stop gorutine 1
// from f1 3
// from f1 4
// from f1 5
// from f1 6
// from f1 7
// from f1 8
// from f1 9
// Time out 10s
