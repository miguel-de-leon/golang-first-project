package main

import "fmt"

type Employee struct {
	id       int
	name     string
	vacation bool
}

// Receiver function
func (e *Employee) SetId(id int) {
	e.id = id
}

func (e *Employee) SetName(name string) {
	e.name = name
}

func (e *Employee) GetId() int {
	return e.id
}

func (e *Employee) GetName() string {
	return e.name
}

func NewEmployee(id int, name string, vacation bool) *Employee {
	return &Employee{id: id, name: name, vacation: vacation}
}

// Main
func main() {
	// 1
	e1 := Employee{}
	fmt.Printf("%v\n", e1)

	// 2
	e2 := Employee{id: 1, name: "Miguel", vacation: true}
	fmt.Printf("%v\n", e2)

	//3 new returns a pointer
	e3 := new(Employee)
	e3.id = 25
	e3.name = "Nombre"
	fmt.Printf("%v\n", *e3)

	// 4
	e4 := NewEmployee(47, "Mike", true)
	fmt.Printf("%v\n", *e4)
}
