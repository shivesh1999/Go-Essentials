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
// interface type - we cannot create values of the type
// interfaces are not generic types
// interfaces are implicit - go automatically take care of that
// interfaces are a contract to help us manage types - make reuse of code
// go throughh docs while using interfaces

/*
In Go, an interface is a type that specifies a set of method signatures (behavior),
but it doesn’t provide the implementation of those methods. If a type provides
implementations for all the methods declared by an interface, that type implicitly
satisfies the interface, and we can use it as that interface type.
*/

/*
An empty interface (interface{}) is a special case because it doesn't specify any methods.
Any type in Go satisfies the empty interface, which makes it useful for accepting any type.
*/

/*
Go allows an interface to be used in function signatures where multiple types might be returned.
This is useful when a function might return different types of results based on different conditions.

func GetShape(name string) Shape {
    switch name {
    case "circle":
        return Circle{Radius: 5}
    case "square":
        return Square{Side: 4}
    default:
        return nil
    }
}

*/

/*
A type switch allows you to compare an interface value against multiple types in a concise way.
This is similar to a regular switch statement, but instead of comparing values, it compares types.

func Describe(i interface{}) {
    switch v := i.(type) {
    case int:
        fmt.Println("Integer:", v)
    case string:
        fmt.Println("String:", v)
    default:
        fmt.Println("Unknown type")
    }
}

*/

/*
Interfaces are commonly used in Go for mocking in unit tests,
allowing you to replace real implementations with fake ones
to test components in isolation.

type Database interface {
    Save(data string) error
}

type MockDatabase struct {
    data string
}

func (m *MockDatabase) Save(data string) error {
    m.data = data
    return nil
}

func main() {
    db := &MockDatabase{}
    db.Save("mock data")
    fmt.Println(db.data) // Output: mock data
}

*/
