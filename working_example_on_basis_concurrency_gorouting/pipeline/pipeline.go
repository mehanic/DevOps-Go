package main

import (
	"context"
	"encoding/base64"
	"fmt"
	// "github.com/PacktPublishing/Go-Programming-Cookbook-Second-Edition/chapter10/pipeline"
)

func main() {
	ctx := context.Background()
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()

	in, out := NewPipeline(ctx, 10, 2)

	go func() {
		for i := 0; i < 20; i++ {
			in <- fmt.Sprint("Message", i)
		}
	}()

	for i := 0; i < 20; i++ {
		<-out
	}
}

//---

// Encode takes plain text as int
// and returns "string => <base64 string encoding>
// as out
func (w *Worker) Encode(ctx context.Context) {
	for {
		select {
		case <-ctx.Done():
			return
		case val := <-w.in:
			w.out <- fmt.Sprintf("%s => %s", val, base64.StdEncoding.EncodeToString([]byte(val)))
		}
	}
}

// NewPipeline initializes the workers and
// connects them, it returns the input of the pipeline
// and the final output
func NewPipeline(ctx context.Context, numEncoders, numPrinters int) (chan string, chan string) {
	inEncode := make(chan string, numEncoders)
	inPrint := make(chan string, numPrinters)
	outPrint := make(chan string, numPrinters)
	for i := 0; i < numEncoders; i++ {
		w := Worker{
			in:  inEncode,
			out: inPrint,
		}
		go w.Work(ctx, Encode)
	}

	for i := 0; i < numPrinters; i++ {
		w := Worker{
			in:  inPrint,
			out: outPrint,
		}
		go w.Work(ctx, Print)
	}
	return inEncode, outPrint
}

// Print prints w.in and repalys it
// on w.out
func (w *Worker) Print(ctx context.Context) {
	for {
		select {
		case <-ctx.Done():
			return
		case val := <-w.in:
			fmt.Println(val)
			w.out <- val
		}
	}
}

// Worker have one role
// that is determined when
// Work is called
type Worker struct {
	in  chan string
	out chan string
}

// Job is a job a worker can do
type Job string

const (
	// Print echo's all input to
	// stdout
	Print Job = "print"
	// Encode base64 encodes input
	Encode Job = "encode"
)

// Work is how to dispatch a worker, they are assigned
// a job here
func (w *Worker) Work(ctx context.Context, j Job) {
	switch j {
	case Print:
		w.Print(ctx)
	case Encode:
		w.Encode(ctx)
	default:
		return
	}
}

// Message0 => TWVzc2FnZTA=
// Message1 => TWVzc2FnZTE=
// Message2 => TWVzc2FnZTI=
// Message3 => TWVzc2FnZTM=
// Message4 => TWVzc2FnZTQ=
// Message5 => TWVzc2FnZTU=
// Message6 => TWVzc2FnZTY=
// Message7 => TWVzc2FnZTc=
// Message9 => TWVzc2FnZTk=
// Message10 => TWVzc2FnZTEw
// Message11 => TWVzc2FnZTEx
// Message12 => TWVzc2FnZTEy
// Message13 => TWVzc2FnZTEz
// Message14 => TWVzc2FnZTE0
// Message15 => TWVzc2FnZTE1
// Message16 => TWVzc2FnZTE2
// Message8 => TWVzc2FnZTg=
// Message18 => TWVzc2FnZTE4
// Message19 => TWVzc2FnZTE5
// Message17 => TWVzc2FnZTE3
