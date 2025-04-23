package kata


// Write functions:

// 1: an(n) with parameter n: returns the first n terms of the series of a(n) (not tested)

// 2: gn(n) with parameter n: returns the first n terms of the series of g(n) (not tested)

// 3: countOnes(n) with parameter n: returns the number of 1 in the series gn(n) 
//     (don't forget to add a `1` at the head) # (tested)
    
// 4:  p(n) with parameter n: returns an array filled with the n first distinct primes in the same order they are found in the sequence gn (not tested)

// 5: maxPn(n) with parameter n: returns the biggest prime number of the above p(n) # (tested)

// 6: anOver(n) with parameter n: returns an array (n terms) of the a(i)/i for every i such g(i) != 1 (not tested but interesting result)

// 7: anOverAverage(n) with parameter n: returns as an *integer* the average of anOver(n) # (tested)
// Note:
// You can write directly functions 3:, 5: and 7:. There is no need to write functions 1:, 2:, 4: 6: except out of pure curiosity.




//import "fmt"

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