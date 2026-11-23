package main

import "fmt"

/*
 // Dedup returns the unique strings from in, preserving first-seen order.
 // It must NOT modify in.
 func Dedup(in []string) []string
 Dedup([]string{"b", "a", "b", "c", "a"}) → [b a c]
 Use a set (map[string]struct{}) and pre-size both your map and the result slice.
 In a comment, state the time complexity, and say why it’s better than checking slices.Contains(result, s) in the loop.
*/

func Dedup(in []string) []string {
	seen := make(map[string]struct{}, len(in))
	result := make([]string, 0, len(in))

	for _, s := range in {
		if _, exists := seen[s]; exists {
			continue
		}

		seen[s] = struct{}{}
		result = append(result, s)
	}
	return result
}

func main() {
	a := []string{"c", "b", "c", "d", "b", "b", "a", "c", "d"}
	res := Dedup(a)
	fmt.Println(res)
}
