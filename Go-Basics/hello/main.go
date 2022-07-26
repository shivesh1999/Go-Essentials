package main

import "fmt"

func main() {
	fmt.Println("Hello World !")
}

/*
running the code - go run filename
go commands - go run - compile and immediately execute
			- go build - just compiles it
			- go fmt - automatically format code
			- go install
			- go git
			- go test - executes test files of the code
package main - package is collection on source code files of a project
			   first line should declare tha package it belongs to
			   executable packages - runnable file
			   resuable packages - helper logic is put here and can you used by other projects
			   package main - used for execuable packages
import "fmt" - give our package access to alll the code of other pacakage ("fmt")
			 - fmt is a default package of standard licrary of go
func fuctionName(){} - function declaration

*/
