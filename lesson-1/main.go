package main

import (
	"errors"
	"fmt"
	"strconv"
)

func main() {
	// fizzbuzz()

	for _, input := range []string{"42", "-5", "abc"} {
		age, err := parseAge(input)
		if err != nil {
			fmt.Println("error:", err)
			continue
		}
		fmt.Println(age)
	}

	final, err := applyDiscount(999, 10)
	if err != nil {
		fmt.Println("error:", err)
		return
	}
	fmt.Println("discounted:", final)
}

func fizzbuzz() {
	for i := range 30 {
		n := i + 1
		switch {
		case n%15 == 0:
			fmt.Println("FizzBuzz")
		case n%3 == 0:
			fmt.Println("Fizz")
		case n%5 == 0:
			fmt.Println("Buzz")
		default:
			fmt.Println(n)
		}

	}
}

func parseAge(s string) (int, error) {
	parsedAge, err := strconv.Atoi(s)
	if err != nil {
		return 0, fmt.Errorf("parse age %q: %w", s, err)
	}

	if parsedAge < 0 {
		return 0, errors.New("age cannot be negative")
	}
	if parsedAge > 150 {
		return 0, errors.New("age cannot be greater than 150")
	}

	return parsedAge, nil

}

func applyDiscount(priceCents int64, percent int) (int64, error) {
	if percent < 0 || percent > 100 {
		return 0, errors.New("percentage must be between 0 - 100")
	}
	// Round half up: scale is 1/100 cent, so +50 adds 0.5 cent before truncating.
	return (priceCents*int64(100-percent) + 50) / 100, nil
}

//BONUS
// var s string
// var n int
// fmt.Println(s == "", n)
//
// This would print true, 0
