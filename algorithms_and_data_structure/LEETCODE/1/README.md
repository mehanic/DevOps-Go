https://leetcode.com/problems/remove-duplicates-from-sorted-array/
                                      
26. Remove Duplicates from Sorted Array
Easy │ 16613  19760  │ 60.0% of 10.2M │ 󰛨 Hints



Given an integer array nums sorted in non-decreasing order, remove the duplicates [in-place](https://en.wikipedia.org/wiki/In-place_algorithm) such that each unique element appears only once. The relative order of the elements should be kept the same. Then return the number of unique elements in nums.

Consider the number of unique elements of nums to be k, to get accepted, you need to do the following things:

* Change the array nums such that the first k elements of nums contain the unique elements in the order they were present in nums initially. The remaining elements of nums are not important as well as the size of nums.

* Return k.

Custom Judge:

The judge will test your solution with the following code:

│ int[] nums = [...]; // Input array
│ int[] expectedNums = [...]; // The expected answer with correct length
│ 
│ int k = removeDuplicates(nums); // Calls your implementation
│ 
│ assert k == expectedNums.length;
│ for (int i = 0; i < k; i++) {
│     assert nums[i] == expectedNums[i];
│ }

If all assertions pass, then your solution will be accepted.



󰛨 Example 1:

│ Input: nums = [1,1,2]
│ Output: 2, nums = [1,2,_]
│ Explanation: Your function should return k = 2, with the first two elements of nums being 1 and 2 respectively.
│ It does not matter what you leave beyond the returned k (hence they are underscores).

󰛨 Example 2:

│ Input: nums = [0,0,1,1,1,2,2,3,3,4]
│ Output: 5, nums = [0,1,2,3,4,_,_,_,_,_]
│ Explanation: Your function should return k = 5, with the first five elements of nums being 0, 1, 2, 3, and 4 respectively.
│ It does not matter what you leave beyond the returned k (hence they are underscores).



 Constraints:

* 1 <= nums.length <= 3 * 10^4

* -100 <= nums[i] <= 100

* nums is sorted in non-decreasing order.
------------------



Вот подробный разбор задачи и объяснение **на русском языке**:

---

## 📘 Условие задачи

Дан **отсортированный по неубыванию** массив целых чисел `nums`. Нужно **удалить все дубликаты** *на месте* (т.е. не используя дополнительную память), так чтобы каждый уникальный элемент встречался **только один раз**. Порядок элементов должен сохраниться.

Функция должна вернуть **количество уникальных элементов** `k`. Первые `k` ячеек массива `nums` должны содержать эти уникальные значения, а всё, что идёт после `k`, можно **игнорировать**.

---

## 🧩 Ограничения:

* Длина массива: `1 <= nums.length <= 30000`
* Значения элементов: от `-100` до `100`
* Массив уже **отсортирован**

---

## ✅ Что нужно сделать

* Удалить повторяющиеся значения
* Изменить массив **внутри себя**, не создавая новый аналогичный массив
* Вернуть **число уникальных элементов (k)**

---

## 🔍 Примеры:

```go
Ввод:  nums = [1, 1, 2]
Вывод: 2, массив становится [1, 2, _]
```

```go
Ввод:  nums = [0,0,1,1,1,2,2,3,3,4]
Вывод: 5, массив становится [0,1,2,3,4, _, _, _, _, _]
```

```go
Ввод:  nums = [1,2,3]
Вывод: 3, массив остаётся [1,2,3]
```

```go
Ввод:  nums = [2,2,2,2]
Вывод: 1, массив становится [2, _, _, _]
```

-------

## 💡 Подсказки:

### Используйте свойство отсортированности:

Раз массив отсортирован, то одинаковые элементы идут **подряд**. Достаточно сравнивать текущий элемент с предыдущим, чтобы понять — это дубликат или нет.

### Не создавайте новый массив:

Нельзя использовать `append`, `copy`, `remove`, `del` — всё должно быть **внутри текущего массива**, без дополнительных структур такого же размера.

### Используйте **два указателя (two-pointer technique)**:

* Один указатель (`slow`) — куда будем записывать следующий уникальный элемент.
* Второй (`fast`) — перебирает элементы массива.

Если `nums[fast] != nums[slow]`, то мы нашли новый уникальный элемент — увеличиваем `slow` и записываем в `nums[slow] = nums[fast]`.

---

Отлично! Давай подробно разберём, **как и почему работает** этот алгоритм удаления дубликатов из отсортированного массива `nums` **вместо, без использования дополнительной памяти**.

---

### 📌 Цель задачи:

Из **отсортированного массива** удалить все **повторяющиеся элементы**, чтобы в первых `k` ячейках остались только **уникальные значения**. При этом:

* 🔁 Порядок должен сохраняться.
* 🧠 Память O(1) — никаких новых массивов.
* 📤 Вернуть `k` — количество уникальных элементов.

---

## ✅ Пошаговое объяснение алгоритма:

```go
func removeDuplicates(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
```

### Шаг 1: Обработка пустого массива

Если массив пустой, то возвращаем 0 — дубликатов нет.

---

```go
	first := 0
```

### Шаг 2: `first` — это индекс для "нового" массива без дубликатов.

Он указывает на последнюю **уникальную** позицию в массиве. Начинаем с `0`, потому что первый элемент точно уникальный (массив отсортирован).

---

```go
	for elem := 1; elem < len(nums); elem++ {
```

### Шаг 3: Перебираем массив с позиции 1 до конца.

`elem` — это текущий индекс, который сравнивается с предыдущим уникальным значением (`nums[first]`).

---

```go
		if nums[elem] != nums[first] {
```

### Шаг 4: Если `nums[elem]` отличается от последнего уникального элемента,

значит мы **нашли новый уникальный элемент**.

---

```go
			first++
			nums[first] = nums[elem]
```

### Шаг 5: Двигаем указатель `first` вперёд и записываем туда новый уникальный элемент.

Тем самым мы **перезаписываем** повторяющиеся значения — это и есть "удаление" дубликатов **на месте**.

---

```go
	return first + 1
```

### Шаг 6: Возвращаем количество уникальных элементов.

Индексы начинаются с 0, поэтому `+1`.

---

### 🧠 Пример пошагово: `nums := []int{1, 1, 2, 2, 3}`

| Индекс `elem` | nums\[elem] | Сравнивается с nums\[first] | Действие          | first | Массив `nums`    |
| ------------- | ----------- | --------------------------- | ----------------- | ----- | ---------------- |
| 1             | 1           | 1                           | пропускаем        | 0     | \[1, 1, 2, 2, 3] |
| 2             | 2           | 1                           | found unique → 🔁 | 1     | \[1, 2, 2, 2, 3] |
| 3             | 2           | 2                           | пропускаем        | 1     | \[1, 2, 2, 2, 3] |
| 4             | 3           | 2                           | found unique → 🔁 | 2     | \[1, 2, 3, 2, 3] |

🔚 Итог: `first = 2`, `k = 3`, `nums[:3] = [1, 2, 3]`

---

### ✅ Почему это эффективно:

* ✔ Работает за O(n), один проход по массиву.
* ✔ Не использует дополнительную память (in-place).
* ✔ Использует свойства **отсортированного массива** — дубликаты идут подряд.

---

Хочешь, покажу такую же реализацию на Python или добавлю обработку edge-case'ов (например, когда все числа одинаковые)?
Вот пошаговая **табличка работы алгоритма `removeDuplicates`** на примере:

### 🔢 Исходный массив:

`nums := []int{1, 1, 2, 2, 3}`

| Шаг | `elem` индекс | `first` индекс | `nums[elem]` | `nums[first]` | Равны? | Действие                              | Массив на текущем шаге |
| --- | ------------- | -------------- | ------------ | ------------- | ------ | ------------------------------------- | ---------------------- |
| -   | -             | 0              | -            | 1             | -      | Начальное состояние                   | `[1, 1, 2, 2, 3]`      |
| 1   | 1             | 0              | 1            | 1             | ✅      | Пропустить (дубликат)                 | `[1, 1, 2, 2, 3]`      |
| 2   | 2             | 0              | 2            | 1             | ❌      | `first++`, `nums[first] = nums[elem]` | `[1, 2, 2, 2, 3]`      |
| 3   | 3             | 1              | 2            | 2             | ✅      | Пропустить                            | `[1, 2, 2, 2, 3]`      |
| 4   | 4             | 1              | 3            | 2             | ❌      | `first++`, `nums[first] = nums[elem]` | `[1, 2, 3, 2, 3]`      |

---

### ✅ Результат:

* `first = 2` → значит, уникальных элементов `k = first + 1 = 3`
* `nums[:k] = [1, 2, 3]` — всё, что нужно оставить
* Остальные значения за `k` можно игнорировать

