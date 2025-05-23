package main

import (
	"fmt"
)

type user struct {
	name  string
	score int
}

func (u user) show() {
	fmt.Printf("name:%s, score:%d\n", u.name, u.score)
}

func (u *user) hit() {
	u.score++
}

func main() {
	u := user{name: "foo", score: 20}
	// fmt.Println()
	u.hit()
	u.show()
}

//name:foo, score:21
