                                    
                  122. Best Time to Buy and Sell Stock II
                  Medium │ 14442  2762  │ 69.3% of 3.6M



You are given an integer array prices where prices[i] is the price of a given stock on the i^th day.

On each day, you may decide to buy and/or sell the stock. You can only hold at most one share of the stock at any time. However, you can buy it then immediately sell it on the same day.

Find and return the maximum profit you can achieve.



󰛨 Example 1:

	│ Input: prices = [7,1,5,3,6,4]
	│ Output: 7
	│ Explanation: Buy on day 2 (price = 1) and sell on day 3 (price = 5), profit = 5-1 = 4.
	│ Then buy on day 4 (price = 3) and sell on day 5 (price = 6), profit = 6-3 = 3.
	│ Total profit is 4 + 3 = 7.

󰛨 Example 2:

	│ Input: prices = [1,2,3,4,5]
	│ Output: 4
	│ Explanation: Buy on day 1 (price = 1) and sell on day 5 (price = 5), profit = 5-1 = 4.
	│ Total profit is 4.

󰛨 Example 3:

	│ Input: prices = [7,6,4,3,1]
	│ Output: 0
	│ Explanation: There is no way to make a positive profit, so we never buy the stock to achieve the maximum profit of 0.



 Constraints:

	* 1 <= prices.length <= 3 * 10^4
	
	* 0 <= prices[i] <= 10^4


Хорошо, давай разберёмся пошагово, почему максимальная прибыль вычисляется именно так и как это делается.

---

### ❓ **Что нужно сделать**

Найти **максимальную прибыль**, которую можно получить, если **в любой день можно купить и/или продать акцию**, но:

* можно держать только **одну акцию за раз**;
* можно **сразу же продать после покупки в тот же день**;
* можно совершать **много сделок** (неограниченно), **но не более одной покупки подряд без продажи**.

---

### ✅ **Основная идея (жадный подход)**

Чтобы максимизировать прибыль, **мы покупаем акцию каждый раз, когда цена пойдет вверх, и продаем на следующий день**.

То есть **если цена в следующий день больше, чем сегодня**, мы покупаем сегодня и продаём завтра.

---

### 📈 Пример: `[7,1,5,3,6,4]`

Дни и цены:

* День 0: 7
* День 1: 1
* День 2: 5  → (прибыль = 5 - 1 = 4)
* День 3: 3
* День 4: 6  → (прибыль = 6 - 3 = 3)
* День 5: 4

**Общая прибыль: 4 + 3 = 7**

Мы берём только такие участки, где цена растёт.

---

### 📉 Пример: `[7,6,4,3,1]`

Цены всё время падают — **нет ни одного дня, где можно получить прибыль**, так что **ответ = 0**.

---

### 🧠 Почему это работает?

Потому что каждый участок роста можно рассматривать как отдельную сделку:

* неважно, держим ли акцию несколько дней до самого пика;
* или покупаем и продаём каждый день, пока цена растёт — результат один и тот же.

Например:

Массив: `[1, 2, 3, 4]`

* Можно купить на 1, продать на 4 → прибыль = 3
* Или:

  * купить на 1, продать на 2 → прибыль = 1
  * купить на 2, продать на 3 → прибыль = 1
  * купить на 3, продать на 4 → прибыль = 1
* Общая прибыль = 1 + 1 + 1 = 3 — **то же самое**.




---

### 🧮 Сложность:

* **Время**: O(n) — проходим один раз по массиву.
* **Память**: O(1) — не используем дополнительную память.

Вот эквивалент алгоритма на **Go (Golang)**, который считает максимальную прибыль по такому же принципу:

```go
func maxProfit(prices []int) int {
    profit := 0
    for i := 1; i < len(prices); i++ {
        if prices[i] > prices[i-1] {
            profit += prices[i] - prices[i-1]
        }
    }
    return profit
}
```

---

### 🔍 Пояснение:

* `prices []int` — массив цен на каждый день.
* `for i := 1; i < len(prices); i++` — начинаем с первого индекса (второй день), чтобы сравнивать с предыдущим.
* `if prices[i] > prices[i-1]` — если цена растёт, считаем прибыль.
* `profit += prices[i] - prices[i-1]` — добавляем разницу к общей прибыли.

---

### 📦 Пример использования:

```go
package main

import (
	"fmt"
)

func maxProfit(prices []int) int {
	profit := 0
	for i := 1; i < len(prices); i++ {
		if prices[i] > prices[i-1] {
			profit += prices[i] - prices[i-1]
		}
	}
	return profit
}

func main() {
	prices := []int{7, 1, 5, 3, 6, 4}
	fmt.Println("Максимальная прибыль:", maxProfit(prices))
}
```

