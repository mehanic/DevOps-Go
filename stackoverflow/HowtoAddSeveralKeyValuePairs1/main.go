package main

import (
    "context"
    "fmt"
)

func WithValues(ctx context.Context, vals map[any]any) context.Context {
    return multiValueContext{ctx, vals}
}

type multiValueContext struct {
    context.Context // Embed to inherit base methods
    vals            map[any]any
}

// Override the Value method to check our custom map first
func (ctx multiValueContext) Value(key any) any {
    if v, ok := ctx.vals[key]; ok {
        return v
    }
    return ctx.Context.Value(key)
}

func main() {
    // Base context
    ctx := context.Background()

    // Our custom key-value pairs
    values := map[any]any{
        "userID":   42,
        "username": "alice",
        "admin":    true,
    }

    // Create a new context with those values
    ctx = WithValues(ctx, values)

    // Retrieve values
    fmt.Println("userID:", ctx.Value("userID"))     // Output: 42
    fmt.Println("username:", ctx.Value("username")) // Output: alice
    fmt.Println("admin:", ctx.Value("admin"))       // Output: true

    // Show fallback behavior: looking up a key not in custom vals
    fmt.Println("nonexistent:", ctx.Value("missing")) // Output: <nil>
}
