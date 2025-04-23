### **Объяснение алгоритма `FindClosestValue` в BST (Binary Search Tree)**  

Этот алгоритм ищет **значение в бинарном дереве поиска (BST)**, которое **наиболее близко** к заданному `target`.

---

## **1. Разбор структуры BST**
```go
type BST struct {
    Value int
    Left  *BST
    Right *BST
}
```
- `Value` – значение текущего узла.
- `Left` – указатель на **левое поддерево**.
- `Right` – указатель на **правое поддерево**.

---

## **2. Основная функция `FindClosestValue`**
```go
func (tree *BST) FindClosestValue(target int) int {
    return tree.findClosestValue(target, math.MaxInt32)
}
```
- Вызывает **рекурсивную** функцию `findClosestValue`.
- Начальное **"ближайшее значение"** (`closest`) устанавливается в **максимальное возможное число** (`math.MaxInt32`).

---

## **3. Рекурсивная функция `findClosestValue`**
```go
func (tree *BST) findClosestValue(target, closest int) int {
    if absdiff(target, closest) > absdiff(target, tree.Value) {
        closest = tree.Value
    }

    if target < tree.Value && tree.Left != nil {
        return tree.Left.findClosestValue(target, closest)
    } else if target > tree.Value && tree.Right != nil {
        return tree.Right.findClosestValue(target, closest)
    }

    return closest
}
```

### **Логика работы**
1. **Сравниваем текущее значение с `closest`**
   ```go
   if absdiff(target, closest) > absdiff(target, tree.Value) {
       closest = tree.Value
   }
   ```
   - Если текущий узел ближе к `target`, обновляем `closest`.

2. **Выбираем направление для поиска**
   ```go
   if target < tree.Value && tree.Left != nil {
       return tree.Left.findClosestValue(target, closest)
   } else if target > tree.Value && tree.Right != nil {
       return tree.Right.findClosestValue(target, closest)
   }
   ```
   - Если `target` **меньше**, идем в **левое поддерево**.
   - Если `target` **больше**, идем в **правое поддерево**.

3. **Когда возвращаем `closest`?**
   - Когда **достигли листового узла** (нет левого/правого поддерева).

---

## **4. Функция `absdiff` (разница в значениях)**
```go
func absdiff(a, b int) int {
    out := math.Abs(float64(a) - float64(b))
    return int(out)
}
```
- Вычисляет **модуль разницы** между числами `a` и `b`.

---

## **5. Пример работы алгоритма**
### **Дерево BST**
```
        10
       /  \
      5   15
     / \   / \
    2   5 13 22
       /
      12
```
### **Код для примера**
```go
func main() {
    root := &BST{Value: 10}
    root.Left = &BST{Value: 5}
    root.Right = &BST{Value: 15}
    root.Left.Left = &BST{Value: 2}
    root.Left.Right = &BST{Value: 5}
    root.Right.Left = &BST{Value: 13}
    root.Right.Right = &BST{Value: 22}
    root.Right.Left.Left = &BST{Value: 12}

    target := 12
    fmt.Println("Closest value:", root.FindClosestValue(target))
}
```
### **Разбор работы для `target = 12`**
1. `10` ближе к `12`, чем `math.MaxInt32`.  
   → **closest = 10**
2. `12 > 10`, идем вправо (`15`).
3. `15` ближе к `12`, чем `10`.  
   → **closest = 15**
4. `12 < 15`, идем влево (`13`).
5. `13` ближе к `12`, чем `15`.  
   → **closest = 13**
6. `12 < 13`, идем влево (`12`).
7. `12` **точное совпадение**  
   → **Возвращаем `12`**.

### **Вывод**
```sh
Closest value: 12
```

---

## **6. Временная сложность**
| Случай  | Время  | Пространство |
|---------|--------|-------------|
| **Средний случай (сбалансированное дерево)** | `O(log(n))` | `O(log(n))` |
| **Худший случай (линейное дерево)** | `O(n)` | `O(n)` |

✅ **Алгоритм эффективно находит ближайшее значение к `target` в BST.**