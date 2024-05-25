package main

import (
	"fmt"
)

// isSubsetWithOrder checks if arr1 is a subset of arr2 with the same order
func isSubsetWithOrder(arr1, arr2 []int) bool {
	// Initialize two pointers for arr1 and arr2
	i, j := 0, 0

	// Iterate through both arrays while checking the order of elements
	for i < len(arr1) && j < len(arr2) {
		if arr1[i] == arr2[j] {
			i++
		}
		j++
	}

	// If all elements of arr1 were found in arr2 in the same order, it is a subset
	return i == len(arr1)
}

func main() {
	// Example arrays
	arr1 := []int{1, 2, 3, 4}
	arr2 := []int{5, 1, 2, 3, 4, 6}

	// Check if arr1 is a subset of arr2 with the same order
	result := isSubsetWithOrder(arr1, arr2)

	if result {
		fmt.Println("arr1 is a subset of arr2 with the same order")
	} else {
		fmt.Println("arr1 is not a subset of arr2 with the same order")
	}
}
