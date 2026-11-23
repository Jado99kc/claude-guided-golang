package main

import (
	"cmp"
	"fmt"
	"slices"
	"strings"
)

type WordCount struct {
	Word  string
	Count int
}

// TopWords returns the n most frequent words in text.
// Words are case-insensitive ("Go" == "go"), split on whitespace.
// Ties are broken alphabetically. If n > number of unique words, return all.
// Input: "the go gopher and the Go team love go" with n = 2 → [{go 3} {the 2}]
// Tools you’ll need: strings.Fields, strings.ToLower, slices.SortFunc, cmp.Compare.
// Handle n <= 0 sensibly and document your choice.
func TopWords(text string, n int) []WordCount {
	strArray := strings.Fields(text)
	result := []WordCount{}
	wordMap := make(map[string]WordCount)
	if n <= 0 {
		return result
	}

	for _, s := range strArray {
		lowerW := strings.ToLower(s)
		if _, exists := wordMap[lowerW]; exists {
			tmp := wordMap[lowerW]
			tmp.Count++
			wordMap[lowerW] = tmp
		} else {
			wordMap[lowerW] = WordCount{Word: lowerW, Count: 1}
		}
	}

	for _, wc := range wordMap {
		result = append(result, wc)
	}

	slices.SortFunc(result, func(a, b WordCount) int {
		if c := cmp.Compare(b.Count, a.Count); c != 0 {
			return c
		}

		return cmp.Compare(a.Word, b.Word)
	})

	if n > len(result) {
		return result
	}
	return slices.Clip(result[:n])
}

func main() {
	a := "the go gopher and the Go team love go"
	fmt.Println(TopWords(a, 2))
}
