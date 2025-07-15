package main

import "fmt"

type Variable[A any] interface {
    Get() A
    Set(A) error
}

type Environment map[string]Variable[any]
type Environment_g[V Variable[any]] map[string]V

func main() {
    e := new(Environment)
    eg := e
    fmt.Println(e, eg)
}


////

type Foo struct {
    v any
}

func (f Foo) Get() any {
    return f.v
}

func (f Foo) Set(v any) error {
    f.v = v
    return nil
}

func main1() {
    e := Environment{}
    g := Environment_g[Foo]{}

    e["key"] = Foo{1}
    g["key"] = Foo{2}
}

////

