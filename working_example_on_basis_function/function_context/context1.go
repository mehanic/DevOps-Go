package main

import (
	"context"
	"fmt"
	"math/rand"
	"time"
)


func main() {
	Exec()
}

// Exec sets two random timers and prints
// a different context value for whichever
// fires first
func Exec() {
	// a base context
	ctx := context.Background()
	ctx = Setup(ctx)

	rand.Seed(time.Now().UnixNano())

	timeoutCtx, cancel := context.WithTimeout(ctx, (time.Duration(rand.Intn(2)) * time.Millisecond))
	defer cancel()

	deadlineCtx, cancel := context.WithDeadline(ctx, time.Now().Add(time.Duration(rand.Intn(2))*time.Millisecond))
	defer cancel()

	for {
		select {
		case <-timeoutCtx.Done():
			fmt.Println(GetValue(ctx, timeoutKey))
			return
		case <-deadlineCtx.Done():
			fmt.Println(GetValue(ctx, deadlineKey))
			return
		}
	}
}

//---

type key string

const (
	timeoutKey  key = "TimeoutKey"
	deadlineKey key = "DeadlineKey"
)

// Setup sets some values
func Setup(ctx context.Context) context.Context {

	ctx = context.WithValue(ctx, timeoutKey, "timeout exceeded")
	ctx = context.WithValue(ctx, deadlineKey, "deadline exceeded")

	return ctx
}

// GetValue grabs a value given a key and
// returns a string representation of the
// value
func GetValue(ctx context.Context, k key) string {

	if val, ok := ctx.Value(k).(string); ok {
		return val
	}
	return ""

}
