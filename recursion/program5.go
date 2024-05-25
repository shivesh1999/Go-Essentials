package main

import "fmt"

func reverse(nums []int, i int, n int) {
	fmt.Println(nums)
	if i >= n/2 {
		return
	}
	nums[i], nums[n-i-1] = nums[n-i-1], nums[i]
	reverse(nums, i+1, len(nums))
}

func swap(i int, j int) {
	temp := i
	i = j
	j = temp
}

func main() {
	arr := []int{1, 3, 2, 5, 4}
	reverse(arr, 0, len(arr)-1)
}
