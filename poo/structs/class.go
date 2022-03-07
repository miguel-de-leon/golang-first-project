package main

import "fmt"

type Employee struct {
	id   int
	name string
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

// Main
func main() {
	e1 := Employee{}
	fmt.Printf("%v", e1)

	e1.id = 1
	e1.name = "Migue"
	fmt.Printf("%v", e1)

	e1.SetId(5)
	e1.SetName("Name2")
	// fmt.Printf("%v", e1)
	fmt.Println(e1.GetId())
	fmt.Println(e1.GetName())
}
