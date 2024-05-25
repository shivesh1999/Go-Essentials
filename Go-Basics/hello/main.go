package main

import "fmt"

func main() {
	fmt.Println("Hello World !")
	var i string = "Hello"
	var j int = 15

	fmt.Printf("i has value: %v and type: %T\n", i, i)
	fmt.Printf("j has value: %v and type: %T", j, j)
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

%v is used to print the value of the arguments
%T is used to print the type of the arguments
The default value of a boolean data type is false
The default value for int is 0, and the default value for string is "".

array
var array_name = [length]datatype{values} // here length is defined
or
var array_name = [...]datatype{values} // here length is inferred

slices
slice_name := []datatype{values}

creating a slice from array
var myarray = [length]datatype{values} // An array
myslice := myarray[start:end] // A slice made from the array

The make() function can also be used to create a slice.
slice_name := make([]type, length, capacity)
slice_name = append(slice_name, element1, element2, ...)
slice3 = append(slice1, slice2...)

*/
