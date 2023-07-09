package main

import (
	"fmt"
	"net/http"
	"sync"
)

var endpoints []string

// implementing go routines using waitgroups
var wg sync.WaitGroup //pointer

// implementing mutex
var mut sync.Mutex // pointer

func main() {
	links := []string{
		"http://google.com",
		"http://facebook.com",
		"http://stackoverflow.com",
		"http://golang.org",
		"http://amazon.com",
	}
	for _, link := range links {
		go getStatusCode(link)
		wg.Add(1)
	}

	wg.Wait()
	fmt.Println(endpoints)
}

func getStatusCode(link string) {
	defer wg.Done()
	resp, err := http.Get(link)
	if err != nil {
		fmt.Println("Not able to fetch the status code")
	}
	mut.Lock()
	endpoints = append(endpoints, link)
	mut.Unlock()
	fmt.Println(link, resp.StatusCode)
}
