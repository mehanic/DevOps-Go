Этот код показывает, как можно использовать **функции как аргументы** в Go, чтобы делать **гибкую фильтрацию** и **подсчёт суммы** элементов массива по определённому критерию. Давай разложим всё по полочкам 👇

---

## 🔹 Что делает весь код?

Суммирует только те числа из слайса (`[]int`), которые соответствуют какому-то условию (например, чётные или положительные).

---

## 🔸 Функция `isEven(n int) bool`

```go
func isEven(n int) bool {
	return n%2 == 0
}
```

Возвращает `true`, если число чётное (делится на 2 без остатка).

Примеры:
```go
isEven(4) // true
isEven(5) // false
```

---

## 🔸 Функция `isPositive(n int) bool`

```go
func isPositive(n int) bool {
	return n > 0
}
```

Возвращает `true`, если число положительное.

Примеры:
```go
isPositive(3)  // true
isPositive(-1) // false
```

---

## 🔸 Функция `sum(numbers []int, criteria func(int) bool) int`

```go
func sum(numbers []int, criteria func(int) bool) int {
	result := 0
	for _, val := range numbers {
		if criteria(val) {
			result += val
		}
	}
	return result
}
```

**Объяснение:**
- `numbers` — слайс чисел, например: `[1, 2, 3, 4]`
- `criteria` — функция, которая проверяет число по какому-то условию (например, `isEven` или `isPositive`)
- `result` — результат (сумма чисел, прошедших проверку)

🔁 Цикл проходит по всем числам:
- Если `criteria(val)` вернёт `true`, то `val` прибавляется к `result`.

---

## 🔸 Пример использования:

```go
nums := []int{1, 2, 3, 4, 5, -2, -3}

evenSum := sum(nums, isEven)      // суммирует только чётные: 2 + 4 + (-2) = 4
positiveSum := sum(nums, isPositive) // суммирует только положительные: 1+2+3+4+5 = 15

fmt.Println(evenSum)     // 4
fmt.Println(positiveSum) // 15
```

---
