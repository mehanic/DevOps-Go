### Объяснение алгоритма `NodeDepths`

Алгоритм `NodeDepths` рассчитывает **сумму глубин всех узлов** в бинарном дереве. Глубина узла – это количество шагов от корня до этого узла. Мы будем использовать рекурсивную функцию для обхода всех узлов дерева и подсчёта глубины каждого из них.

#### Структура бинарного дерева:
```go
type BinaryTree struct {
    Value int
    Left, Right *BinaryTree
}
```
- **Value**: Значение текущего узла.
- **Left**: Указатель на левое поддерево.
- **Right**: Указатель на правое поддерево.

### Описание функции `NodeDepths`

1. **Основная функция:**
    ```go
    func NodeDepths(root *BinaryTree) int {
        return nodeDepthsHelper(root, 0)
    }
    ```
    - Принимает корень дерева `root` и вызывает вспомогательную рекурсивную функцию `nodeDepthsHelper`, передавая глубину `0` для начального уровня (глубина корня = 0).

2. **Вспомогательная функция `nodeDepthsHelper`:**
    ```go
    func nodeDepthsHelper(root *BinaryTree, depth int) int {
        if root == nil {
            return 0
        }
        return depth + nodeDepthsHelper(root.Left, depth+1) + nodeDepthsHelper(root.Right, depth+1)
    }
    ```
    - Функция рекурсивно вычисляет сумму глубин для каждого узла:
        - Если текущий узел равен `nil`, то возвращаем 0, так как пустое поддерево не имеет глубины.
        - В противном случае, возвращаем текущую глубину узла (`depth`) + рекурсивно вызываем функцию для левого и правого поддерева, увеличивая глубину на 1.
        
### Пример

Рассмотрим бинарное дерево:

```
        1
       / \
      2   3
     / \
    4   5
```

- Узел **1** (корень) имеет глубину 0.
- Узел **2** имеет глубину 1 (один шаг от корня).
- Узел **3** имеет глубину 1 (один шаг от корня).
- Узел **4** имеет глубину 2 (два шага от корня).
- Узел **5** имеет глубину 2 (два шага от корня).

### Шаги выполнения:

1. Вызов `NodeDepths(root)`:
    - Корень дерева (значение `1`), глубина `0`.
    - Рекурсивно вызываем `nodeDepthsHelper` для левого и правого поддерева.
  
2. В `nodeDepthsHelper(root, 0)`:
    - Для узла **1**:
      - Глубина = 0.
      - Рекурсивно считаем для левого поддерева (узел **2**, глубина = 1).
      - Рекурсивно считаем для правого поддерева (узел **3**, глубина = 1).

3. В `nodeDepthsHelper(root, 1)` для узла **2**:
    - Глубина = 1.
    - Рекурсивно считаем для левого поддерева (узел **4**, глубина = 2).
    - Рекурсивно считаем для правого поддерева (узел **5**, глубина = 2).

4. В `nodeDepthsHelper(root, 2)` для узлов **4** и **5**:
    - Для каждого узла глубина = 2.
    - Левый и правый поддеревья пустые, глубина для них = 0.

### Итоговый расчет:
- Узел **1**: глубина = 0
- Узел **2**: глубина = 1
- Узел **3**: глубина = 1
- Узел **4**: глубина = 2
- Узел **5**: глубина = 2

**Сумма всех глубин = 0 + 1 + 1 + 2 + 2 = 6.**

### Пример кода:
```go
package main

import "fmt"

type BinaryTree struct {
    Value int
    Left, Right *BinaryTree
}

// Основная функция
func NodeDepths(root *BinaryTree) int {
    return nodeDepthsHelper(root, 0)
}

// Вспомогательная рекурсивная функция
func nodeDepthsHelper(root *BinaryTree, depth int) int {
    if root == nil {
        return 0
    }
    // Суммируем текущую глубину и рекурсивно обрабатываем левое и правое поддеревья
    return depth + nodeDepthsHelper(root.Left, depth+1) + nodeDepthsHelper(root.Right, depth+1)
}

func main() {
    // Строим пример дерева
    root := &BinaryTree{Value: 1}
    root.Left = &BinaryTree{Value: 2}
    root.Right = &BinaryTree{Value: 3}
    root.Left.Left = &BinaryTree{Value: 4}
    root.Left.Right = &BinaryTree{Value: 5}

    // Получаем сумму всех глубин
    result := NodeDepths(root)
    fmt.Println("Сумма глубин всех узлов:", result) // Output: 6
}
```

### Временная сложность:
- **Время:** O(n), где `n` — количество узлов в дереве. Мы посещаем каждый узел один раз.
- **Пространство:** O(h), где `h` — высота дерева. Это пространство для рекурсивного стека (глубина рекурсии). В худшем случае (когда дерево вырождается в список), это будет O(n), а в среднем случае (сбалансированное дерево) — O(log n).

### Итог:
Этот алгоритм позволяет эффективно подсчитать сумму глубин всех узлов в бинарном дереве, при этом используя рекурсивный обход дерева.