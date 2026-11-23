package main

import "fmt"

func addItem(items []string, item string) {
	items = append(items, item)
}

func main() {
	a := make([]string, 0, 4)
	a = append(a, "x")
	addItem(a, "y")
	fmt.Println(len(a), a) // line 1 this prints the original slice with the value appended by the append function
	fmt.Println(a[:2])     // line 2  ← the interesting one - this line prints both x and why. Since a copy is passed to add item and the copy still has capacity available since only index 0 is populated it adds y to index 1.
	//However that length is only updated locally inside addItem. the main function headers have not been updated. Nevertheless, the pointer of the slice has data at index 1 which is why when explixiy requestinf a slice of index 0-1 (:2 syntax) the value of y is present

	counts := map[string]int{}
	counts["go"]++
	counts["go"]++
	fmt.Println(counts["go"], counts["rust"]) // line 3 - values start at 0 since its an int so the syntax ++ increments the value fo the go key by 2. accessing a value that doesnt exist such as rust return the value zero value which is zefo in case of int
}

// My prediction for the output is
// ["x"]
// ["x", "y"]
// 2 0
