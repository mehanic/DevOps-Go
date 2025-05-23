
package main

import "fmt"

func main() {

    // Это срез, если бы было написано nums := [3]int{2, 3, 4} - то это был бы массив
    nums := []int{2, 3, 4}
    sum := 0

    for _, num := range nums {
        sum += num
    }

    fmt.Println("sum:", sum)

    for index, num := range nums {
        fmt.Println("index:", index, "num:", num)
    }

    languages := map[string]string{"p": "php", "g": "golang"}
    for k, v := range languages {
        fmt.Printf("%s -> %s\n", k, v)
    }

    for k := range languages {
        fmt.Println("key:", k)
    }

    for i, c := range "golang" {
        fmt.Println(i, c)
    }

}


// sum: 9
// index: 0 num: 2
// index: 1 num: 3
// index: 2 num: 4
// p -> php
// g -> golang
// key: p
// key: g
// 0 103
// 1 111
// 2 108
// 3 97
// 4 110
// 5 103
