package main

import "fmt"

func max() (int, int) {
	return 10, 20
}

func main() {
	a := "used"
	//b := "not used" // 沒有用到的變數會報錯
	fmt.Println(a)

	c, _ := max() // blank identifier (_) 可以用來承接不會用到的回傳值, 避免定義用不到的變數造成錯誤
	fmt.Println(c)
}

// used
// 10
