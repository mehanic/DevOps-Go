package main

import "fmt"

// Определение структуры Interval
type Interval struct {
	Start int
	End   int
}

// Функция нахождения всех пересечений между двумя списками интервалов
func identifyAllIntervalOverlaps(intervals1 []Interval, intervals2 []Interval) []Interval {
	var overlaps []Interval
	i, j := 0, 0

	for i < len(intervals1) && j < len(intervals2) {
		var A, B Interval
		if intervals1[i].Start <= intervals2[j].Start {
			A, B = intervals1[i], intervals2[j]
		} else {
			A, B = intervals2[j], intervals1[i]
		}

		// Проверка на пересечение
		if A.End >= B.Start {
			overlap := Interval{
				Start: B.Start,
				End:   min(A.End, B.End),
			}
			overlaps = append(overlaps, overlap)
		}

		// Продвижение указателя того интервала, который заканчивается раньше
		if intervals1[i].End < intervals2[j].End {
			i++
		} else {
			j++
		}
	}

	return overlaps
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

// Пример использования
func main() {
	intervals1 := []Interval{{1, 5}, {10, 14}, {16, 18}}
	intervals2 := []Interval{{2, 6}, {8, 10}, {11, 20}}

	result := identifyAllIntervalOverlaps(intervals1, intervals2)

	for _, interval := range result {
		fmt.Printf("[%d, %d] ", interval.Start, interval.End)
	}
}
