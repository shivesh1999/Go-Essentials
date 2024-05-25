package main

import "fmt"

//all types has a third custom type bot
// if any type has getgreeting function in it with return type string
// that type is a member of bot
// now you can also call greeting function from bot type
type bot interface {
	getGreeting() string
}

// we can also add arguments in interface
// we can return mutiple functions

type englishBot struct{}
type spanishBot struct{}

func printGreeting(b bot) {
	fmt.Println(b.getGreeting())
}

func (eb englishBot) getGreeting() string {
	// very custom logic for english greeting
	return "Hi there !"
}

func (sb spanishBot) getGreeting() string {
	return "Hola"
}

func main() {
	eb := englishBot{}
	sb := spanishBot{}
	printGreeting(eb)
	printGreeting(sb)
}

// overloading is not possible in GO
// cocerete type - we can create values of the type
// interface type - we cannot create calues of the type
// interfaces are not generic types
// interfaces are implicit - go automatically take care of that
// interfaces are a contract to help us manage typesb - make reuse of code
// go throughh docs while using interfaces
