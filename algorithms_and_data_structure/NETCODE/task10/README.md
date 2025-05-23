Вот оптимизированное решение для задачи **Two Sum** на Go.  

### **Алгоритм (O(n) - Hash Map)**
1. Используем **хеш-таблицу (map[int]int])** для хранения уже просмотренных чисел и их индексов.
2. Для каждого числа `nums[i]` вычисляем **нужное дополнение**:  
   \[
   complement = target - nums[i]
   \]
3. Если `complement` уже есть в **map**, значит мы нашли пару.
4. Возвращаем индексы в порядке `[меньший индекс, больший индекс]`.
5. Если не нашли — продолжаем.

---

### **Реализация (Go)**
```go
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
```

---

### **Разбор кода**
1. **Создаем map** `seen` для хранения уже просмотренных чисел и их индексов.
2. **Итерируемся по массиву** `nums`:
   - Вычисляем `complement = target - nums[i]`.
   - Проверяем, есть ли `complement` в `seen`.
   - Если да — возвращаем индексы `[j, i]`, где `j` — индекс найденного дополнения.
   - Если нет — записываем `nums[i]` в `seen` с индексом `i`.

---

### **Сложность**
- **Временная сложность**: **O(n)** (проходим массив один раз).
- **Доп. память**: **O(n)** (используем map для хранения чисел).

**🔥 Готовое эффективное решение!** 🚀