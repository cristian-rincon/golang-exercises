package main

import "fmt"

type Employee struct {
	id       int
	name     string
	vacation bool
}

func NewEmployee (id int, name string, vacation bool ) *Employee {
	return &Employee{
		id:       id,
		name:     name,
		vacation: vacation,
	}
}

// Receiver functions
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

func main() {
	// Create an employee 1
	employee := Employee{}
	employee.SetId(1)
	employee.SetName("Juan")
	fmt.Println(employee)

	// initialize a new Employee 2
	employee2 := Employee{
		id:       2,
		name:     "Pedro",
		vacation: true,
	}
	fmt.Printf("%v\n", employee2)

	// initialize a new Employee 3
	employee3 := new(Employee)
	fmt.Println(employee3)

	// initialize a new Employee 4
	employee4 := NewEmployee(3, "Pablo", false)
	fmt.Printf("%v\n", *employee4)
}
