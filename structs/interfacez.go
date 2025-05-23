// インターフェース：メソッドの一覧を定義したデータ型
// それらのメソッドをある構造体が実装していれば、その構造体をインターフェース型として扱ってよいという特徴がある

package main

import (
	"fmt"
)

// インターフェースの定義　type インターフェース名　interface
type greeter interface {
	greet()
}

type japanese struct{}
type american struct{}

func (j japanese) greet() {
	fmt.Println("konnnichiwa!")
}

func (a american) greet() {
	fmt.Println("Hello!")
}

func main() {
	greeters := []greeter{japanese{}, american{}}
	for _, greeter := range greeters {
		greeter.greet()
	}

}
