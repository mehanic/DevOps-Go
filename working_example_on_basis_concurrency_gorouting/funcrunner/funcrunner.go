package main

import (
	"fmt"
	"time"
)

// Run concurrent tasks
func Run(tasks []func() error, N int, M int) error {
	// start task pool
	taskChannel := make(chan int, len(tasks))
	// error pool
	errorChannel := make(chan error, len(tasks))

	// routine count
	routineNum := N
	if len(tasks) < N {
		routineNum = len(tasks)
	}

	// quit task pool
	quitChannel := make(chan struct{}, routineNum)

	defer func() {
		close(errorChannel)
		close(quitChannel)
	}()

	// feel pool with tasks id-s
	for i := 0; i < len(tasks); i++ {
		taskChannel <- i
	}
	close(taskChannel)

	for i := 0; i < routineNum; i++ {

		go func() {

			defer func() {
				quitChannel <- struct{}{}
			}()

			for {
				if len(errorChannel) >= M {
					return
				}
				taskID, ok := <-taskChannel
				if !ok {
					return
				}
				err := tasks[taskID]()
				if err != nil {
					errorChannel <- err
				}
			}

		}()

	}

	// wait for goroutine finish
	for i := 0; i < routineNum; i++ {
		<-quitChannel
	}

	if len(errorChannel) >= M {
		return fmt.Errorf("Max number of errors reached")
	}

	return nil
}


// Example tasks
func task1() error {
	fmt.Println("Running Task 1")
	time.Sleep(1 * time.Second)
	return nil
}

func task2() error {
	fmt.Println("Running Task 2")
	time.Sleep(2 * time.Second)
	return nil
}

func taskWithError() error {
	fmt.Println("Running Task with Error")
	time.Sleep(1 * time.Second)
	return fmt.Errorf("Task failed")
}

func main() {
	// Define tasks
	tasks := []func() error{
		task1,
		task2,
		taskWithError,
		task1,
		task2,
	}

	// Run tasks with a maximum of 2 workers and allow up to 1 error
	err := Run(tasks, 2, 1)
	if err != nil {
		fmt.Printf("Error: %v\n", err)
	} else {
		fmt.Println("All tasks completed successfully")
	}
}


// Running Task 1
// Running Task 2
// Running Task with Error
// Running Task 1
// Error: Max number of errors reached
