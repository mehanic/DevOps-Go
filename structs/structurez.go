// 構造体:複数の値を意味のあるまとまりとして新しい型を定義することができる
package main

import (
	"fmt"
)

type user struct {
	name  string
	score int
}

func main() {
	// // 書き方１
	// u := new(user)
	// u.name = "baa"
	// u.score = 20
	// fmt.Println(u)

	// 書き方２：変数宣言とともに値を定義する
	// u:=user{"baa",20}
	u := user{name: "baa", score: 20}
	fmt.Println(u)
}
