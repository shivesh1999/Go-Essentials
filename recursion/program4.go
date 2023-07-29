package main

import "fmt"

func getSum(i int, sum int) {
	if i < 1 {
		fmt.Println(sum)
		return
	}
	getSum(i-1, sum+i)
}

func getsum(n int) int {
	if n == 0 {
		return 0
	}
	return n + getsum(n-1)
}

func factorial(num int) int {
	if num == 1 {
		return 1
	}
	return num * factorial(num-1)
}

func main() {
	getSum(5, 0)
	fmt.Println(getsum(5))
	fmt.Println(factorial(5))
}

// f(5,0)->f(4,1)->f(3,2)->f(2,3)->f(1,4)->f(0,5)
