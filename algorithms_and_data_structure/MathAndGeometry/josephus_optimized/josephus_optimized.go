package main

import "fmt"

func josephusOptimized(n int, k int) int {
    res := 0
    for i := 2; i <= n; i++ {
        res = (res + k) % i
    }
    return res
}

func main() {
    fmt.Println(josephusOptimized(7, 3)) // Вывод: 3 (пример)
}
