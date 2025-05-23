Для решения задачи на Go можно построить функции, основываясь на описанном алгоритме. Вот возможные реализации для функций `CountOnes`, `MaxPn`, и `AnOverAverage`.

### 1. **CountOnes**
Эта функция должна подсчитывать количество единиц в последовательности разностей `g(n)`, не забывая о добавлении единицы в начало.

```go
package kata

import "fmt"

// Функция для вычисления наибольшего общего делителя (gcd)
func gcd(a, b int64) int64 {
    for b != 0 {
        a, b = b, a%b
    }
    return a
}

// Функция CountOnes подсчитывает количество единиц в последовательности g(n)
func CountOnes(n int64) int {
    a := []int64{7} // Начальный элемент последовательности a(1)
    g := []int64{1} // Начальный элемент для g(1)

    // Генерация последовательностей a(n) и g(n)
    for i := int64(2); i <= n; i++ {
        aNext := a[i-2] + gcd(i, a[i-2])
        a = append(a, aNext)

        gNext := aNext - a[i-2]
        g = append(g, gNext)
    }

    // Подсчет количества единиц в g(n)
    count := 0
    for _, value := range g {
        if value == 1 {
            count++
        }
    }

    return count
}
```

### 2. **MaxPn**
Эта функция должна находить максимальное простое число из последовательности простых чисел, полученных из последовательности `g(n)`.

```go
package kata

// Проверка числа на простоту
func isPrime(n int64) bool {
    if n <= 1 {
        return false
    }
    for i := int64(2); i*i <= n; i++ {
        if n%i == 0 {
            return false
        }
    }
    return true
}

// Функция MaxPn находит максимальное простое число в последовательности p(n)
func MaxPn(n int64) int64 {
    a := []int64{7} // Начальный элемент последовательности a(1)
    g := []int64{1} // Начальный элемент для g(1)
    primes := []int64{}

    // Генерация последовательностей a(n) и g(n)
    for i := int64(2); i <= n; i++ {
        aNext := a[i-2] + gcd(i, a[i-2])
        a = append(a, aNext)

        gNext := aNext - a[i-2]
        g = append(g, gNext)
    }

    // Извлечение простых чисел из g(n)
    for _, value := range g {
        if value != 1 && isPrime(value) {
            primes = append(primes, value)
        }
    }

    // Поиск максимального простого числа в последовательности
    maxPrime := int64(0)
    for _, prime := range primes {
        if prime > maxPrime {
            maxPrime = prime
        }
    }

    return maxPrime
}
```

### 3. **AnOverAverage**
Эта функция должна вычислять среднее значение для последовательности \(a(i) / i\), где \(g(i) \neq 1\).

```go
package kata

// Функция AnOverAverage вычисляет среднее значение a(i) / i для всех i, где g(i) != 1
func AnOverAverage(n int64) int {
    a := []int64{7} // Начальный элемент последовательности a(1)
    g := []int64{1} // Начальный элемент для g(1)
    anOver := []float64{}

    // Генерация последовательностей a(n) и g(n)
    for i := int64(2); i <= n; i++ {
        aNext := a[i-2] + gcd(i, a[i-2])
        a = append(a, aNext)

        gNext := aNext - a[i-2]
        g = append(g, gNext)
    }

    // Заполнение последовательности a(i)/i для всех i, где g(i) != 1
    for i := int64(1); i <= n; i++ {
        if g[i] != 1 {
            anOver = append(anOver, float64(a[i])/float64(i))
        }
    }

    // Вычисление среднего значения
    sum := 0.0
    for _, value := range anOver {
        sum += value
    }

    average := sum / float64(len(anOver))

    // Возвращаем среднее значение как целое число
    return int(average)
}
```

### Объяснение:
1. **CountOnes**:
   - Мы генерируем последовательности \(a(n)\) и \(g(n)\), начиная с заданных начальных значений. Затем мы подсчитываем количество единиц в последовательности \(g(n)\).
   
2. **MaxPn**:
   - Мы генерируем последовательности \(a(n)\) и \(g(n)\), затем из последовательности \(g(n)\) извлекаем простые числа и возвращаем максимальное из них.

3. **AnOverAverage**:
   - Мы генерируем последовательности \(a(n)\) и \(g(n)\), затем вычисляем \(a(i) / i\) для всех \(i\), где \(g(i) \neq 1\), и возвращаем среднее значение этого отношения.

### Примечания:
- Мы используем функцию `gcd` для вычисления наибольшего общего делителя, которая необходима для вычисления последовательности \(a(n)\).
- Для проверки простоты числа используется стандартный метод, где проверяется делимость числа на все числа до его квадратного корня.
- В функции `AnOverAverage` мы вычисляем среднее значение как тип `float64` и затем преобразуем его в целое число.