
package array

import "fmt"

func josepheus(intList []int, skip int) {
	skip -= 1
	idx := 0
	for len(intList) > 0 {
		idx = (skip + idx) % len(intList)
		fmt.Printf("%v", intList[idx])
		intList = append(intList[:idx], intList[idx+1:]...) // pop
	}
}

/*
the reason for hashing is that we have to find the index of the item which needs to be removed.
So for e.g. if you iterate with the initial list of folks with every 3rd item eliminated:

INPUT
intList := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
skip := 3

While Iteration:

intList = 123456789
len(intList) = 9
skip = 2 # as intList starts from 0
idx = (0 + 2) % 9 #here previous index was 0
so 3rd element which is 3 in this case eliminated
intList = 12456789
len(intList) = 8
idx = (2 + 2) % 8 #here previous index was 2
so 3rd element starting from 4th person which is 6 would be deleted.
and so on
The reason why we have to do this way is I am not putting the people who escape at the back of list so ideally in 2 while iteration the list should have been
45678912 and then no hashing needed to be done, which means you can directly remove the third element
*/
