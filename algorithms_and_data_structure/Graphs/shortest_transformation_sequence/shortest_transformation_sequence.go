package main

import (
	"container/list"
	"fmt"
)

func shortestTransformationSequence(start, end string, dictionary []string) int {
	dictSet := make(map[string]bool)
	for _, word := range dictionary {
		dictSet[word] = true
	}

	if !dictSet[start] || !dictSet[end] {
		return 0
	}
	if start == end {
		return 1
	}

	lowerCaseAlphabet := "abcdefghijklmnopqrstuvwxyz"
	queue := list.New()
	queue.PushBack(start)
	visited := map[string]bool{start: true}
	dist := 0

	for queue.Len() > 0 {
		levelSize := queue.Len()
		for i := 0; i < levelSize; i++ {
			elem := queue.Front()
			currWord := elem.Value.(string)
			queue.Remove(elem)

			if currWord == end {
				return dist + 1
			}

			for j := 0; j < len(currWord); j++ {
				for _, c := range lowerCaseAlphabet {
					nextWord := currWord[:j] + string(c) + currWord[j+1:]
					if dictSet[nextWord] && !visited[nextWord] {
						visited[nextWord] = true
						queue.PushBack(nextWord)
					}
				}
			}
		}
		dist++
	}

	return 0
}

func main() {
	dict := []string{"hit", "hot", "dot", "dog", "cog", "lot", "log"}
	start := "hit"
	end := "cog"
	fmt.Println(shortestTransformationSequence(start, end, dict)) // Output: 5
}
