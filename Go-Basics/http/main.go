package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
)

type logWriter struct{}

func main() {
	resp, err := http.Get("http://google.com")
	if err != nil {
		fmt.Println("Error :", err)
		os.Exit(1)
	}
	// bs := make([]byte, 99999) // 9999- is the ammount of elements
	// // intailaizing was important as the read function does not automatically resizes the slice
	// resp.Body.Read(bs)
	// fmt.Println(string(bs))
	lw := logWriter{}
	io.Copy(lw, resp.Body)
}

func (logWriter) Write(bs []byte) (int, error) {
	fmt.Println(string(bs))
	fmt.Println("Bytes converted ", len(bs))
	return len(bs), nil
}
