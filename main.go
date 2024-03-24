package main

import (
	"fmt"
	"math/rand"
)

func makeRandomSlice(numItems, max int) []int {
	s := make([]int, numItems)

	for i := 0; i < numItems; i++ {
		s[i] = rand.Intn(max)
	}

	return s
}

func printSlice(s []int, max int) {
	for i, v := range s {
		if i < max {
			fmt.Println(v)
		} else {
			break
		}
	}
}

func checkSorted(s []int) {
	sorted := true
	for i := 1; i < len(s); i++ {
		if s[i-1] > s[i] {
			sorted = false
		}
	}

	if sorted {
		fmt.Println("The slice is sorted.")
	} else {
		fmt.Println("The slice is NOT sorted!")
	}
}

// quicksort ...
func quicksort(a []int) []int {
	if len(a) < 2 {
		return a
	}

	p := partition(a)

	quicksort(a[:p])

	quicksort(a[p+1:])

	return a
}

// partition divides the slice into two partitions, returning the pivot index.
func partition(a []int) int {
	lo := 0
	hi := len(a) - 1
	pivot := a[hi]

	// temporary pivot index
	p := lo - 1

	// scan the slice and move the elements less than the pivot value before the temporary pivot index
	for i, v := range a {
		if i == len(a)-1 {
			continue
		}

		if v <= pivot {
			p++
			a[p], a[i] = a[i], a[p]
		}
	}

	p++
	a[p], a[hi] = a[hi], a[p]

	return p
}

func main() {
	// Get the number of items and maximum item value.
	var numItems, max int
	fmt.Printf("# Items: ")
	fmt.Scanln(&numItems)
	fmt.Printf("Max: ")
	fmt.Scanln(&max)

	// Make and display the unsorted slice.
	slice := makeRandomSlice(numItems, max)
	printSlice(slice, 40)
	fmt.Println()

	// Sort and display the result.
	quicksort(slice)
	printSlice(slice, 40)

	// Verify that it's sorted.
	checkSorted(slice)
}
