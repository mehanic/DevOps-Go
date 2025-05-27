package main

import "fmt"

func josephus(n int, k int) int {
    if n == 1 {
        return 0
    }
    return (josephus(n-1, k) + k) % n
}

func main() {
    fmt.Println(josephus(7, 3)) // Пример вызова, вывод: 3
}
