package main

import "fmt"

type contactInfo struct {
	email string
	zip   int
}

type person struct {
	firstName string
	lastName  string
	//embedding in stucts - we can take custom types as well
	contact contactInfo
	//using just contactinfo it will work
}

// type Person struct {
//     Name string `json:"name"`
//     Age  int    `json:"age"`
// }

/*
Go allows you to define methods on structs. Methods are functions that are associated with a
type. You define methods with a receiver, which is the struct on which the method will
operate.
*/
// func (receiver ReceiverType) MethodName() {
//     // Method body
// }
func (p person) print() {
	fmt.Printf(" %+v ", p)
}

/*
*person - is the pointer to a type(person) - we are looking pointer to a person
*pointerToPerson - takes the pointer and change it to a value
 */
func (pointerToPerson *person) updateName(newFirstName string) {
	// (*pointerToPerson) is the value of jim type
	(*pointerToPerson).firstName = newFirstName
}

func main() {
	// //Creating structs
	// alex := person{"alex", "Anderson"}
	// //alex := person{firstname: "alex",lastname: "Anderson"}
	// fmt.Println(alex)
	// fmt.Printf("%+v", alex)
	// //Updating structs
	// alex.firstName = "bob"
	// fmt.Println(alex)
	//Embedding in stucts
	jim := person{
		firstName: "jim",
		lastName:  "lace",
		contact: contactInfo{
			email: "jim@gmail.com",
			zip:   94000,
		},
	}
	// go works on pass by value - go makes a copy of this data
	// instead of replacing it
	//updating using pointers
	// &variableName - & is a operator which gives us access to the memory address
	// *pointerName - * gives us the value in that memory address
	jimPointer := &jim
	jimPointer.updateName("jimmy")
	//or we can write
	//go allow the shorcut to convert type person to pointer to person if you have pointer as parameter
	//for slices it functions differently
	// incase of slices slice has three paramerters - pointer to head,capacity, length in one address that points to underlying array
	// in case of fnction it is passed by value but it has the pointer
	// these are called refernce type data structure - slices, maps, channels, pointers, functions - dont worry about pointers
	// non refeernce types - int,float,string,bool, structs - use pointers in this case
	// jim.updateName("timmy")
	jim.print()

}
