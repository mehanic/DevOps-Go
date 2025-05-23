package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"runtime"
	"runtime/trace"
	"sync"
	"time"
)

func main() {
	f, err := os.Create("trace.out")
	if err != nil {
		log.Fatalln("Error:", err)
	}
	defer func() {
		if err := f.Close(); err != nil {
			log.Fatalln("Error:", err)
		}
	}()

	if err := trace.Start(f); err != nil {
		log.Fatalln("Error:", err)
	}
	defer trace.Stop()

	ctx, t := trace.NewTask(context.Background(), "main")
	defer t.End()
	fmt.Println("The number of logical CPU Cores:", runtime.NumCPU())

	// task(ctx, "Task1")
	// task(ctx, "Task2")
	// task(ctx, "Task3")

	var wg sync.WaitGroup
	wg.Add(3)
	go cTask(ctx, &wg, "Task1")
	go cTask(ctx, &wg, "Task2")
	go cTask(ctx, &wg, "Task3")
	wg.Wait()
	fmt.Println("main func finished")
}

func task(ctx context.Context, name string) {
	defer trace.StartRegion(ctx, name).End() // チェーンメソッドでは最後のメソッド(End())だけが遅延して呼ばれる
	time.Sleep(time.Second)
	fmt.Println(name)
}

func cTask(ctx context.Context, wg *sync.WaitGroup, name string) {
	defer trace.StartRegion(ctx, name).End()
	defer wg.Done()
	time.Sleep(time.Second)
	fmt.Println(name)
}

// ─λ go run main.go                                                                         0 (0.000s) < 16:06:09
// The number of logical CPU Cores: 12
// Task1
// Task2
// Task3
// main func finished


