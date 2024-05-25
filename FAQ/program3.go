package main

import "fmt"

func checkSubset(arr1 []int, arr2 []int) bool {
	arr1Map := make(map[int]bool, len(arr1))
	for i := 0; i < len(arr1); i++ {
		arr1Map[arr1[i]] = true
	}
	for i := 0; i < len(arr2); i++ {
		if _, isPresent := arr1Map[arr2[i]]; !isPresent {
			return false
		}
	}
	return true
}

func main() {
	arr1 := []int{1, 2, 3, 4, 5}
	arr2 := []int{1, 4, 5}
	isSubset := checkSubset(arr1, arr2)
	fmt.Println(isSubset)
}
