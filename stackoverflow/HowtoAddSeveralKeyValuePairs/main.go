package main

import (
    "context"
    "fmt"
)

func WithValues(ctx context.Context, kv ...interface{}) context.Context {
    if len(kv)%2 != 0 {
        panic("odd number of key-value pairs")
    }
    for i := 0; i < len(kv); i += 2 {
        ctx = context.WithValue(ctx, kv[i], kv[i+1])
    }
    return ctx
}

func main() {
    ctx := context.Background()

    // Use the custom WithValues function
    ctx = WithValues(ctx,
        "k1", "v1",
        "k2", "v2",
        "k3", "v3",
    )

    fmt.Println(ctx.Value("k1")) // Output: v1
    fmt.Println(ctx.Value("k2")) // Output: v2
    fmt.Println(ctx.Value("k3")) // Output: v3
}
