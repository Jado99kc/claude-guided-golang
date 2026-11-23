package main

import "fmt"

func splitAt(s []int, i int) (left, right []int) {
	return s[:i:i], s[i:len(s):len(s)]
}
func main() {

	nums := []int{10, 20, 30, 40, 50}
	fmt.Println("nums:", nums)
	// later:
	left, right := splitAt(nums, 2)
	fmt.Println("left:", left)
	fmt.Println("right:", right)
	left = append(left, 999)
	fmt.Println("left:", left)
	fmt.Println("right:", right)
	fmt.Println("nums:", nums)
}

// My prediction is that the left array will now be [10 20 999] right will be [30 40 50 999] and the nums slice will remain unchanged
