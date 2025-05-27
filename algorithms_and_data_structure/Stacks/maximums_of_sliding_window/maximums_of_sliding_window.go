package main

import "fmt"

type pair struct {
	val int
	idx int
}

func maximumsOfSlidingWindow(nums []int, k int) []int {
	res := []int{}
	dq := []pair{} // Используем срез как deque
	left, right := 0, 0

	for right < len(nums) {
		// 1) Удаляем из конца dq все элементы, которые меньше или равны текущему
		for len(dq) > 0 && dq[len(dq)-1].val <= nums[right] {
			dq = dq[:len(dq)-1]
		}
		// 2) Добавляем текущий элемент
		dq = append(dq, pair{nums[right], right})

		// Если окно длиной k
		if right-left+1 == k {
			// 3) Удаляем элементы из начала dq, которые вышли за окно
			if len(dq) > 0 && dq[0].idx < left {
				dq = dq[1:]
			}
			// Максимум окна — первый элемент dq
			res = append(res, dq[0].val)
			left++
		}
		right++
	}
	return res
}

func main() {
	nums := []int{1, 3, -1, -3, 5, 3, 6, 7}
	k := 3
	result := maximumsOfSlidingWindow(nums, k)
	fmt.Println(result) // [3 3 5 5 6 7]
}
