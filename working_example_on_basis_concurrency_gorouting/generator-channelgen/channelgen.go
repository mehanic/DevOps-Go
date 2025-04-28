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
	g := NewGenInt64(context.Background())
	for i := 0; i < 100; i++ {
		go func(i int) {
			fmt.Println(i, g.Next())
		}(i)
	}
	time.Sleep(time.Second)
}

// /─λ go run channelgen.go                                                                                                         0 (0.001s) < 14:04:27
// 9 0
// 0 1
// 1 2
// 2 3
// 3 4
// 4 5
// 5 6
// 6 7
// 7 8
// 8 9
// 25 10
// 76 69
// 34 20
// 35 21
// 36 22
// 37 23
// 38 24
// 39 25
// 40 26
// 65 27
// 99 70
// 80 73
// 79 72
// 33 19
// 10 52
// 42 28
// 43 29
// 44 30
// 32 18
// 23 91
// 17 74
// 50 36
// 55 41
// 59 45
// 86 80
// 60 46
// 30 16
// 98 99
// 31 17
// 19 87
// 45 31
// 11 53
// 20 88
// 12 54
// 21 89
// 15 55
// 22 90
// 14 56
// 64 50
// 13 51
// 46 32
// 85 79
// 47 33
// 78 71
// 48 34
// 24 92
// 49 35
// 92 93
// 16 57
// 93 94
// 77 58
// 94 95
// 66 59
// 95 96
// 67 60
// 96 97
// 68 61
// 97 98
// 69 62
// 70 63
// 57 43
// 71 64
// 51 37
// 52 38
// 41 11
// 53 39
// 54 40
// 81 75
// 82 76
// 26 12
// 58 44
// 83 77
// 72 65
// 27 13
// 73 66
// 56 42
// 74 67
// 75 68
// 88 82
// 84 78
// 87 81
// 28 14
// 90 84
// 62 48
// 63 49
// 89 83
// 29 15
// 91 85
// 61 47
// 18 86
