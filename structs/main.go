package main

import "fmt"

type contactInfo struct {
	email   string
	zipCode int
}

type person struct {
	firstName string
	lastName  string
	contactInfo
}

func main() {
	// #1 declare a struct
	// alex := person{"Alex", "Anderson"}

	// #2 declare a struct
	// alex := person{firstName: "Alex", lastName: "Anderson"}

	// #3 declare a struct
	/* var alex person

	alex.firstName = " Alex"
	alex.lastName = " Anderson"

	fmt.Println(alex)
	fmt.Printf("%+v", alex) */

	jim := person{
		firstName: "Jim",
		lastName:  "Party",
		contactInfo: contactInfo{
			email:   "jim@gmail.com",
			zipCode: 94000,
		},
	}

	// jim => type of person
	jim.updateName("Jimmy")
	jim.print()
}

// pointer shortcut
// receiver *person => type of *person, or a pointer to a person
func (pointerToPerson *person) updateName(newFirstName string) {
	(*pointerToPerson).firstName = newFirstName
}

func (p person) print() {
	fmt.Printf("%+v", p)
}
