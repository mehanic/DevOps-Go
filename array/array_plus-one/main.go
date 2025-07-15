package array

import (
	"fmt"

	"github.com/qkraudghgh/algorithms/utils"
)

func plusOne(digits []int) []int {
	digits[len(digits)-1] = digits[len(digits)-1] + 1
	res := []int{}
	ten := 0.0
	i := len(digits) - 1
	for i >= 0 || ten == 1 {
		sum := 0
		if i >= 0 {
			sum += digits[i]
		}
		if ten != 0 {
			sum += 1
		}
		res = append(res, sum%10)
		ten = float64(sum) / 10.0
		i -= 1
	}

	return utils.ReverseInts(res)
}

func main() {
	result := plusOne([]int{9, 9, 9})
	fmt.Println(result)
}
