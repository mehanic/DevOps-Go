package main

import (
	"fmt"
)

func gasStations(gas []int, cost []int) int {
	totalGas, totalCost := 0, 0
	for i := 0; i < len(gas); i++ {
		totalGas += gas[i]
		totalCost += cost[i]
	}

	// Если общего количества бензина меньше, чем общих затрат — невозможно пройти круг
	if totalGas < totalCost {
		return -1
	}

	start, tank := 0, 0
	for i := 0; i < len(gas); i++ {
		tank += gas[i] - cost[i]
		if tank < 0 {
			// Следующая станция — новый старт
			start = i + 1
			tank = 0
		}
	}

	return start
}

func main() {
	gas := []int{1, 2, 3, 4, 5}
	cost := []int{3, 4, 5, 1, 2}
	fmt.Println("Start station index:", gasStations(gas, cost)) // Output: 3
}
