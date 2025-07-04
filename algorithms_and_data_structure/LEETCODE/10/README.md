
Determine if a 9 x 9 Sudoku board is valid. Only the filled cells need to be validated according to the following rules:

	1. Each row must contain the digits 1-9 without repetition.
	
	2. Each column must contain the digits 1-9 without repetition.
	
	3. Each of the nine 3 x 3 sub-boxes of the grid must contain the digits 1-9 without repetition.

Note:

	* A Sudoku board (partially filled) could be valid but is not necessarily solvable.
	
	* Only the filled cells need to be validated according to the mentioned rules.



󰛨 Example 1:

[img](https://upload.wikimedia.org/wikipedia/commons/thumb/f/ff/Sudoku-by-L2G-20050714.svg/250px-Sudoku-by-L2G-20050714.svg.png)

	│ Input: board = 
	│ [["5","3",".",".","7",".",".",".","."]
	│ ,["6",".",".","1","9","5",".",".","."]
	│ ,[".","9","8",".",".",".",".","6","."]
	│ ,["8",".",".",".","6",".",".",".","3"]
	│ ,["4",".",".","8",".","3",".",".","1"]
	│ ,["7",".",".",".","2",".",".",".","6"]
	│ ,[".","6",".",".",".",".","2","8","."]
	│ ,[".",".",".","4","1","9",".",".","5"]
	│ ,[".",".",".",".","8",".",".","7","9"]]
	│ Output: true

󰛨 Example 2:

	│ Input: board = 
	│ [["8","3",".",".","7",".",".",".","."]
	│ ,["6",".",".","1","9","5",".",".","."]
	│ ,[".","9","8",".",".",".",".","6","."]
	│ ,["8",".",".",".","6",".",".",".","3"]
	│ ,["4",".",".","8",".","3",".",".","1"]
	│ ,["7",".",".",".","2",".",".",".","6"]
	│ ,[".","6",".",".",".",".","2","8","."]
	│ ,[".",".",".","4","1","9",".",".","5"]
	│ ,[".",".",".",".","8",".",".","7","9"]]
	│ Output: false
	│ Explanation: Same as Example 1, except with the 5 in the top left corner being modified to 8. Since there are two 8's in the top left 3x3 sub-box, it is invalid.



 Constraints:

	* board.length == 9
	
	* board[i].length == 9
	
	* board[i][j] is a digit 1-9 or '.'.
----------------------------------------------


Отлично! Давай пошагово, **с полным объяснением**, разберем этот код на Go (Golang), который **проверяет валидность Sudoku доски** размером 9x9.

---

## 📌 Цель

Проверить, корректно ли расставлены цифры в Судоку:

* В **каждой строке** не должно быть повторяющихся цифр.
* В **каждом столбце** не должно быть повторяющихся цифр.
* В **каждом из девяти 3x3 блоков** не должно быть повторяющихся цифр.

---

## 🔍 Входные данные

```go
board [][]byte
```

Это двумерный срез (`slice`) 9x9, где каждая ячейка:

* либо содержит цифру от `'1'` до `'9'` (в виде байта),
* либо `'.'` — пустая клетка (ее мы игнорируем).

---

## 📦 Объявление структур

```go
rows := make([]map[byte]bool, 9)
cols := make([]map[byte]bool, 9)
blocks := make([]map[byte]bool, 9)
```

Мы создаём:

* `rows[i]` — множество значений, встреченных в строке `i`
* `cols[j]` — множество значений, встреченных в столбце `j`
* `blocks[k]` — множество значений, встреченных в блоке `k`

> `map[byte]bool` — это способ отслеживать, добавлена ли уже конкретная цифра.

---

## 🧱 Инициализация множеств

```go
for i := 0; i < 9; i++ {
	rows[i] = make(map[byte]bool)
	cols[i] = make(map[byte]bool)
	blocks[i] = make(map[byte]bool)
}
```

Создаём по 9 словарей (массивы из словарей), чтобы в каждом отслеживать повторения.

---

## 🔁 Обход всей доски

```go
for r := 0; r < 9; r++ {
	for c := 0; c < 9; c++ {
```

Циклы по всем 81 ячейкам (по строкам и столбцам).

---

## ⛔ Пропуск пустых клеток

```go
val := board[r][c]
if val == '.' {
	continue
}
```

Если ячейка пустая (не содержит цифру) — пропускаем.

---

## 🔢 Нахождение номера блока

```go
blockId := getBlockNumber(r, c)
```

Функция:

```go
func getBlockNumber(row, col int) int {
	return (row/3)*3 + col/3
}
```

Позволяет определить, в какой из 9 **блоков 3x3** находится клетка `(r, c)`.

### Как это работает:

| row | col | blockId |
| --- | --- | ------- |
| 0–2 | 0–2 | 0       |
| 0–2 | 3–5 | 1       |
| 0–2 | 6–8 | 2       |
| 3–5 | 0–2 | 3       |
| 3–5 | 3–5 | 4       |
| 3–5 | 6–8 | 5       |
| 6–8 | 0–2 | 6       |
| 6–8 | 3–5 | 7       |
| 6–8 | 6–8 | 8       |

---

## ❌ Проверка на дубликаты

```go
if rows[r][val] || cols[c][val] || blocks[blockId][val] {
	return false
}
```

Проверяем:

* Было ли уже такое значение `val` в:

  * строке `r`
  * столбце `c`
  * блоке `blockId`

Если **да** — значит судоку невалиден.

---

## ✅ Запоминаем значение

Если нет дубликатов, добавляем значение во все 3 множества:

```go
rows[r][val] = true
cols[c][val] = true
blocks[blockId][val] = true
```

---

## ✅ Финальный результат

```go
return true
```

Если прошли всю доску без нарушений — судоку валиден.

---

## 🧪 Пример: клетка (4,5)

* Пусть в `board[4][5] = '3'`
* Строка = 4
* Столбец = 5
* Блок = `(4/3)*3 + (5/3) = 1*3 + 1 = 4` — центральный блок
* Если `'3'` уже была в `rows[4]`, `cols[5]` или `blocks[4]` — ошибка.
* Иначе — записываем `'3'` туда.

---

## 🧠 Заключение

Этот алгоритм:

* Проходит всю доску 1 раз → **O(81)** = **O(1)** (константа).
* Использует эффективную структуру `map` для отслеживания дубликатов.
* Прост и работает быстро.



Давай разберём **эту часть кода** пошагово, построчно и очень подробно:

```go
for r := 0; r < 9; r++ {
	for c := 0; c < 9; c++ {
```

### 🔁 Внешний и внутренний циклы

* `r` — индекс строки (от `0` до `8`)
* `c` — индекс столбца (от `0` до `8`)

Ты проходишь **все 81 клетку** Sudoku-доски (9 строк × 9 столбцов).

---

```go
val := board[r][c]
```

### 🧱 Получаем значение в ячейке

Ты читаешь **символ** из ячейки на позиции `(r, c)`.

* Пример: `val = '5'`, если в ячейке стоит пятёрка.
* Или `val = '.'`, если клетка пустая.

---

```go
if val == '.' {
	continue
}
```

### ⛔ Пропускаем пустые клетки

Если клетка **не содержит цифру**, то она **не участвует в проверке**.
Символ `'.'` означает, что клетка пока не заполнена.

—> `continue` пропускает оставшийся код в этом цикле и переходит к следующей ячейке.

---

```go
blockId := getBlockNumber(r, c)
```

### 🧮 Вычисляем ID блока

Вызов функции `getBlockNumber(r, c)` возвращает номер 3×3 блока, к которому принадлежит ячейка.

**Формула:**

```go
func getBlockNumber(row, col int) int {
	return (row / 3) * 3 + (col / 3)
}
```

Пример:

* `(4, 5)` → `(4/3)*3 + (5/3) = 1*3 + 1 = 4` → центральный блок (нумерация от 0 до 8)

---

```go
if rows[r][val] || cols[c][val] || blocks[blockId][val] {
	return false
}
```

### ❌ Проверка на дубликаты

Ты проверяешь, **была ли уже такая цифра (`val`)**:

* в строке `r` → `rows[r][val]`
* в столбце `c` → `cols[c][val]`
* в блоке `blockId` → `blocks[blockId][val]`

Если **хоть где-то уже была** эта цифра — это нарушение правил Sudoku → возвращаем `false`.

---

💡 **К чему это всё сводится?**
Ты проверяешь **уникальность** каждой цифры:

* в своей строке
* в своём столбце
* в своём блоке 3x3

Это **основные правила Sudoku**.

---

Если хочешь, я могу:

* визуально показать, как работает `blockId`
* добавить отладочные `print`-выводы
* переписать на Python или Java
* сделать проверку пользовательского ввода

Скажи, что интересно!
