package main

import (
	"context"
	"fmt"
	"time"
	//"github.com/PacktPublishing/Go-Programming-Cookbook-Second-Edition/chapter10/channels"
)

func main() {
	ch := make(chan string)
	done := make(chan bool)

	ctx := context.Background()
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()

	go Printer(ctx, ch)
	go Sender(ch, done)

	time.Sleep(2 * time.Second)
	done <- true
	cancel()
	time.Sleep(3 * time.Second)
}

// Printer will print anything sent on the ch chan
// and will print tock every 200 milliseconds
// this will repeat forever until a context is
// Done, i.e. timed out or cancelled
func Printer(ctx context.Context, ch chan string) {
	t := time.Tick(200 * time.Millisecond)
	for {
		select {
		case <-ctx.Done():
			fmt.Println("printer done.")
			return
		case res := <-ch:
			fmt.Println(res)
		case <-t:
			fmt.Println("tock")
		}
	}
}

// Sender sends "tick"" on ch until done is
// written to, then it sends "sender done."
// and exits
func Sender(ch chan string, done chan bool) {
	t := time.Tick(100 * time.Millisecond)
	for {
		select {
		case <-done:
			ch <- "sender done."
			return
		case <-t:
			ch <- "tick"
		}
	}
}

// tick
// tock
// tick
// tick
// tick
// tock
// tick
// tock
// tick
// tick
// tock
// tick
// tick
// tick
// tock
// tick
// tock
// tick
// tick
// tick
// tock
// tick
// tock
// tick
// tick
// tock
// tick
// tick
// tock
// tick
// printer done.
