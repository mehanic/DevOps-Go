// goroutine : 並行処理

package main

import (
	"fmt"
	"time"
)

func task1() {
	time.Sleep(time.Second * 2)
	fmt.Println("task1 finised!")
}

func task2() {
	fmt.Println("task2 finised!")

}

// go のprefixでを付すことで、並列計算が可能になる
func main() {
	// 同時に走らせたい処理がある場合には go を付けて goroutine にすることができる
	go task1()
	go task2()

	// goroutine が終わる前に main 関数自体が終わってしまうと何もおきないので、下記を付与
	time.Sleep(time.Second * 3)

}

// ╰─λ go run goroutinez.go                                                                  0 (0.001s) < 15:48:47
// task2 finised!
// task1 finised!
