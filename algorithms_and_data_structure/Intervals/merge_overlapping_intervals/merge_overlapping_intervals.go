package main

import (
	"fmt"
	"sort"
)

// Определение структуры Interval
type Interval struct {
	Start int
	End   int
}

// Функция объединения пересекающихся интервалов
func mergeOverlappingIntervals(intervals []Interval) []Interval {
	if len(intervals) == 0 {
		return nil
	}

	// Сортируем интервалы по старту
	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i].Start < intervals[j].Start
	})

	merged := []Interval{intervals[0]}

	for _, B := range intervals[1:] {
		A := merged[len(merged)-1]
		// Если интервалы не пересекаются, добавляем B в результат
		if A.End < B.Start {
			merged = append(merged, B)
		} else {
			// Если пересекаются, объединяем их
			merged[len(merged)-1].End = max(A.End, B.End)
		}
	}

	return merged
}

// Вспомогательная функция для вычисления максимума
func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

// Пример использования
func main() {
	intervals := []Interval{
		{Start: 1, End: 4},
		{Start: 2, End: 5},
		{Start: 7, End: 9},
		{Start: 8, End: 10},
	}

	merged := mergeOverlappingIntervals(intervals)
	for _, interval := range merged {
		fmt.Printf("[%d, %d] ", interval.Start, interval.End)
	}
}
