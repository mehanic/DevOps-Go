package main

import (
	"fmt"
)

func combinationsOfSumK(nums []int, target int) [][]int {
	var res [][]int
	var combination []int
	dfs(combination, 0, nums, target, &res)
	return res
}

func dfs(combination []int, startIndex int, nums []int, target int, res *[][]int) {
	// Условие завершения: если target == 0, сохраняем комбинацию
	if target == 0 {
		// Создаём копию combination, чтобы избежать мутаций
		temp := make([]int, len(combination))
		copy(temp, combination)
		*res = append(*res, temp)
		return
	}

	// Если target стал отрицательным, прекращаем поиск
	if target < 0 {
		return
	}

	// Перебираем числа, начиная с текущего индекса
	for i := startIndex; i < len(nums); i++ {
		combination = append(combination, nums[i])                   // добавляем число в комбинацию
		dfs(combination, i, nums, target-nums[i], res)               // рекурсивный вызов с обновлённой целью
		combination = combination[:len(combination)-1]               // backtrack
	}
}

func main() {
	nums := []int{2, 3, 6, 7}
	target := 7
	result := combinationsOfSumK(nums, target)
	fmt.Println(result) // [[7] [2 2 3]] например, в зависимости от входа
}
