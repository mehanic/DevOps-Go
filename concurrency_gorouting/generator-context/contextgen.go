package main

import (
	"context"
	"fmt"
	"time"
)

type genInt64 struct {
	ch chan int64
}

func (g genInt64) Next() int64 {
	return <-g.ch
}

func NewGenInt64(ctx context.Context) genInt64 {
	g := genInt64{ch: make(chan int64)}
	go func() {
		for i := int64(0); ; i++ {
			select {
			case g.ch <- i:
				// do nothing
			case <-ctx.Done():
				close(g.ch)
				return
			}
		}
	}()
	return g
}

func main() {
	ctx, cancel := context.WithTimeout(context.Background(), time.Millisecond*100)
	defer cancel()
	g := NewGenInt64(ctx)
	for i := range g.ch {
		go func(i int64) {
			fmt.Println(i, g.Next())
		}(i)
	}
	time.Sleep(time.Second)
}

// 71431 71433
// 71455 71457
// 71452 71458
// 71451 71453
// 71456 71460
// 71459 71462
// 71461 71466
// 71463 71465
// 71467 71469
// 71464 71470
// 71468 71472
// 71471 0
