package main

import "fmt"

func main() {
	numbers := [...]int{1, 2, 3, 4, 5}
	sumi := 0
	sumi = sumi + numbers[0]
	sumi = sumi + numbers[1]
	sumi = sumi + numbers[2]
	sumi = sumi + numbers[3]
	sumi = sumi + numbers[4]
	fmt.Println(sumi)
	main1()
}

func main1() {
	numbers := [...]int{1, 2, 3, 4, 5}
	sumi := 0
	for _, val := range numbers {
		sumi += val
	}
	fmt.Println(sumi)
}
