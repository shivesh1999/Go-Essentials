package main

import "fmt"

func print(num int) {
	if num == 4 {
		return
	}
	fmt.Println(num)
	num++
	print(num)
}

func main() {
	print(1)
}
