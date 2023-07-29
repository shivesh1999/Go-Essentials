package main

import "fmt"

func print(num int) {
	if num < 1 {
		return
	}
	num--
	print(num)
	fmt.Println(num)
}

func main() {
	print(5)
}
