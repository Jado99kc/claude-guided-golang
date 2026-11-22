package main

import (
	"cmp"
	"fmt"
	"strings"
)

type WordCount struct {
	Word  string
	Count int
}

// TopWords returns the n most frequent words in text.
// Words are case-insensitive ("Go" == "go"), split on whitespace.
// Ties are broken alphabetically. If n > number of unique words, return all.
func TopWords(text string, n int) []WordCount {
	result := []WordCount{}

	if n <= 0 {
		return result
	}

	strArr := strings.Fields(text)
	for _, w := range strArr {
		lowerW := strings.ToLower(w)
		if len(result) == 0 {
			result = append(result, WordCount{Word: lowerW, Count: 1})
			continue
		}
		for _, x := range result {
			lowerX := strings.ToLower(x.Word)
			if cmp.Compare(lowerW, lowerX) == 0 {
				x.Count++
			} else {
				result = append(result, WordCount{Word: lowerW, Count: 1})
			}
		}
	}

	return result
}

func main() {
	a := "the go gopher and the Go team love go"
	fmt.Println(TopWords(a, 2))
}
