package main

import (
	"fmt"
	"net/http"
	"time"
)

func main() {
	links := []string{
		"http://google.com",
		"http://facebook.com",
		"http://stackoverflow.com",
		"http://golang.org",
		"http://amazon.com",
	}
	// creating a channel
	c := make(chan string)
	for _, link := range links {
		go checkLink(link, c)
	}
	// fmt.Println(<-c)
	// fmt.Println(<-c)
	// fmt.Println(<-c)
	// fmt.Println(<-c)
	// fmt.Println(<-c)
	// fmt.Println(<-c)
	for l := range c {

		go func(link string) {
			time.Sleep(5 * time.Second)
			checkLink(link, c)
		}(l)
	}
}

// go routine - takes every line of code and executes one after another
// go scheduler - make only one rountine until its finished
// not executed paralelly- by default it attempts to use one core
// fo muticore processor ischeduler allocates different go rountine in different cores
// concurrency is not parllaelism
// schedule work to be done throghout each other, we dont want one go routine to be over before switching - concurrency
// paralellism is when multyiple things at the same time - possible in multi core processor
// every program has a default go routine
// add go to function call to use go routines

func checkLink(link string, c chan string) {
	_, err := http.Get(link)
	if err != nil {
		fmt.Println(link, "might be down!")
		c <- link
		return
	}
	fmt.Println(link, "is up!")
	c <- link
}

//channels - channels are the only way to communicate with go routines
// data in channesls should have a same type for sharing
// intialize a channel and add it as a argument in the go routine function call
// send data to channel - channel<-5
// receiveing data from a channel - mynumber<-channel
// channels also face blockers when expect values from the channel
// by default it will wait for only one go routine
//recieving somethigg from a channel is blocker
// we never try to access same variable from different child routinr
// we pass it by arguments
