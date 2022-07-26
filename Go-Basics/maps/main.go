package main

import (
	"fmt"
)

func main() {
	/*
		maps
		key value pair of data
		statically typed - all keys same type, all value have same type
	*/

	//declaring a map
	colors := map[string]string{
		"red":   "#ff0000",
		"green": "#4bf745",
	}
	fmt.Println(colors)
	//colors:= make(map[string]string)
	//adding values
	colors["white"] = "#ffffff"
	fmt.Println(colors)
	//deleting values
	delete(colors, "white")
	fmt.Println(colors)
	printMap(colors)
}

func printMap(c map[string]string) {
	for color, hex := range c {
		fmt.Println("hex for", color, "is", hex)
	}
}

/*
Difference betweeen structs and maps
maps have same type, structs can have differenr values
map is a refernce type, structs are value type
use map then the values of closesly related properties
use struct when items has many types
*/
