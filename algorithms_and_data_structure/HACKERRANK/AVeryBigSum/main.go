package main
func aVeryBigSum(ar []int64) int64 {
    var sum int64 = 0           // 1. инициализация суммы
    for _, num := range ar {    // 2. цикл по массиву
        sum += num              // 3. прибавляем каждый элемент
    }
    return sum                  // 4. возвращаем результат
}
