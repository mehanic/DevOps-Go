### **Объяснение алгоритма на Go для начинающих**

Этот код на Go проверяет, является ли строка **палиндромом**.  
**Палиндром** – это строка, которая читается одинаково **справа налево и слева направо** (без учета регистра и пробелов).  
Например:
- ✅ `"Was it a car or a cat I saw?"` → `"wasitacaroracatisaw"` → **палиндром**
- ❌ `"tab a cat"` → `"tabacat"` → **не палиндром**

---

## **Разбор кода по шагам**

### **1. Определение функции `isPalindrome`**
```go
func isPalindrome(s string) bool {
```
- **Принимает строку `s`**.
- **Возвращает `true`, если строка является палиндромом**, иначе `false`.

---

### **2. Фильтрация строки**
```go
var filtered []rune
```
- Создаем **пустой список (`filtered`)** для хранения **только букв и цифр** в нижнем регистре.

```go
for _, ch := range s {
	if unicode.IsLetter(ch) || unicode.IsDigit(ch) {
		filtered = append(filtered, unicode.ToLower(ch))
	}
}
```
- **Цикл перебирает каждый символ `ch` в строке `s`**.
- **Проверка**:  
  - Если `ch` **буква** (`unicode.IsLetter(ch)`) или **цифра** (`unicode.IsDigit(ch)`), то:
    - Преобразуем в **нижний регистр** (`unicode.ToLower(ch)`)
    - Добавляем в `filtered`

✅ **Пример работы фильтрации:**  
Входная строка: `"Was it a car or a cat I saw?"`  
После фильтрации: `"wasitacaroracatisaw"`

---

### **3. Двухуказательный метод (Two-Pointer Technique)**
```go
left, right := 0, len(filtered)-1
```
- `left` указывает на **начало** строки.
- `right` указывает на **конец** строки.

```go
for left < right {
	if filtered[left] != filtered[right] {
		return false
	}
	left++
	right--
}
```
- Пока `left < right`:
  - Если символы **не совпадают**, значит это **не палиндром** → `return false`
  - Если символы **совпадают**, **двигаем указатели**:
    - `left++` (вправо)
    - `right--` (влево)
- Если **все символы совпадают**, то это палиндром.

✅ **Пример работы алгоритма:**  
Проверяем `"wasitacaroracatisaw"`:
```
w == w   ✅
a == a   ✅
s == s   ✅
i == i   ✅
t == t   ✅
a == a   ✅
c == c   ✅
a == a   ✅
r == r   ✅
(середина) ✅ → палиндром
```

---

### **4. Вывод результата**
```go
return true
```
- Если строка прошла проверку, значит **это палиндром**.

---

### **5. Тестирование алгоритма**
```go
func main() {
	examples := []string{
		"Was it a car or a cat I saw?", // true
		"tab a cat",                   // false
		"A man, a plan, a canal: Panama!", // true
		"RaceCar",                     // true
		"hello",                        // false
	}

	for _, example := range examples {
		fmt.Printf("isPalindrome(%q) = %v\n", example, isPalindrome(example))
	}
}
```
- **Массив тестовых строк** проверяется циклом `for`.
- **Выводится результат проверки**.

✅ **Вывод программы:**
```
isPalindrome("Was it a car or a cat I saw?") = true
isPalindrome("tab a cat") = false
isPalindrome("A man, a plan, a canal: Panama!") = true
isPalindrome("RaceCar") = true
isPalindrome("hello") = false
```

---

## **Итог**
- Фильтруем строку: **удаляем все, кроме букв и цифр**.
- Используем **двухуказательный метод** для проверки, читается ли строка одинаково **слева направо и справа налево**.
- Временная сложность: **O(n)** (линейная, где `n` — длина строки).
- Простая и **эффективная проверка палиндрома** на Go! 🚀