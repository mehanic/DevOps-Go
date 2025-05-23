package main

import (
	"context"
	"fmt"
)

func Greetings(ctx context.Context) {

	fmt.Printf("Welcome %v ,\nWe learn golang here!!!", ctx.Value("name"))

	/*

		We can return the name for future usage

		return ctx.Value("name")
	*/
}

func main() {

	ctx := context.Background() // It does not have to be "ctx", can be any name.

	ctx = context.WithValue(ctx, "name", "Berkay") // "myKey", "myValue"

	Greetings(ctx)
}
