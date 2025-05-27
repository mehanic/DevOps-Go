package main

import (
	"container/list"
	"fmt"
)

func shortestTransformationSequenceOptimized(start, end string, dictionary []string) int {
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

	startQueue := list.New()
	startQueue.PushBack(start)
	startVisited := map[string]bool{start: true}

	endQueue := list.New()
	endQueue.PushBack(end)
	endVisited := map[string]bool{end: true}

	levelStart, levelEnd := 0, 0

	for startQueue.Len() > 0 && endQueue.Len() > 0 {
		levelStart++
		if exploreLevel(startQueue, startVisited, endVisited, dictSet) {
			return levelStart + levelEnd + 1
		}

		levelEnd++
		if exploreLevel(endQueue, endVisited, startVisited, dictSet) {
			return levelStart + levelEnd + 1
		}
	}

	return 0
}

func exploreLevel(queue *list.List, visited, otherVisited, dictSet map[string]bool) bool {
	lowerCaseAlphabet := "abcdefghijklmnopqrstuvwxyz"

	for size := queue.Len(); size > 0; size-- {
		elem := queue.Front()
		currentWord := elem.Value.(string)
		queue.Remove(elem)

		for i := 0; i < len(currentWord); i++ {
			for _, c := range lowerCaseAlphabet {
				nextWord := currentWord[:i] + string(c) + currentWord[i+1:]
				if otherVisited[nextWord] {
					return true
				}
				if dictSet[nextWord] && !visited[nextWord] {
					visited[nextWord] = true
					queue.PushBack(nextWord)
				}
			}
		}
	}
	return false
}

func main() {
	dictionary := []string{"hit", "hot", "dot", "dog", "cog", "log", "lot"}
	start := "hit"
	end := "cog"
	fmt.Println(shortestTransformationSequenceOptimized(start, end, dictionary)) // Output: 5
}
