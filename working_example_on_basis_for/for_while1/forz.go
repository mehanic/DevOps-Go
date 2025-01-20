// for文：goには繰り返し処理はforしかない（whileはない）
package main

import (
	"fmt"
)

func main() {
	// for i := 0; i < 10; i++ {
	// 	// if i == 3 {
	// 	// 	break
	// 	// }
	// 	// if i == 3 {
	// 	// 	continue
	// 	// }
	// 	// fmt.Println(i)
	// }

	// while文的な書き方
	// i := 0
	// for i < 10 {
	// 	fmt.Println(i)
	// 	i++
	// }

	// while文的な書き方その２(あわや無限ループ)
	i := 0
	for {
		fmt.Println(i)
		i++
		if i == 3 {
			break
		}
	}

}
