package main

import (
	"fmt"
	"sync"
)

func main() {
	var wg sync.WaitGroup
	ch1 := make(chan string, 5)
	ch2 := make(chan string, 5)
	task1 := []string{"hello", "bello", "mello", "tello", "gello"}
	task2 := []string{"hillo", "billo", "millo", "tillo", "gillo"}
	func1 := func(wg *sync.WaitGroup, tasks ...string) {
		fmt.Println(tasks)
		defer wg.Done()
		for _, t := range tasks {
			ch1 <- t
		}
		close(ch1)
	}
	func2 := func(wg *sync.WaitGroup, tasks ...string) {
		fmt.Println(tasks)
		defer wg.Done()
		for _, t := range tasks {
			ch2 <- t
		}
		close(ch2)
	}
	func3 := func(wg *sync.WaitGroup) {
		var msg1 string
		task1_complete := false
		task2_complete := false
		defer wg.Done()
		for {
			select {
			case msg1 = <-ch1:
				if msg1 != "" {
					fmt.Println(msg1)
				}
				for tsk := range ch1 {
					fmt.Println(tsk)
				}
				task1_complete = true
			case msg1 = <-ch2:
				if msg1 != "" {
					fmt.Println(msg1)
				}
				for tsk := range ch2 {
					fmt.Println(tsk)
				}
				task2_complete = true
			}
			if task1_complete && task2_complete {
				break
			}
		}
	}

	wg.Add(1)
	go func1(&wg, task1...)
	wg.Add(1)
	go func2(&wg, task2...)
	wg.Add(1)
	go func3(&wg)
	wg.Wait()
}
