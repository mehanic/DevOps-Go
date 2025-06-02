package main

import "fmt"

func TwoSum(array []int, target int) []int {

	for i := 0; i < len(array)-1; i++ {
		firstNum := array[i]
		for j := i + 1; j < len(array); j++ {
			secondNum := array[j]
			if firstNum + secondNum == target {
				return []int{firstNum, secondNum}
			}	
}

}
return []int{}
}

func main(){
array := []int{3,4,5,6}
target := 7
results := TwoSum(array, target) 
fmt.Println(results)
main1()
}



func main1(){
array := []int{4,5,6}
target := 10
results := TwoSum(array, target)
fmt.Println(results)
}
