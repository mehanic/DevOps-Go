package main

import (
	"fmt"
	"sync"
)

var wg sync.WaitGroup

type Income struct {
	Source string
	Amount int
}

func main() {
	// var for bank balance
	var bankBalance int
	var balance sync.Mutex

	// print out starting values
	fmt.Printf("Initial account balance $%d.00", bankBalance)
	fmt.Println()

	// define weekly revenue
	incomes := []Income{
		{
			Source: "main job",
			Amount: 500,
		},
		{
			Source: "gifts",
			Amount: 10,
		},
		{
			Source: "part time job",
			Amount: 50,
		},
		{
			Source: "investments",
			Amount: 100,
		},
	}

	wg.Add(len(incomes))

	// loop through 52 weeks and print out how much is made; keep a running total
	for i, income := range incomes {
		go func(i int, income Income) {
			defer wg.Done()

			for week := 1; week <= 52; week++ {
				balance.Lock()
				temp := bankBalance
				temp += income.Amount
				bankBalance = temp
				balance.Unlock()
				fmt.Printf("on week %d, you earned $%d.00 from %s\n", week, income.Amount, income.Source)
			}
		}(i, income)
	}

	wg.Wait()
	// print out final balance
	fmt.Printf("final bank balance $%d.00", bankBalance)
	fmt.Println()
}

// ╰─λ go run main.go                                                                        0 (0.001s) < 12:55:12
// Initial account balance $0.00
// on week 1, you earned $100.00 from investments
// on week 2, you earned $100.00 from investments
// on week 3, you earned $100.00 from investments
// on week 4, you earned $100.00 from investments
// on week 5, you earned $100.00 from investments
// on week 6, you earned $100.00 from investments
// on week 7, you earned $100.00 from investments
// on week 8, you earned $100.00 from investments
// on week 9, you earned $100.00 from investments
// on week 10, you earned $100.00 from investments
// on week 11, you earned $100.00 from investments
// on week 12, you earned $100.00 from investments
// on week 13, you earned $100.00 from investments
// on week 14, you earned $100.00 from investments
// on week 15, you earned $100.00 from investments
// on week 16, you earned $100.00 from investments
// on week 17, you earned $100.00 from investments
// on week 18, you earned $100.00 from investments
// on week 19, you earned $100.00 from investments
// on week 20, you earned $100.00 from investments
// on week 21, you earned $100.00 from investments
// on week 22, you earned $100.00 from investments
// on week 23, you earned $100.00 from investments
// on week 24, you earned $100.00 from investments
// on week 25, you earned $100.00 from investments
// on week 26, you earned $100.00 from investments
// on week 27, you earned $100.00 from investments
// on week 28, you earned $100.00 from investments
// on week 29, you earned $100.00 from investments
// on week 30, you earned $100.00 from investments
// on week 31, you earned $100.00 from investments
// on week 32, you earned $100.00 from investments
// on week 33, you earned $100.00 from investments
// on week 34, you earned $100.00 from investments
// on week 35, you earned $100.00 from investments
// on week 36, you earned $100.00 from investments
// on week 37, you earned $100.00 from investments
// on week 38, you earned $100.00 from investments
// on week 39, you earned $100.00 from investments
// on week 1, you earned $50.00 from part time job
// on week 2, you earned $50.00 from part time job
// on week 3, you earned $50.00 from part time job
// on week 4, you earned $50.00 from part time job
// on week 5, you earned $50.00 from part time job
// on week 6, you earned $50.00 from part time job
// on week 7, you earned $50.00 from part time job
// on week 8, you earned $50.00 from part time job
// on week 9, you earned $50.00 from part time job
// on week 10, you earned $50.00 from part time job
// on week 11, you earned $50.00 from part time job
// on week 12, you earned $50.00 from part time job
// on week 13, you earned $50.00 from part time job
// on week 14, you earned $50.00 from part time job
// on week 15, you earned $50.00 from part time job
// on week 16, you earned $50.00 from part time job
// on week 17, you earned $50.00 from part time job
// on week 18, you earned $50.00 from part time job
// on week 19, you earned $50.00 from part time job
// on week 20, you earned $50.00 from part time job
// on week 21, you earned $50.00 from part time job
// on week 22, you earned $50.00 from part time job
// on week 23, you earned $50.00 from part time job
// on week 24, you earned $50.00 from part time job
// on week 25, you earned $50.00 from part time job
// on week 26, you earned $50.00 from part time job
// on week 27, you earned $50.00 from part time job
// on week 28, you earned $50.00 from part time job
// on week 29, you earned $50.00 from part time job
// on week 30, you earned $50.00 from part time job
// on week 31, you earned $50.00 from part time job
// on week 32, you earned $50.00 from part time job
// on week 40, you earned $100.00 from investments
// on week 41, you earned $100.00 from investments
// on week 42, you earned $100.00 from investments
// on week 1, you earned $500.00 from main job
// on week 2, you earned $500.00 from main job
// on week 3, you earned $500.00 from main job
// on week 4, you earned $500.00 from main job
// on week 5, you earned $500.00 from main job
// on week 6, you earned $500.00 from main job
// on week 7, you earned $500.00 from main job
// on week 8, you earned $500.00 from main job
// on week 9, you earned $500.00 from main job
// on week 10, you earned $500.00 from main job
// on week 11, you earned $500.00 from main job
// on week 12, you earned $500.00 from main job
// on week 13, you earned $500.00 from main job
// on week 33, you earned $50.00 from part time job
// on week 34, you earned $50.00 from part time job
// on week 35, you earned $50.00 from part time job
// on week 36, you earned $50.00 from part time job
// on week 37, you earned $50.00 from part time job
// on week 38, you earned $50.00 from part time job
// on week 39, you earned $50.00 from part time job
// on week 40, you earned $50.00 from part time job
// on week 41, you earned $50.00 from part time job
// on week 42, you earned $50.00 from part time job
// on week 43, you earned $50.00 from part time job
// on week 44, you earned $50.00 from part time job
// on week 45, you earned $50.00 from part time job
// on week 46, you earned $50.00 from part time job
// on week 47, you earned $50.00 from part time job
// on week 48, you earned $50.00 from part time job
// on week 49, you earned $50.00 from part time job
// on week 50, you earned $50.00 from part time job
// on week 51, you earned $50.00 from part time job
// on week 52, you earned $50.00 from part time job
// on week 43, you earned $100.00 from investments
// on week 44, you earned $100.00 from investments
// on week 45, you earned $100.00 from investments
// on week 46, you earned $100.00 from investments
// on week 47, you earned $100.00 from investments
// on week 48, you earned $100.00 from investments
// on week 49, you earned $100.00 from investments
// on week 50, you earned $100.00 from investments
// on week 1, you earned $10.00 from gifts
// on week 2, you earned $10.00 from gifts
// on week 3, you earned $10.00 from gifts
// on week 4, you earned $10.00 from gifts
// on week 5, you earned $10.00 from gifts
// on week 6, you earned $10.00 from gifts
// on week 7, you earned $10.00 from gifts
// on week 8, you earned $10.00 from gifts
// on week 9, you earned $10.00 from gifts
// on week 10, you earned $10.00 from gifts
// on week 11, you earned $10.00 from gifts
// on week 12, you earned $10.00 from gifts
// on week 13, you earned $10.00 from gifts
// on week 14, you earned $10.00 from gifts
// on week 15, you earned $10.00 from gifts
// on week 16, you earned $10.00 from gifts
// on week 17, you earned $10.00 from gifts
// on week 18, you earned $10.00 from gifts
// on week 19, you earned $10.00 from gifts
// on week 20, you earned $10.00 from gifts
// on week 21, you earned $10.00 from gifts
// on week 22, you earned $10.00 from gifts
// on week 23, you earned $10.00 from gifts
// on week 24, you earned $10.00 from gifts
// on week 25, you earned $10.00 from gifts
// on week 26, you earned $10.00 from gifts
// on week 27, you earned $10.00 from gifts
// on week 28, you earned $10.00 from gifts
// on week 29, you earned $10.00 from gifts
// on week 30, you earned $10.00 from gifts
// on week 31, you earned $10.00 from gifts
// on week 32, you earned $10.00 from gifts
// on week 33, you earned $10.00 from gifts
// on week 34, you earned $10.00 from gifts
// on week 35, you earned $10.00 from gifts
// on week 36, you earned $10.00 from gifts
// on week 37, you earned $10.00 from gifts
// on week 38, you earned $10.00 from gifts
// on week 39, you earned $10.00 from gifts
// on week 40, you earned $10.00 from gifts
// on week 41, you earned $10.00 from gifts
// on week 42, you earned $10.00 from gifts
// on week 43, you earned $10.00 from gifts
// on week 44, you earned $10.00 from gifts
// on week 45, you earned $10.00 from gifts
// on week 46, you earned $10.00 from gifts
// on week 47, you earned $10.00 from gifts
// on week 48, you earned $10.00 from gifts
// on week 49, you earned $10.00 from gifts
// on week 50, you earned $10.00 from gifts
// on week 51, you earned $10.00 from gifts
// on week 52, you earned $10.00 from gifts
// on week 14, you earned $500.00 from main job
// on week 15, you earned $500.00 from main job
// on week 16, you earned $500.00 from main job
// on week 17, you earned $500.00 from main job
// on week 18, you earned $500.00 from main job
// on week 19, you earned $500.00 from main job
// on week 20, you earned $500.00 from main job
// on week 21, you earned $500.00 from main job
// on week 22, you earned $500.00 from main job
// on week 23, you earned $500.00 from main job
// on week 24, you earned $500.00 from main job
// on week 25, you earned $500.00 from main job
// on week 26, you earned $500.00 from main job
// on week 27, you earned $500.00 from main job
// on week 28, you earned $500.00 from main job
// on week 29, you earned $500.00 from main job
// on week 30, you earned $500.00 from main job
// on week 31, you earned $500.00 from main job
// on week 32, you earned $500.00 from main job
// on week 33, you earned $500.00 from main job
// on week 34, you earned $500.00 from main job
// on week 35, you earned $500.00 from main job
// on week 36, you earned $500.00 from main job
// on week 37, you earned $500.00 from main job
// on week 38, you earned $500.00 from main job
// on week 39, you earned $500.00 from main job
// on week 40, you earned $500.00 from main job
// on week 41, you earned $500.00 from main job
// on week 42, you earned $500.00 from main job
// on week 43, you earned $500.00 from main job
// on week 44, you earned $500.00 from main job
// on week 45, you earned $500.00 from main job
// on week 46, you earned $500.00 from main job
// on week 47, you earned $500.00 from main job
// on week 48, you earned $500.00 from main job
// on week 49, you earned $500.00 from main job
// on week 51, you earned $100.00 from investments
// on week 52, you earned $100.00 from investments
// on week 50, you earned $500.00 from main job
// on week 51, you earned $500.00 from main job
// on week 52, you earned $500.00 from main job
// final bank balance $34320.00
