package lesson3

import (
	"bufio"
	"fmt"
	"sort"
	"strings"
)

const maxReturnItems = 10

type Token struct {
	Value        string
	EntriesCount int
}

func (token Token) String() string {
	return fmt.Sprintf("Token{Value: %q, EntriesCount: %d}", token.Value, token.EntriesCount)
}

func toSortedSlice(mp *map[string]int) []Token {
	tokensSlice := make([]Token, 0, len(*mp))

	for key, value := range *mp {
		tokensSlice = append(tokensSlice, Token{Value: key, EntriesCount: value})
	}

	sort.Slice(tokensSlice, func(i, j int) bool {
		return tokensSlice[i].EntriesCount > tokensSlice[j].EntriesCount
	})

	return tokensSlice
}

func PrintTenMajorTokens(in string) ([]Token, error) {
	if in == "" {
		return make([]Token, 0), nil
	}

	tokensMap := make(map[string]int)
	scanner := bufio.NewScanner(strings.NewReader(in))
	scanner.Split(bufio.ScanWords)

	for scanner.Scan() {
		token := scanner.Text()
		val, present := tokensMap[token]

		if present {
			val++
			tokensMap[token] = val
			continue
		}

		tokensMap[token] = 1
	}

	if err := scanner.Err(); err != nil {
		return nil, err
	}

	tokensSlice := toSortedSlice(&tokensMap)
	tokensCount := len(tokensSlice)

	if tokensCount > maxReturnItems {
		return tokensSlice[:maxReturnItems], nil
	}

	return tokensSlice, nil
}
