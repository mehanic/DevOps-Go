package main

import "fmt"

func nextLargestNumberToTheRight(nums []int) []int {
    n := len(nums)
    res := make([]int, n)
    stack := []int{}

    for i := n - 1; i >= 0; i-- {
        // Пока стек не пуст и верхний элемент <= nums[i], извлекаем из стека
        for len(stack) > 0 && stack[len(stack)-1] <= nums[i] {
            stack = stack[:len(stack)-1]
        }
        // Если стек пуст, значит нет следующего большего элемента справа
        if len(stack) == 0 {
            res[i] = -1
        } else {
            res[i] = stack[len(stack)-1]
        }
        // Добавляем текущий элемент в стек
        stack = append(stack, nums[i])
    }

    return res
}

func main() {
    nums := []int{2, 1, 2, 4, 3}
    fmt.Println(nextLargestNumberToTheRight(nums)) // Output: [4 2 4 -1 -1]
}
