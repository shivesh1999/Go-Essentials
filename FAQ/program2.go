package main

import "fmt"

func fibonacci(num int, fibonacciChannels chan int) {
	x := 0
	y := 1
	for i := 0; i < num; i++ {
		fibonacciChannels <- x
		x, y = y, x+y
	}
	close(fibonacciChannels)
}

func main() {
	fibonacciChannels := make(chan int)
	go fibonacci(10, fibonacciChannels)
	for num := range fibonacciChannels {
		fmt.Println(num)
	}
}
