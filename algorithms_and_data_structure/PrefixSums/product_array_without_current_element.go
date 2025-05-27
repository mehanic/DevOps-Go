package main

func productArrayWithoutCurrentElement(nums []int) []int {
    n := len(nums)
    res := make([]int, n)

    // Инициализируем res единицами
    for i := range res {
        res[i] = 1
    }

    // Заполняем res накопленным произведением слева направо
    for i := 1; i < n; i++ {
        res[i] = res[i-1] * nums[i-1]
    }

    // Умножаем res на накопленное произведение справа налево
    rightProduct := 1
    for i := n - 1; i >= 0; i-- {
        res[i] *= rightProduct
        rightProduct *= nums[i]
    }

    return res
}
