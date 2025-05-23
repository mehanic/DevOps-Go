package main

import (
	"fmt"
	"sync"
	"time"
)

var wg sync.WaitGroup

func main() {
	wg.Add(2)
	go foo()
	go bar()
	wg.Wait()
}

func foo() {
	for i := 0; i < 45; i++ {
		fmt.Println("Foo: ", i)
		time.Sleep(time.Duration(3 * time.Millisecond))
	}
	wg.Done()
}
func bar() {
	for i := 0; i < 45; i++ {
		fmt.Println("bar: ", i)
		time.Sleep(time.Duration(20 * time.Millisecond))
	}
	wg.Done()
}

// Executing concurrent method you need to define wait group and add the number of method to it and
//ask the wait group to wait. Also you need to put waitgroup.Done in those method to take those method from wait group once completed.
//when wg.done is excecuted it will take that method from wait group. wg.wait becomes 0 , the program will execute the main func

// concurrency - idependenty executing lots of things at once
// parallelism - simulatenous execution of (possibly related) computation

//go run -race main.go  ---> to check the race condition

// ╰─λ go run main.go                                                                        0 (0.001s) < 13:02:43
// bar:  0
// Foo:  0
// Foo:  1
// Foo:  2
// Foo:  3
// Foo:  4
// Foo:  5
// bar:  1
// Foo:  6
// Foo:  7
// Foo:  8
// Foo:  9
// Foo:  10
// Foo:  11
// Foo:  12
// bar:  2
// Foo:  13
// Foo:  14
// Foo:  15
// Foo:  16
// Foo:  17
// Foo:  18
// bar:  3
// Foo:  19
// Foo:  20
// Foo:  21
// Foo:  22
// Foo:  23
// Foo:  24
// bar:  4
// Foo:  25
// Foo:  26
// Foo:  27
// Foo:  28
// Foo:  29
// Foo:  30
// bar:  5
// Foo:  31
// Foo:  32
// Foo:  33
// Foo:  34
// Foo:  35
// Foo:  36
// bar:  6
// Foo:  37
// Foo:  38
// Foo:  39
// Foo:  40
// Foo:  41
// Foo:  42
// bar:  7
// Foo:  43
// Foo:  44
// bar:  8
// bar:  9
// bar:  10
// bar:  11
// bar:  12
// bar:  13
// bar:  14
// bar:  15
// bar:  16
// bar:  17
// bar:  18
// bar:  19
// bar:  20
// bar:  21
// bar:  22
// bar:  23
// bar:  24
// bar:  25
// bar:  26
// bar:  27
// bar:  28
// bar:  29
// bar:  30
// bar:  31
// bar:  32
// bar:  33
// bar:  34
// bar:  35
// bar:  36
// bar:  37
// bar:  38
// bar:  39
// bar:  40
// bar:  41
// bar:  42
// bar:  43
// bar:  44
