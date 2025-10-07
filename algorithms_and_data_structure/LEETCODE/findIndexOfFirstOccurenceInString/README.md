Given two strings needle and haystack, return the index of the first occurrence of needle in haystack, or -1 if needle is not part of haystack.

 

Example 1:

Input: haystack = "sadbutsad", needle = "sad"
Output: 0
Explanation: "sad" occurs at index 0 and 6.
The first occurrence is at index 0, so we return 0.
Example 2:

Input: haystack = "leetcode", needle = "leeto"
Output: -1
Explanation: "leeto" did not occur in "leetcode", so we return -1.
 

Constraints:

1 <= haystack.length, needle.length <= 104
haystack and needle consist of only lowercase English characters.




Разберём построчно, **почему и как работает** эта функция `strStr` на Go:

```go
func strStr(haystack string, needle string) int {
```

🔹 Это объявление функции. Она принимает две строки:

* `haystack` — строка, в которой мы ищем подстроку;
* `needle` — подстрока, которую нужно найти.

---

```go
for i := 0; i <= len(haystack)-len(needle); i++ {
```

🔹 Цикл `for` проходит по `haystack` от начала до той позиции, где ещё можно вместить `needle`.

* `len(haystack) - len(needle)` — это последняя возможная позиция, где подстрока `needle` ещё может полностью уместиться.
* Например: если `haystack = "sadbutsad"`, длина = 9, и `needle = "sad"`, длина = 3, то последний возможный индекс начала — `9 - 3 = 6`.

---

```go
if haystack[i:i+len(needle)] == needle {
```

🔹 Здесь мы берём подстроку `haystack` от позиции `i` до `i+len(needle)` и сравниваем её с `needle`.

* Например, если `i = 0`, то `haystack[0:3] == "sad"`, что может совпасть с `needle == "sad"`.

---

```go
return i
```

🔹 Если подстрока совпала с `needle`, возвращаем индекс `i` — позицию первого вхождения.

---

```go
return -1
```

🔹 Если цикл прошёл полностью, и ни одно совпадение не было найдено — возвращаем `-1`, как указано в условии задачи.

---

### 🧠 Почему работает правильно:

* Проходим все возможные позиции в `haystack`, где `needle` может появиться.
* Сравниваем кусок строки напрямую с `needle`.
* Как только найдена первая подстрока, идентичная `needle`, возвращаем её индекс — задача решена.

Хочешь, я покажу, как добавить [обработку случая, если `needle` пустая строка](f) или объясню, как работает [поиск через KMP алгоритм](f)?
