package main

import "fmt"

func twoSum(nums []int, target int) []int {
    seen := make(map[int]int) // Храним числа и их индексы

    for i, num := range nums {
        complement := target - num

        if j, found := seen[complement]; found {
            return []int{j, i} // Возвращаем индексы в порядке [меньший, больший]
        }

        seen[num] = i // Запоминаем число и его индекс
    }
    
    return nil // По условию задачи всегда есть решение, но добавим этот случай
}

func main() {
    // Примеры
    fmt.Println(twoSum([]int{3, 4, 5, 6}, 7))  // [0,1]
    fmt.Println(twoSum([]int{4, 5, 6}, 10))    // [0,2]
    fmt.Println(twoSum([]int{5, 5}, 10))       // [0,1]
}
