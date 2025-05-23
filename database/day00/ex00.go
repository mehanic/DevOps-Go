package main

import (
	"bufio"
	"flag"
	"fmt"
	"math"
	"os"
	"sort"
	"strconv"
)

const (
	maxAllowedNum = 100_000
	minAllowedNum = -100_000
)

var (
	flagMean, flagMedian, flagMode, flagSD bool
)

func init() {
	flag.BoolVar(&flagMean, "mean", false, "Флаг для вывода значения Mean")
	flag.BoolVar(&flagMedian, "median", false, "Флаг для вывода значения Median")
	flag.BoolVar(&flagMode, "mode", false, "Флаг для вывода значения Mode")
	flag.BoolVar(&flagSD, "sd", false, "Флаг для вывода значения SD")
	flag.Parse()
}

// группа целых чисел от -100000 до 100000
func main() {
	array := []int{}
	scaner := bufio.NewScanner(os.Stdin)
	for scaner.Scan() {
		num, err := strconv.Atoi(scaner.Text())
		if err != nil {
			fmt.Println("ошибка: неверный ввод, введите число снова")
			continue
		}
		if num < minAllowedNum || num > maxAllowedNum {
			fmt.Println("вне диапазона")
			continue
		}
		array = append(array, num)
	}
	sort.Ints(array)
	if flagMean {
		fmt.Printf("Mean: %.2f\n", Mean(array))
	}
	if flagMedian {
		fmt.Printf("Median: %.2f\n", Median(array))
	}
	if flagMode {
		fmt.Printf("Mode: %d\n", Mode(array))
	}
	if flagSD {
		fmt.Printf("SD: %.2f\n", SD(array))
	}
}

func Mean(arr []int) float64 {
	var result float64
	for i := 0; i < len(arr); i++ {
		result += float64(arr[i])
	}

	return result / float64(len(arr))
}

func Median(arr []int) float64 {
	if len(arr)%2 == 1 {
		return float64(arr[len(arr)/2])
	}

	return float64(arr[len(arr)/2]+arr[len(arr)/2-1]) / 2.
}

func Mode(arr []int) int {
	var max int
	mymap := make(map[int]int)
	for i := 0; i < len(arr); i++ {
		mymap[arr[i]]++
		if value, ok := mymap[arr[i]]; ok {
			if value > max {
				max = value
			}
		}
	}

	var min_val int
	var count int
	for key, value := range mymap {
		if max == value {
			if count == 0 {
				min_val = key
				count++
			}
			if key < min_val {
				min_val = key
			}
		}
	}

	return min_val
}

func SD(arr []int) float64 {
	mean := Mean(arr)
	var summ float64
	for i := 0; i < len(arr); i++ {
		summ += math.Pow((float64(arr[i]) - mean), 2)
	}
	return math.Sqrt(summ / float64(len(arr)-1))
}
