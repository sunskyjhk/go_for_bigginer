package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println("Go Sorting Tutorial")

	// our unsorted int array
	unsorted := []int{1, 3, 2, 6, 3, 4}
	fmt.Println(unsorted)

	sort.Ints(unsorted)
	fmt.Println(unsorted)
}
