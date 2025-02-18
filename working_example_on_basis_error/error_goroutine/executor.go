package main

import (
	"errors"
	"fmt"
	"sync"
)

var (
	ErrIncorrectErrorsSize  = errors.New("max errors size must be greater then zero")
	ErrIncorrectWorkersSize = errors.New("workers size must be greater then zero")
	ErrIncorrectBufferSize  = errors.New("buffer size must be greater then zero")
	ErrIncorrectState       = errors.New("executor is closed")
)

type UnitOfWork func() error

type Executor interface {
	Closed() bool
	Shutdown(awaitFinish bool) bool
	Execute(tasks []UnitOfWork) error
}

type errorsAwareExecutor struct {
	lock          sync.RWMutex
	workers       int
	maxErrors     int
	closed        bool
	tasksChannel  chan UnitOfWork
	errorsChannel chan error
	controlChan   chan struct{}
	wg            sync.WaitGroup
}

func (e *errorsAwareExecutor) Closed() bool {
	e.lock.RLock()
	result := e.closed
	e.lock.RUnlock()
	return result
}

func (e *errorsAwareExecutor) Shutdown(awaitFinish bool) bool {
	if e.Closed() {
		return false
	}

	e.lock.Lock()

	// State may have changed from another goroutine, so check again with a WriteLock
	if e.closed {
		e.lock.Unlock()
		return false
	}

	e.closed = true
	e.lock.Unlock()

	close(e.tasksChannel)
	close(e.controlChan)

	if awaitFinish {
		e.wg.Wait()
	}

	return true
}

func (e *errorsAwareExecutor) Execute(tasks []UnitOfWork) error {
	e.lock.Lock()
	defer e.lock.Unlock()

	if e.closed {
		return ErrIncorrectState
	}

	for i := range tasks {
		e.tasksChannel <- tasks[i]
	}

	return nil
}

func NewUnbufferedExecutor(maxErrors, maxWorkers int) (Executor, error) {
	return newExecutor(maxErrors, maxWorkers, maxWorkers)
}

func NewBufferedExecutor(maxErrors, maxWorkers, bufferSize int) (Executor, error) {
	return newExecutor(maxErrors, maxWorkers, bufferSize)
}

func newExecutor(maxErrors, maxWorkers, buffersSize int) (Executor, error) {
	if err := validateParams(maxErrors, maxWorkers, buffersSize); err != nil {
		return nil, err
	}

	tasksChan := make(chan UnitOfWork, buffersSize)
	errorsChan := make(chan error, maxErrors)
	controlChan := make(chan struct{})

	executor := &errorsAwareExecutor{
		workers:       maxWorkers,
		maxErrors:     maxErrors,
		tasksChannel:  tasksChan,
		errorsChannel: errorsChan,
		controlChan:   controlChan,
	}

	startErrorsController(executor)
	executor.wg.Add(maxWorkers)

	for i := 0; i < maxWorkers; i++ {
		runNewWorker(executor)
	}

	return executor, nil
}

func startErrorsController(executor *errorsAwareExecutor) {
	go func(errorsChan <-chan error, maxErrorsCount int, controlChan <-chan struct{}, executorStopCallback func()) {
		var currentErrors int

	loop:
		for {
			select {
			case err, isOpen := <-errorsChan:
				if isOpen {
					currentErrors++
					fmt.Printf("Received new error (total: %d/%d): %s.\n", currentErrors, maxErrorsCount, err)

					if currentErrors > maxErrorsCount {
						fmt.Printf("Errors limit exceeded: %d. Stopping executor...\n", maxErrorsCount)
						executorStopCallback()
						break loop
					}

					continue
				}

				// If channel is closed
				break loop

			case <-controlChan:
				break loop
			}
		}
	}(executor.errorsChannel, executor.maxErrors, executor.controlChan, func() { executor.Shutdown(false) })
}

func runNewWorker(executor *errorsAwareExecutor) {
	go func(tasksChan <-chan UnitOfWork, errorsChan chan<- error) {
		for {
			task, isAlive := <-tasksChan

			if !isAlive {
				break
			} else {
				err := task()

				if err != nil {
					errorsChan <- err
				}
			}
		}
	}(executor.tasksChannel, executor.errorsChannel)
}

func validateParams(maxErrors, maxWorkers, buffersSize int) error {
	if maxErrors <= 0 {
		return ErrIncorrectErrorsSize
	}

	if maxWorkers <= 0 {
		return ErrIncorrectWorkersSize
	}

	if buffersSize <= 0 {
		return ErrIncorrectBufferSize
	}

	return nil
}

// The main function to start executing tasks
func main() {
	// Example usage
	executor, err := NewBufferedExecutor(3, 5, 10)
	if err != nil {
		fmt.Println("Failed to create executor:", err)
		return
	}

	tasks := make([]UnitOfWork, 0)

	// Creating some tasks
	for i := 1; i <= 10; i++ {
		taskNum := i
		tasks = append(tasks, func() error {
			fmt.Printf("Executing task #%d\n", taskNum)
			if taskNum%4 == 0 {
				return fmt.Errorf("task #%d failed", taskNum)
			}
			return nil
		})
	}

	// Execute the tasks
	err = executor.Execute(tasks)
	if err != nil {
		fmt.Println("Failed to execute tasks:", err)
		return
	}

	// Shutdown after completion
	executor.Shutdown(true)
	fmt.Println("All tasks completed or stopped.")
}


// Executing task #1
// Executing task #3
// Executing task #4
// Executing task #5
// Executing task #6
// Executing task #7
// Executing task #8
// Executing task #9
// Executing task #10
// Executing task #2
// fatal error: all goroutines are asleep - deadlock!

// goroutine 1 [semacquire]:
// sync.runtime_Semacquire(0x9?)
// 	/home/mehanic/.gvm/gos/go1.23.0/src/runtime/sema.go:71 +0x25
// sync.(*WaitGroup).Wait(0xc00008e070?)
// 	/home/mehanic/.gvm/gos/go1.23.0/src/sync/waitgroup.go:118 +0x48
// main.(*errorsAwareExecutor).Shutdown(0xc00003a1e0, 0x1)
// 	/home/mehanic/structure/13.error/error-goroutine/executor.go:62 +0x7e
// main.main()
// 	/home/mehanic/structure/13.error/error-goroutine/executor.go:214 +0x237
// exit status 2
