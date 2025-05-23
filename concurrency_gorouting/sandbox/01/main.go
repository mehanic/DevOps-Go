package main

import (
	"fmt"
	"log"
	"math/rand"
	"sync"
	"time"
)

func main() {
	taskNum := rand.Intn(10) + 1
	fmt.Printf("%d tasks\n", taskNum)

	var wg sync.WaitGroup
	wg.Add(taskNum)
	for i := 1; i <= taskNum; i++ {
		go task(&wg, i)
	}
	wg.Wait()
}

func task(wg *sync.WaitGroup, taskId int) {
	defer wg.Done()
	timeRequired := rand.Intn(10) + 1
	time.Sleep(time.Duration(timeRequired) * time.Second)
	log.Printf("task%d done.(%dsec)", taskId, timeRequired)
}

// ╰─λ go run main.go                                                                                                               0 (0.001s) < 14:06:55
// 8 tasks
// 2024/10/04 14:07:01 task8 done.(3sec)
// 2024/10/04 14:07:02 task1 done.(4sec)
// 2024/10/04 14:07:04 task5 done.(6sec)
// 2024/10/04 14:07:06 task6 done.(8sec)
// 2024/10/04 14:07:07 task4 done.(9sec)
// 2024/10/04 14:07:07 task3 done.(9sec)
// 2024/10/04 14:07:08 task7 done.(10sec)
// 2024/10/04 14:07:08 task2 done.(10sec)
