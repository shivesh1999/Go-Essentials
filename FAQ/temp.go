package main

func main() {
	arr1 := []int{1, 2, 3, 4, 2, 3, 1, 5}
	// arr2 := []int{}
	for i := 0; i < len(arr1)-1; i++ {
		for j := i + 1; j < len(arr1); j++ {
			if arr1[j] == arr1[i] {
				continue
			}
		}
	}
}
