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

// Точка начала или конца интервала
type Point struct {
	Time int
	Type rune // 'S' для Start, 'E' для End
}

// Нахождение максимального количества перекрывающихся интервалов
func largestOverlapOfIntervals(intervals []Interval) int {
	var points []Point

	// Преобразование интервалов в точки начала и конца
	for _, interval := range intervals {
		points = append(points, Point{Time: interval.Start, Type: 'S'})
		points = append(points, Point{Time: interval.End, Type: 'E'})
	}

	// Сортировка: сначала по времени, потом 'E' перед 'S' при равных значениях
	sort.Slice(points, func(i, j int) bool {
		if points[i].Time == points[j].Time {
			return points[i].Type == 'E'
		}
		return points[i].Time < points[j].Time
	})

	activeIntervals := 0
	maxOverlaps := 0

	for _, p := range points {
		if p.Type == 'S' {
			activeIntervals++
		} else {
			activeIntervals--
		}
		if activeIntervals > maxOverlaps {
			maxOverlaps = activeIntervals
		}
	}

	return maxOverlaps
}

// Пример использования
func main() {
	intervals := []Interval{
		{1, 5},
		{2, 6},
		{4, 8},
		{7, 9},
	}

	result := largestOverlapOfIntervals(intervals)
	fmt.Println("Максимальное количество перекрывающихся интервалов:", result)
}
