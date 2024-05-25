package main

import "fmt"

func fibonacci(num int) {
	x := 0
	y := 1
	for i := 0; i < num; i++ {
		fmt.Println(x)
		z := x + y
		x = y
		y = z
	}
}

func main() {
	fibonacci(10)
}
