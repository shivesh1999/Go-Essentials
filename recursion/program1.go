package main

import "fmt"

func print() {
	fmt.Println(1)
	print()
}

// this program will run into stack overflow
// function call are waiting
func main() {
	print()
}
