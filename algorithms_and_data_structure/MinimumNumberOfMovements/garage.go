package array

import (
	"fmt"
	"log"
	"reflect"
	"github.com/qkraudghgh/algorithms/utils"
)

func garage(beg, end []int) int {
	i := 0
	moves := 0
	for !reflect.DeepEqual(beg, end) {
		if beg[i] != 0 && beg[i] != end[i] {
			car := beg[i]
			empty, err := utils.FindInt(beg, 0)
			if err != nil {
				log.Fatal(err)
			}
			finalPos, err := utils.FindInt(end, beg[i])
			if err != nil {
				log.Fatal(err)
			}
			if empty != finalPos {
				beg[finalPos], beg[empty] = beg[empty], beg[finalPos]
				fmt.Println(beg)
				empty, err = utils.FindInt(beg, 0)
				if err != nil {
					log.Fatal(err)
				}
				carLotate, err := utils.FindInt(beg, car)
				if err != nil {
					log.Fatal(err)
				}
				beg[carLotate], beg[empty] = beg[empty], beg[carLotate]
				fmt.Println(beg)
				moves += 2
			} else {
				carLotate, err := utils.FindInt(beg, car)
				if err != nil {
					log.Fatal(err)
				}
				beg[carLotate], beg[empty] = beg[empty], beg[carLotate]
				fmt.Println(beg)
				moves += 1
			}
		}
		i += 1
		if i == len(beg) {
			i = 0
		}
	}
	return moves
}