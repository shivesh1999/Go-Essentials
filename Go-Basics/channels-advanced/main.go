package main

import (
	"fmt"
	"sync"
)

func main() {
	fmt.Println("channels in go lang")
	mych := make(chan int, 2)
	//second paramenter is used for buffer channel
	wg := &sync.WaitGroup{}

	// -> not avaliable
	//fmt.Println(<-mych)
	wg.Add(2)

	// receive only
	go func(ch <-chan int, wg *sync.WaitGroup) {
		fmt.Println(<-mych)
		wg.Done()
	}(mych, wg)
	// send only
	go func(ch chan<- int, wg *sync.WaitGroup) {
		mych <- 5
		mych <- 6
		close(mych)
		wg.Done()
	}(mych, wg)

	wg.Wait()

}
