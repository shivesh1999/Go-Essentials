package main

import (
	"fmt"
	"time"
)

func listenToChan(ch chan int) {
	for {
		i := <-ch
		fmt.Println("Got ", i, " from channel")
		time.Sleep(1 * time.Second)
	}
}
func main() {
	ch := make(chan int, 10)
	go listenToChan(ch)
	for i := 1; i <= 100; i++ {
		ch <- i
	}
	fmt.Println("done!")
	close(ch)
}
