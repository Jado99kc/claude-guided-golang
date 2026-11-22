package main

import "fmt"

// Dedup returns the unique strings from in, preserving first-seen order.
// It must NOT modify in.
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
	a := []string{"b", "a", "b", "c", "a"}
	res := Dedup(a)

	fmt.Println(res)
}
