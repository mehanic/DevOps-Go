package __8

import (
	"fmt"
	"sync"
)

func main() {
	wg := sync.WaitGroup{}
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func(n int) {
			defer wg.Done()
			fmt.Println(n)
		}(i)
	}
	wg.Wait()
	fmt.Println("All work done")
}

//---

func main1() {
	ch := make(chan string, 1)
	go func() {
		for _, word := range []string{"hello", "world"} {
			ch <- word
			close(ch)
		}
	}()
	for word := range ch {
		fmt.Println(word)
	}
}


for {
	select {
case v := <-inCh1:
go fmt.Println("received(inCh1): ", v)
case v := <-inCh2:
go fmt.Println("received(inCh2): ", v)
}
}


select {
case s := <-ch:
fmt.Printf("had a string(%s) on the channel\n", s)
default:
fmt.Println("channel was empty")
}

//-------------

func printWords(in1, in2 chan string, exit chan struct{}, wg*sync.WaitGroup) {
defer wg.Done()
for {
         select{
         case <-exit:
			 fmt.Println("exiting")
			 return
         case str := <-in1:
              fmt.Println("in1: ", str)
         case str := <-in2:
              fmt.Println("in2: ", str)
}
}
}

func main2() {
	in1 := make(chan string)
	in2 := make(chan string)
	wg := &sync.WaitGroup{}
	exit := make(chan struct{})
	wg.Add(1)
	go printWords(in1, in2, exit, wg)
	in1 <- "hello"
	in2 <- "world"
	close(exit)
	wg.Wait()
}