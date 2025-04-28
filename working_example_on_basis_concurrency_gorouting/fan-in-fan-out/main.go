package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	borings := fanIn(fanOut("Clem", "John", "Bob")...)
	for i := 0; i < 20; i++ {
		fmt.Println(<-borings)
	}
	fmt.Println("leaving")
}

func boring(name string) <-chan string {
	out := make(chan string)
	go func() {
		for i := 0; ; i++ {
			msg := fmt.Sprintf("%d %s is boring", i, name)
			out <- msg
			time.Sleep(time.Duration(rand.Intn(3) * int(time.Second)))
		}
	}()
	return out
}

func fanOut(names ...string) []<-chan string {
	outChs := make([]<-chan string, len(names))
	for index, name := range names {
		outChs[index] = boring(name)
	}
	return outChs
}

func fanIn(channels ...<-chan string) <-chan string {
	out := make(chan string)

	for _, channel := range channels {
		go func(c <-chan string) {
			for {
				out <- <-c
			}
		}(channel)
	}

	return out
}

// 0 Bob is boring
// 0 John is boring
// 1 John is boring
// 2 John is boring
// 0 Clem is boring
// 1 Bob is boring
// 3 John is boring
// 2 Bob is boring
// 1 Clem is boring
// 2 Clem is boring
// 3 Clem is boring
// 3 Bob is boring
// 4 Bob is boring
// 4 John is boring
// 4 Clem is boring
// 5 Clem is boring
// 6 Clem is boring
// 7 Clem is boring
// 8 Clem is boring
// 5 John is boring
// leaving
