package main

import (
    "context"
    "fmt"
)

type Values struct {
    m map[string]string
}

func (v Values) Get(key string) string {
    return v.m[key]
}

func main() {
    v := Values{map[string]string{
        "1": "one",
        "2": "two",
    }}

    c := context.Background()
    c2 := context.WithValue(c, "myvalues", v)

    val := c2.Value("myvalues").(Values)
    fmt.Println(val.Get("2")) // Output: two
}
